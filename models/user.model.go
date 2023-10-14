package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string `gorm:"size:255;not null;unique;"`
	Password     string `gorm:"size:255;not null;"`
	Name         string `gorm:"size:255;not null;"`
	ProfileImage string `gorm:"size:255;"`
	Verified     bool   `gorm:"default:false"`
}

func (user *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}
