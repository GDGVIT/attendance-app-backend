package controllers

import (
	"net/http"

	"github.com/GDGVIT/attendance-app-backend/infra/logger"
	"github.com/GDGVIT/attendance-app-backend/models"
	"github.com/GDGVIT/attendance-app-backend/repository"
	"github.com/gin-gonic/gin"
)

type TeamController struct {
	teamRepo       *repository.TeamRepository
	teamMemberRepo *repository.TeamMemberRepository
	userRepo       *repository.UserRepository
}

func NewTeamController() *TeamController {
	teamRepo := repository.NewTeamRepository()
	teamMemberRepo := repository.NewTeamMemberRepository()
	userRepo := repository.NewUserRepository()
	return &TeamController{teamRepo, teamMemberRepo, userRepo}
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

	c.JSON(http.StatusOK, createdTeam)
}

// GetTeamByInviteCode retrieves team details by invite code.
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
