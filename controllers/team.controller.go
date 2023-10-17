package controllers

import (
	"net/http"
	"strconv"

	"github.com/GDGVIT/attendance-app-backend/infra/logger"
	"github.com/GDGVIT/attendance-app-backend/models"
	"github.com/GDGVIT/attendance-app-backend/repository"
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

// --- Can be done by super admin ---

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
