package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// Location struct to hold lat long and alt for Meeting
type Location struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
}

// Members can start marking attendance after meeting has been started (MeetingPeriod = true), and attendance is open (AttendancePeriod = true). Their attendance will be OnTime = true.
// If they mark attendance after attendance closed (AttendancePeriod = false), but while meeting still ongoing (MeetingPeriod = True), their attendance will be OnTime = false.
// They cannot mark attendance after meeting has ended (MeetingOver = true), which is set when MeetingPeriod = true -> false.
// A meeting can only be deleted if MeetingPeriod = false and AttendancePeriod = false and MeetingOver = false. I.e., meeting hasn't started yet.
type Meeting struct {
	gorm.Model
	TeamID           uint      `gorm:"not null"`
	Title            string    `gorm:"size:255;not null"`
	Description      string    `gorm:"size:255;not null"`
	Venue            string    `gorm:"size:255;not null"`
	Location         Location  `gorm:"embedded"`
	StartTime        time.Time `gorm:"not null"` // Unix timestamp, for info purposes only. Attendance will start on manual start.
	MeetingPeriod    bool      `gorm:"default:false"`
	AttendancePeriod bool      `gorm:"default:false"` // Members can mark attendance while true. Can only be started after meeting has started. Is ended alongside meeting end if not ended before.
	MeetingOver      bool      `gorm:"default:false"` // Will not show meeting on dashboard if true, can be seen in some history tab
	AttendanceOver   bool      `gorm:"default:false"`
}

// add isvalid check to model to check if venue, title, description are not empty strings or missing
// add isvalid check to model to check if location is valid
// add isvalid check to model to check if starttime is in the future
// all this in a beforecreate hook
func (m *Meeting) BeforeCreate(tx *gorm.DB) error {
	if m.TeamID == 0 {
		return gorm.ErrInvalidData
	}
	if m.Venue == "" || m.Title == "" || m.Description == "" {
		return errors.New("meeting venue, title or description cannot be empty")
	}
	if m.StartTime.Before(time.Now()) {
		return errors.New("meeting start time cannot be in the past")
	}
	if m.MeetingPeriod || m.AttendancePeriod || m.MeetingOver || m.AttendanceOver {
		return errors.New("meeting cannot be created with any of the periods set to true")
	}
	return nil
}

type MeetingAttendance struct {
	gorm.Model
	UserID             uint      `gorm:"primaryKey;not null"`
	MeetingID          uint      `gorm:"primaryKey;not null"`
	AttendanceMarkedAt time.Time `gorm:"not null"`
	OnTime             bool
}

func (ma *MeetingAttendance) BeforeCreate(tx *gorm.DB) error {
	if ma.UserID == 0 || ma.MeetingID == 0 {
		return gorm.ErrInvalidData
	}
	if ma.AttendanceMarkedAt.IsZero() {
		return errors.New("attendance marked at cannot be zero time")
	}
	return nil
}

type MeetingAttendanceListResponse struct {
	ID                 uint
	MeetingID          uint
	AttendanceMarkedAt time.Time
	OnTime             bool
	User               User
	MeetingName        string
	TeamName           string
}

type UserUpcomingMeetingsListResponse struct {
	Meeting Meeting
	Team    Team
}
