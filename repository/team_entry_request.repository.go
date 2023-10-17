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

func (ter *TeamEntryRequestRepository) GetTeamEntryRequestByID(requestID uint) (models.TeamEntryRequest, error) {
	var request models.TeamEntryRequest
	if err := ter.db.First(&request, requestID).Error; err != nil {
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

// GetTeamEntryRequestsByUserID retrieves all TeamEntryRequests for a given UserID.
func (ter *TeamEntryRequestRepository) GetTeamEntryRequestsByUserID(userID uint) ([]models.TeamEntryRequest, error) {
	var requests []models.TeamEntryRequest
	if err := ter.db.Find(&requests, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return requests, nil
}

// GetTeamEntryRequestsByUserIDAndStatus retrieves all TeamEntryRequests for a given UserID and status.
func (ter *TeamEntryRequestRepository) GetTeamEntryRequestsByUserIDAndStatus(userID uint, status string) ([]models.TeamEntryRequest, error) {
	var requests []models.TeamEntryRequest
	if err := ter.db.Find(&requests, "user_id = ? AND status = ?", userID, status).Error; err != nil {
		return nil, err
	}
	return requests, nil
}

// GetTeamEntryRequestsByTeamIDAndStatus retrieves all TeamEntryRequests for a given TeamID and status.
func (ter *TeamEntryRequestRepository) GetTeamEntryRequestsByTeamIDAndStatus(teamID uint, status string) ([]models.TeamEntryRequest, error) {
	var requests []models.TeamEntryRequest
	if err := ter.db.Find(&requests, "team_id = ? AND status = ?", teamID, status).Error; err != nil {
		return nil, err
	}
	return requests, nil
}

// GetTeamEntryRequestByTeamIDAndUserID retrieves a TeamEntryRequest by its TeamID and UserID.
func (ter *TeamEntryRequestRepository) GetTeamEntryRequestByTeamIDAndUserID(teamID, userID uint) (models.TeamEntryRequest, error) {
	var request models.TeamEntryRequest
	if err := ter.db.First(&request, "team_id = ? AND user_id = ?", teamID, userID).Error; err != nil {
		return models.TeamEntryRequest{}, err
	}
	return request, nil
}

// DeleteTeamEntryRequestByID deletes a TeamEntryRequest by its ID.
func (ter *TeamEntryRequestRepository) DeleteTeamEntryRequestByID(requestID uint) error {
	return ter.db.Where("id = ?", requestID).Delete(&models.TeamEntryRequest{}).Error
}
