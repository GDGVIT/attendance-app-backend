package repository

import (
	"github.com/GDGVIT/attendance-app-backend/infra/database"
	"github.com/GDGVIT/attendance-app-backend/models"
	"gorm.io/gorm"
)

type TeamEntryRequestRepository struct {
	db *gorm.DB
}

func NewTeamEntryRequestRepository() *TeamEntryRequestRepository {
	return &TeamEntryRequestRepository{database.DB}
}

func (ter *TeamEntryRequestRepository) CreateTeamEntryRequest(request models.TeamEntryRequest) (models.TeamEntryRequest, error) {
	if err := ter.db.Create(&request).Error; err != nil {
		return models.TeamEntryRequest{}, err
	}
	return request, nil
}

func (ter *TeamEntryRequestRepository) GetTeamEntryRequestsByTeamID(teamID uint) ([]models.TeamEntryRequest, error) {
	var requests []models.TeamEntryRequest
	if err := ter.db.Where("team_id = ?", teamID).Find(&requests).Error; err != nil {
		return nil, err
	}
	return requests, nil
}

func (ter *TeamEntryRequestRepository) UpdateTeamEntryRequestStatus(requestID uint, status string) (models.TeamEntryRequest, error) {
	// update status and return the updated request
	var request models.TeamEntryRequest
	if err := ter.db.First(&request, requestID).Error; err != nil {
		return models.TeamEntryRequest{}, err
	}
	request.Status = status
	if err := ter.db.Save(&request).Error; err != nil {
		return models.TeamEntryRequest{}, err
	}
	return request, nil
}