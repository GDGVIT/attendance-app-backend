package repository

import (
	"github.com/GDGVIT/attendance-app-backend/infra/database"
	"github.com/GDGVIT/attendance-app-backend/models"
	"gorm.io/gorm"
)

type TeamRepository struct {
	db *gorm.DB
}

func NewTeamRepository() *TeamRepository {
	return &TeamRepository{database.DB}
}

type TeamRepositoryInterface interface {
	CreateTeam(team models.Team) (models.Team, error)
	GetTeamByID(id uint) (models.Team, error)
	GetTeamByInvite(inviteCode string) (models.Team, error)
	UpdateTeam(team models.Team) (models.Team, error)
	DeleteTeamByID(id uint) error
	GetUnprotectedTeams() ([]models.Team, error)
	UpdateTeamSuperAdmin(teamID, userID uint) (models.Team, error)
}

// CreateTeam creates a new team in the database.
func (tr *TeamRepository) CreateTeam(team models.Team) (models.Team, error) {
	if err := tr.db.Create(&team).Error; err != nil {
		return models.Team{}, err
	}
	return team, nil
}

// GetTeamByID retrieves a team by its ID.
func (tr *TeamRepository) GetTeamByID(id uint) (models.Team, error) {
	var team models.Team
	if err := tr.db.First(&team, id).Error; err != nil {
		return team, err
	}
	return team, nil
}

// GetTeamByInvite retrieves a team by its invite code.
func (tr *TeamRepository) GetTeamByInvite(inviteCode string) (models.Team, error) {
	var team models.Team
	if err := tr.db.Where("invite = ?", inviteCode).First(&team).Error; err != nil {
		return team, err
	}
	return team, nil
}

// UpdateTeam updates an existing team record.
func (tr *TeamRepository) UpdateTeam(team models.Team) (models.Team, error) {
	if err := tr.db.Save(&team).Error; err != nil {
		return models.Team{}, err
	}
	return team, nil
}

// DeleteTeamByID deletes a team by its ID.
func (tr *TeamRepository) DeleteTeamByID(id uint) error {
	return tr.db.Unscoped().Delete(&models.Team{}, id).Error
}

// GetUnprotectedTeams retrieves all unprotected teams.
func (tr *TeamRepository) GetUnprotectedTeams() ([]models.Team, error) {
	var teams []models.Team
	if err := tr.db.Find(&teams, "protected = ?", false).Error; err != nil {
		return nil, err
	}
	return teams, nil
}

// UpdateTeamSuperAdmin changes super admin of a team and returns the updated team.
func (tr *TeamRepository) UpdateTeamSuperAdmin(teamID, userID uint) (models.Team, error) {
	var team models.Team
	if err := tr.db.First(&team, teamID).Error; err != nil {
		return models.Team{}, err
	}
	team.SuperAdminID = userID
	if err := tr.db.Save(&team).Error; err != nil {
		return models.Team{}, err
	}
	return team, nil
}
