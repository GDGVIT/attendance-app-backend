package models

import "gorm.io/gorm"

const (
	SuperAdminRole = "super_admin"
	AdminRole      = "admin"
	MemberRole     = "member"
)

type TeamMember struct {
	gorm.Model
	TeamID uint `gorm:"primaryKey"`
	UserID uint `gorm:"primaryKey"`
	Role   string
}
