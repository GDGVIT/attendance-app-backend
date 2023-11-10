package controllers

import (
	"net/http"

	"github.com/GDGVIT/attendance-app-backend/infra/logger"
	"github.com/GDGVIT/attendance-app-backend/models"
	"github.com/gin-gonic/gin"
)

// GetMyDetails returns details of the authenticated user.
func (uc *UserController) GetMyDetails(c *gin.Context) {
	// Retrieve the authenticated user from the context
	currentUser, _ := c.Get("user")

	// Type-assert the user to the models.User struct
	user := currentUser.(*models.User)

	// Respond with the user details
	c.JSON(http.StatusOK, user)
}

// UpdateMyDetails updates the name and profile picture of the authenticated user.
func (uc *UserController) UpdateMyDetails(c *gin.Context) {
	// Retrieve the authenticated user from the context
	currentUser, _ := c.Get("user")

	// Type-assert the user to the models.User struct
	user := currentUser.(*models.User)

	// Define a struct to bind the JSON request
	var userUpdateRequest struct {
		Name         string `json:"name"`
		ProfileImage string `json:"profile_image"`
	}

	// Bind the JSON request to the userUpdateRequest struct
	if err := c.ShouldBindJSON(&userUpdateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the user's name and profile picture
	if userUpdateRequest.Name != "" {
		user.Name = userUpdateRequest.Name
	}

	if userUpdateRequest.ProfileImage != "" {
		user.ProfileImage = userUpdateRequest.ProfileImage
	}

	// Update the user in the database
	err := uc.userRepo.SaveUser(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	// Respond with the updated user details
	c.JSON(http.StatusOK, gin.H{"message": "User details updated successfully."})
}

// GetMyTeams returns the teams the authenticated user is a member of, along with their role in that team.
func (uc *UserController) GetMyTeams(c *gin.Context) {
	// Retrieve the authenticated user from the context
	currentUser, _ := c.Get("user")

	// Type-assert the user to the models.User struct
	user := currentUser.(*models.User)

	// Get role filter from query params
	role := c.Query("role")

	// Retrieve the teams the user is a member of
	var teamMembers []models.TeamMember
	err := error(nil)
	if role == "" {
		teamMembers, err = uc.teamMemberRepo.GetTeamMembersByUserID(user.ID)
		if err != nil {
			logger.Errorf("Failed to get teams for user %d: "+err.Error(), user.ID)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve teams"})
			return
		}
	} else {
		teamMembers, err = uc.teamMemberRepo.GetTeamMembersByUserAndRole(user.ID, role)
		if err != nil {
			logger.Errorf("Failed to get teams for user %d by %s: "+err.Error(), user.ID, role)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve teams"})
			return
		}
	}

	// Create a slice of {team: models.Team{}, role: string} objects
	teams := make([]struct {
		Team models.Team
		Role string
	}, len(teamMembers))

	// Populate the teams slice with the teams the user is a member of
	for i, teamMember := range teamMembers {
		team, err := uc.teamRepo.GetTeamByID(teamMember.TeamID)
		if err != nil {
			logger.Errorf("Failed to get team %d while populating: "+err.Error(), teamMember.TeamID)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve teams"})
			return
		}

		teams[i].Team = team
		teams[i].Role = teamMember.Role
	}

	// Respond with the teams
	c.JSON(http.StatusOK, teams)
}

// GetMyRequests returns the requests the authenticated user has sent to join teams, with a query filter for the status of the requests.
func (uc *UserController) GetMyRequests(c *gin.Context) {
	// Retrieve the authenticated user from the context
	currentUser, _ := c.Get("user")

	// Type-assert the user to the models.User struct
	user := currentUser.(*models.User)

	// Get status filter from query params
	status := c.Query("status")

	// Retrieve the requests the user has sent
	var requests []models.TeamEntryRequest
	err := error(nil)
	if status == "" {
		requests, err = uc.teamEntryRequestRepo.GetTeamEntryRequestsByUserID(user.ID)
		if err != nil {
			logger.Errorf("Failed to get requests for user %d: "+err.Error(), user.ID)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve requests"})
			return
		}
	} else {
		requests, err = uc.teamEntryRequestRepo.GetTeamEntryRequestsByUserIDAndStatus(user.ID, status)
		if err != nil {
			logger.Errorf("Failed to get requests for user %d by %s: "+err.Error(), user.ID, status)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve requests"})
			return
		}
	}

	// Create a slice of {request: models.TeamEntryRequest{}, team: models.Team{}} objects
	requestsWithTeams := make([]struct {
		Request models.TeamEntryRequest
		Team    models.Team
	}, len(requests))

	// Populate the requestsWithTeams slice with the requests the user has sent
	for i, request := range requests {
		team, err := uc.teamRepo.GetTeamByID(request.TeamID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve requests"})
			return
		}

		requestsWithTeams[i].Request = request
		requestsWithTeams[i].Team = team
	}

	// Respond with the requests
	c.JSON(http.StatusOK, requestsWithTeams)
}
