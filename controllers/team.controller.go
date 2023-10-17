package controllers

import (
	"net/http"
	"strconv"

	"github.com/GDGVIT/attendance-app-backend/infra/logger"
	"github.com/GDGVIT/attendance-app-backend/models"
	"github.com/GDGVIT/attendance-app-backend/repository"
	"github.com/GDGVIT/attendance-app-backend/utils/email"
	"github.com/GDGVIT/attendance-app-backend/utils/team"
	"github.com/gin-gonic/gin"
)

type TeamController struct {
	teamRepo             *repository.TeamRepository
	teamMemberRepo       *repository.TeamMemberRepository
	userRepo             *repository.UserRepository
	teamEntryRequestRepo *repository.TeamEntryRequestRepository
}

func NewTeamController() *TeamController {
	teamRepo := repository.NewTeamRepository()
	teamMemberRepo := repository.NewTeamMemberRepository()
	userRepo := repository.NewUserRepository()
	teamEntryRequestRepo := repository.NewTeamEntryRequestRepository()
	return &TeamController{teamRepo, teamMemberRepo, userRepo, teamEntryRequestRepo}
}

// --- Can be done by any logged in user ---

// CreateTeam creates a new team in the database and designates creator as super admin.
func (tc *TeamController) CreateTeam(c *gin.Context) {
	// Extract data from the JSON request
	var teamRequest struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		Protected   bool   `json:"protected"`
	}

	if err := c.ShouldBindJSON(&teamRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the current user (super admin)
	currentUser, _ := c.Get("user")
	superAdmin, ok := currentUser.(*models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get the current user"})
		return
	}

	// Create the team with SuperAdminID set
	team := models.Team{
		Name:         teamRequest.Name,
		Description:  teamRequest.Description,
		Protected:    teamRequest.Protected,
		SuperAdminID: superAdmin.ID,
	}

	// Create the team in the database
	createdTeam, err := tc.teamRepo.CreateTeam(team)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create team"})
		logger.Errorf("Failed to create team: " + err.Error())
		return
	}

	// Create a TeamMember entry for the super admin
	teamMember := models.TeamMember{
		TeamID: createdTeam.ID,
		UserID: superAdmin.ID,
		Role:   models.SuperAdminRole,
	}

	_, err = tc.teamMemberRepo.CreateTeamMember(teamMember)
	if err != nil {
		// Handle the error, for example, by rolling back the team creation
		tc.teamRepo.DeleteTeamByID(createdTeam.ID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create team member"})
		logger.Errorf("Failed to create team member: " + err.Error())
		return
	}

	c.JSON(http.StatusCreated, createdTeam)
}

// GetTeamByInviteCode retrieves team details by invite code.
// Intended for users pulling up basic team details before joining.
func (tc *TeamController) GetTeamByInviteCode(c *gin.Context) {
	inviteCode := c.Param("inviteCode")

	team, err := tc.teamRepo.GetTeamByInvite(inviteCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid invite."})
		return
	}

	superAdmin, err := tc.userRepo.GetUserByID(team.SuperAdminID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve super admin"})
		logger.Errorf("Failed to retrieve super admin: " + err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"team":       team,
		"superAdmin": superAdmin,
	})
}

