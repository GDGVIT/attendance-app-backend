package models

import "gorm.io/gorm"

const (
	TeamEntryRequestPending  = "pending"
	TeamEntryRequestApproved = "approved"
	TeamEntryRequestRejected = "rejected"
)

type TeamEntryRequest struct {
	gorm.Model
	TeamID uint   `gorm:"not null"`
	UserID uint   `gorm:"not null"`
	Status string `gorm:"not null;default:'pending'"` // 'pending', 'approved', 'rejected'
}
