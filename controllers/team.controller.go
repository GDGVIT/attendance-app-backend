package controllers

import (
	"net/http"

	"github.com/GDGVIT/attendance-app-backend/models"
	"github.com/GDGVIT/attendance-app-backend/repository"
	"github.com/gin-gonic/gin"
)

type TeamController struct {
	teamRepo       *repository.TeamRepository
	teamMemberRepo *repository.TeamMemberRepository
}

func NewTeamController() *TeamController {
	teamRepo := repository.NewTeamRepository()
	teamMemberRepo := repository.NewTeamMemberRepository()
	return &TeamController{teamRepo, teamMemberRepo}
}

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
		return
	}

	c.JSON(http.StatusOK, createdTeam)
}