// JoinTeamByInviteCode joins a team by invite code. If unprotected, the user is added as a member. If protected, the user is added as a pending member via TeamRequests.
func (tc *TeamController) JoinTeamByInviteCode(c *gin.Context) {
	// Get the current user
	currentUser, _ := c.Get("user")
	user, ok := currentUser.(*models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get the current user"})
		return
	}

	// Get the team by invite code
	inviteCode := c.Param("inviteCode")
	team, err := tc.teamRepo.GetTeamByInvite(inviteCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid invite."})
		return
	}

	// Check if the user is already a member of the team
	teamMember, err := tc.teamMemberRepo.GetTeamMemberByID(team.ID, user.ID)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User is already a member of the team"})
		return
	}

	// Check if the team is protected
	if team.Protected {
		// Create a TeamRequest
		teamRequest := models.TeamEntryRequest{
			TeamID: team.ID,
			UserID: user.ID,
		}

		_, err := tc.teamEntryRequestRepo.CreateTeamEntryRequest(teamRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create team request"})
			logger.Errorf("Failed to create team request: " + err.Error())
			return
		}

		adminTeam, err := tc.teamMemberRepo.GetAdminTeamByTeamID(team.ID)
		// list of admin team emails
		var adminEmails []string
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve admin team"})
			logger.Errorf("Failed to retrieve admin team: " + err.Error())
			return
		}
		for _, admin := range adminTeam {
			adminUser, err := tc.userRepo.GetUserByID(admin.UserID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve admin user"})
				logger.Errorf("Failed to retrieve admin user: " + err.Error())
				return
			}
			adminEmails = append(adminEmails, adminUser.Email)
		}

		// Email the admins and super admins of the team
		email.SendRequestNotifToTeamAdmins(adminEmails, user.Name, user.Email, team.Name)

		c.JSON(http.StatusCreated, gin.H{"message": "Team request created", "protected": true})
		return
	}

	// Create a TeamMember entry
	teamMember = models.TeamMember{
		TeamID: team.ID,
		UserID: user.ID,
		Role:   models.MemberRole,
	}

	_, err = tc.teamMemberRepo.CreateTeamMember(teamMember)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create team member"})
		logger.Errorf("Failed to create team member: " + err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Team member created", "protected": false})
}

// --- Can be done by team member, admin and super admin ---

