package models

import (
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
	TeamID           uint     `gorm:"not null"`
	Title            string   `gorm:"size:255;not null;"`
	Description      string   `gorm:"size:255;not null;"`
	Venue            string   `gorm:"size:255;not null;"`
	Location         Location `gorm:"embedded"`
	StartTime        int64    `gorm:"not null"` // Unix timestamp, for info purposes only. Attendance will start on manual start.
	MeetingPeriod    bool     `gorm:"default:false"`
	AttendancePeriod bool     `gorm:"default:false"` // Members can mark attendance while true. Can only be started after meeting has started. Is ended alongside meeting end if not ended before.
	MeetingOver      bool     `gorm:"default:false"` // Will not show meeting on dashboard if true, can be seen in some history tab
}
