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
