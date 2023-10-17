package repository

import (
	"github.com/GDGVIT/attendance-app-backend/infra/database"
	"github.com/GDGVIT/attendance-app-backend/models"
	"gorm.io/gorm"
)

type TeamMemberRepository struct {
	db *gorm.DB
}

func NewTeamMemberRepository() *TeamMemberRepository {
	return &TeamMemberRepository{database.DB}
}

// CreateTeamMember creates a new TeamMember record.
func (tmr *TeamMemberRepository) CreateTeamMember(teamMember models.TeamMember) (models.TeamMember, error) {
	if err := tmr.db.Create(&teamMember).Error; err != nil {
		return models.TeamMember{}, err
	}
	return teamMember, nil
}

// GetTeamMemberByID retrieves a TeamMember record by its primary keys (TeamID and UserID).
func (tmr *TeamMemberRepository) GetTeamMemberByID(teamID, userID uint) (models.TeamMember, error) {
	var teamMember models.TeamMember
	if err := tmr.db.First(&teamMember, "team_id = ? AND user_id = ?", teamID, userID).Error; err != nil {
		return models.TeamMember{}, err
	}
	return teamMember, nil
}

// UpdateTeamMember updates an existing TeamMember record.
func (tmr *TeamMemberRepository) UpdateTeamMember(teamMember models.TeamMember) (models.TeamMember, error) {
	if err := tmr.db.Save(&teamMember).Error; err != nil {
		return models.TeamMember{}, err
	}
	return teamMember, nil
}

// DeleteTeamMember deletes a TeamMember record by its primary keys (TeamID and UserID).
func (tmr *TeamMemberRepository) DeleteTeamMember(teamID, userID uint) error {
	return tmr.db.Where("team_id = ? AND user_id = ?", teamID, userID).Delete(&models.TeamMember{}).Error
}

// GetTeamMembersByTeamID retrieves all TeamMembers for a given TeamID.
func (tmr *TeamMemberRepository) GetTeamMembersByTeamID(teamID uint) ([]models.TeamMember, error) {
	var teamMembers []models.TeamMember
	if err := tmr.db.Find(&teamMembers, "team_id = ?", teamID).Error; err != nil {
		return nil, err
	}
	return teamMembers, nil
}

// GetTeamMembersByUserID retrieves all TeamMembers for a given UserID.
func (tmr *TeamMemberRepository) GetTeamMembersByUserID(userID uint) ([]models.TeamMember, error) {
	var teamMembers []models.TeamMember
	if err := tmr.db.Find(&teamMembers, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return teamMembers, nil
}

// GetTeamMembersByUserAndRole retrieves all TeamMembers for a given UserID and Role.
func (tmr *TeamMemberRepository) GetTeamMembersByUserAndRole(userID uint, role string) ([]models.TeamMember, error) {
	var teamMembers []models.TeamMember
	if err := tmr.db.Find(&teamMembers, "user_id = ? AND role = ?", userID, role).Error; err != nil {
		return nil, err
	}
	return teamMembers, nil
}

// GetTeamMembersByTeamAndRole retrieves all TeamMembers for a given TeamID and Role.
func (tmr *TeamMemberRepository) GetTeamMembersByTeamAndRole(teamID uint, role string) ([]models.TeamMember, error) {
	var teamMembers []models.TeamMember
	if err := tmr.db.Find(&teamMembers, "team_id = ? AND role = ?", teamID, role).Error; err != nil {
		return nil, err
	}
	return teamMembers, nil
}
