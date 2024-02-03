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

type MeetingRepositoryInterface interface {
	CreateMeeting(meeting models.Meeting) (models.Meeting, error)
	GetMeetingByID(id uint) (models.Meeting, error)
	UpdateMeeting(meeting models.Meeting) (models.Meeting, error)
	DeleteMeetingByID(id uint) error
	GetMeetingsByTeamID(teamID uint) ([]models.Meeting, error)
	GetMeetingsByTeamIDAndMeetingOver(teamID uint, meetingOver bool) ([]models.Meeting, error)
	AddMeetingAttendance(meetingAttendance models.MeetingAttendance) error
	GetMeetingAttendanceByMeetingID(meetingID uint) ([]models.MeetingAttendance, error)
	GetMeetingAttendanceByMeetingIDAndOnTime(meetingID uint, onTime bool) ([]models.MeetingAttendance, error)
	GetMeetingAttendanceByUserIDAndMeetingID(userID, meetingID uint) (models.MeetingAttendance, error)
	GetMeetingAttendancesByUserID(userID uint) ([]models.MeetingAttendance, error)
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

// GetMeetingsByTeamID fetches all meetings of a team.
func (mr *MeetingRepository) GetMeetingsByTeamID(teamID uint) ([]models.Meeting, error) {
	var meetings []models.Meeting
	if err := mr.db.Where("team_id = ?", teamID).Find(&meetings).Error; err != nil {
		return nil, err
	}
	return meetings, nil
}

// GetMeetingsByTeamIDAndMeetingOver
func (mr *MeetingRepository) GetMeetingsByTeamIDAndMeetingOver(teamID uint, meetingOver bool) ([]models.Meeting, error) {
	var meetings []models.Meeting
	if err := mr.db.Where("team_id = ? AND meeting_over = ?", teamID, meetingOver).Find(&meetings).Error; err != nil {
		return nil, err
	}
	return meetings, nil
}

// AddMeetingAttendance adds attendance record for meeting and user
func (mr *MeetingRepository) AddMeetingAttendance(meetingAttendance models.MeetingAttendance) error {
	if err := mr.db.Create(&meetingAttendance).Error; err != nil {
		return err
	}
	return nil
}

// GetMeetingAttendanceByMeetingID fetches all attendance records for a meeting
func (mr *MeetingRepository) GetMeetingAttendanceByMeetingID(meetingID uint) ([]models.MeetingAttendance, error) {
	var meetingAttendance []models.MeetingAttendance
	if err := mr.db.Where("meeting_id = ?", meetingID).Find(&meetingAttendance).Error; err != nil {
		return nil, err
	}
	return meetingAttendance, nil
}

// GetMeetingAttendanceByMeetingIDAndOnTime
func (mr *MeetingRepository) GetMeetingAttendanceByMeetingIDAndOnTime(meetingID uint, onTime bool) ([]models.MeetingAttendance, error) {
	var meetingAttendance []models.MeetingAttendance
	if err := mr.db.Where("meeting_id = ? AND on_time = ?", meetingID, onTime).Find(&meetingAttendance).Error; err != nil {
		return nil, err
	}
	return meetingAttendance, nil
}

// GetMeetingAttendanceByUserIDAndMeetingID
func (mr *MeetingRepository) GetMeetingAttendanceByUserIDAndMeetingID(userID, meetingID uint) (models.MeetingAttendance, error) {
	var meetingAttendance models.MeetingAttendance
	if err := mr.db.Where("user_id = ? AND meeting_id = ?", userID, meetingID).First(&meetingAttendance).Error; err != nil {
		return meetingAttendance, err
	}
	return meetingAttendance, nil
}

// GetMeetingAttendanceByUserID fetches all attendance records for a user across meetings and teams
func (mr *MeetingRepository) GetMeetingAttendancesByUserID(userID uint) ([]models.MeetingAttendance, error) {
	var userAttendance []models.MeetingAttendance
	if err := mr.db.Where("user_id = ?", userID).Find(&userAttendance).Error; err != nil {
		return nil, err
	}
	return userAttendance, nil
}