// GetTeamByID retrieves a team by ID.
func (tc *TeamController) GetTeamByID(c *gin.Context) {
	// Get the team ID from the route parameter
	teamID, err := strconv.Atoi(c.Param("teamID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}

	// Get the team
	team, err := tc.teamRepo.GetTeamByID(uint(teamID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}

	c.JSON(http.StatusOK, team)
}

// GetTeamMembers retrieves all team members for a given team.
func (tc *TeamController) GetTeamMembers(c *gin.Context) {
	// Get the team ID from the route parameter
	teamID, err := strconv.Atoi(c.Param("teamID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}

	// Get the team
	_, err = tc.teamRepo.GetTeamByID(uint(teamID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}

	// Get role filter from query params
	role := c.Query("role")

	var teamMembers []models.TeamMember
	if role == "" {
		// Retrieve the team members
		teamMembers, err = tc.teamMemberRepo.GetTeamMembersByTeamID(uint(teamID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve team members"})
			return
		}
	} else {
		// Retrieve the team members by role
		teamMembers, err = tc.teamMemberRepo.GetTeamMembersByTeamAndRole(uint(teamID), role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve team members"})
			return
		}
	}

	// Create a slice of users, with corresponding roles from teammembers added to each
	users := make([]struct {
		User models.User
		Role string
	}, len(teamMembers))

	// Populate the teamMembers slice with the team members and their corresponding users
	for i, teamMember := range teamMembers {
		user, err := tc.userRepo.GetUserByID(teamMember.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
			return
		}
		users[i] = struct {
			User models.User
			Role string
		}{user, teamMember.Role}
	}

	c.JSON(http.StatusOK, users)
}

// --- Can be done by team admin and super admin ---

// GetTeamRequests retrieves all team requests for a given team, and allows filtering by status.
func (tc *TeamController) GetTeamRequests(c *gin.Context) {
	// Get the team ID from the route parameter
	teamID, err := strconv.Atoi(c.Param("teamID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}

	// Get the team
	_, err = tc.teamRepo.GetTeamByID(uint(teamID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}

	// Get status filter from query params
	status := c.Query("status")

	// Retrieve the team requests
	var requests []models.TeamEntryRequest
	if status == "" {
		requests, err = tc.teamEntryRequestRepo.GetTeamEntryRequestsByTeamID(uint(teamID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve team requests"})
			return
		}
	} else {
		requests, err = tc.teamEntryRequestRepo.GetTeamEntryRequestsByTeamIDAndStatus(uint(teamID), status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve team requests"})
			return
		}
	}

	// Create a slice of {request: models.TeamEntryRequest{}, user: models.User{}} objects
	requestsWithUsers := make([]struct {
		Request models.TeamEntryRequest
		User    models.User
	}, len(requests))

	// Populate the requests slice with the requests and their corresponding users
	for i, request := range requests {
		user, err := tc.userRepo.GetUserByID(request.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
			return
		}
		requestsWithUsers[i] = struct {
			Request models.TeamEntryRequest
			User    models.User
		}{request, user}
	}

	c.JSON(http.StatusOK, requestsWithUsers)
}

// UpdateTeamRequestStatus updates the status of a team request Patch /team/:teamID/requests/:requestID
func (tc *TeamController) UpdateTeamRequestStatus(c *gin.Context) {
	// Get the request ID from the route parameter
	requestID, err := strconv.Atoi(c.Param("requestID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}

	// Get the request
	request, err := tc.teamEntryRequestRepo.GetTeamEntryRequestByID(uint(requestID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Request not found"})
		return
	}

	if request.Status == models.TeamEntryRequestApproved {
		c.JSON(http.StatusBadRequest, gin.H{"error": "already-approved", "message": "Request has already been approved. You can remove the user from the team instead if you wish to do so."})
		return
	}

	// Get the team
	team, err := tc.teamRepo.GetTeamByID(request.TeamID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}

	// Get the status from the JSON request
	var requestUpdateRequest struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestUpdateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid input."})
		return
	}

	// Update the request status
	updatedRequest, err := tc.teamEntryRequestRepo.UpdateTeamEntryRequestStatus(request.ID, requestUpdateRequest.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to update request"})
		return
	}

	// If the request was accepted, add the user as a member of the team
	if requestUpdateRequest.Status == models.TeamEntryRequestApproved {
		teamMember := models.TeamMember{
			TeamID: request.TeamID,
			UserID: request.UserID,
			Role:   models.MemberRole,
		}

		_, err := tc.teamMemberRepo.CreateTeamMember(teamMember)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create team member"})
			logger.Errorf("Failed to create team member: " + err.Error())
			return
		}
	}

	// Get the user
	user, err := tc.userRepo.GetUserByID(request.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}

	// Email the user
	email.SendRequestStatusNotifToUser(user.Email, user.Name, team.Name, updatedRequest.Status)

	c.JSON(http.StatusOK, updatedRequest)
}

// --- Can be done by team super admin ---

// UpdateTeam updates a team's name and description.
func (tc *TeamController) UpdateTeam(c *gin.Context) {
	// Bind the JSON request to a TeamUpdateRequest struct
	var teamUpdateRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&teamUpdateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Get the team ID from the route parameter
	teamID, err := strconv.Atoi(c.Param("teamID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}

	// Update the team's name and description
	team, err := tc.teamRepo.GetTeamByID(uint(teamID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}

	if teamUpdateRequest.Name != "" {
		team.Name = teamUpdateRequest.Name
	}
	if teamUpdateRequest.Description != "" {
		team.Description = teamUpdateRequest.Description
	}

	// Save the updated team
	updatedTeam, err := tc.teamRepo.UpdateTeam(team)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update team"})
		return
	}

	// Respond with the updated team details
	c.JSON(http.StatusOK, updatedTeam)
}

// RegenerateInviteCode regenerates a team's invite code.
func (tc *TeamController) RegenerateInviteCode(c *gin.Context) {
	// Get the team ID from the route parameter
	teamID, err := strconv.Atoi(c.Param("teamID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}

	// Get the team
	currteam, err := tc.teamRepo.GetTeamByID(uint(teamID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}

	// Regenerate the invite code
	currteam.Invite = team.GenerateInviteCode()

	// Save the updated team
	updatedTeam, err := tc.teamRepo.UpdateTeam(currteam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update team"})
		return
	}

	// Respond with the updated team details
	c.JSON(http.StatusOK, updatedTeam)
}
