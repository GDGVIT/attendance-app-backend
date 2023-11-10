package models

import (
	"github.com/GDGVIT/attendance-app-backend/utils/team"
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	Name         string `gorm:"not null;unique"`
	Description  string
	SuperAdminID uint // Foreign key to the user who is the super admin of this team
	// Meetings    []Meeting
	Protected bool   `gorm:"default:false"`   // If true, then users will need to be approved by the super admin to join this team
	Invite    string `gorm:"unique;not null"` // Invite code for this team, length 10
}

// gorm on create hook to generate invite code if not provided
func (t *Team) BeforeCreate(tx *gorm.DB) error {
	if t.Invite == "" {
		t.Invite = team.GenerateInviteCode()
	}
	return nil
}
