package repository

import (
	"github.com/GDGVIT/attendance-app-backend/infra/database"
	"github.com/GDGVIT/attendance-app-backend/models"
	"gorm.io/gorm"
)

type MeetingRepository struct {
	db *gorm.DB
}

func NewMeetingRepository() *MeetingRepository {
	return &MeetingRepository{database.DB}
}

// CreateMeeting creates a new meeting in the database.
func (mr *MeetingRepository) CreateMeeting(meeting models.Meeting) (models.Meeting, error) {
	// should return error if not null field doesnt have value
	if err := mr.db.Create(&meeting).Error; err != nil {
		return models.Meeting{}, err
	}
	return meeting, nil
}

// GetMeetingByID retrieves a meeting by its ID.
func (mr *MeetingRepository) GetMeetingByID(id uint) (models.Meeting, error) {
	var meeting models.Meeting
	if err := mr.db.First(&meeting, id).Error; err != nil {
		return meeting, err
	}
	return meeting, nil
}

// UpdateMeeting updates an existing meeting record.
func (mr *MeetingRepository) UpdateMeeting(meeting models.Meeting) (models.Meeting, error) {
	if err := mr.db.Save(&meeting).Error; err != nil {
		return models.Meeting{}, err
	}
	return meeting, nil
}

// DeleteMeetingByID deletes a meeting by its ID.
func (mr *MeetingRepository) DeleteMeetingByID(id uint) error {
	return mr.db.Delete(&models.Meeting{}, id).Error
}
