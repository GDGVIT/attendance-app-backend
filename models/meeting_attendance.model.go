package models

import (
	"time"

	"gorm.io/gorm"
)

type MeetingAttendance struct {
	gorm.Model
	UserID             uint `gorm:"primaryKey"`
	MeetingID          uint `gorm:"primaryKey"`
	AttendanceMarkedAt time.Time
	OnTime             bool `gorm:"default:true"`
}
