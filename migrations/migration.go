package migrations

import (
	"github.com/GDGVIT/attendance-app-backend/infra/database"
	"github.com/GDGVIT/attendance-app-backend/models"
)

// Migrate Add list of model add for migrations
// TODO later separate migration each models
func Migrate() {
	var migrationModels = []interface{}{
		&models.User{},
		&models.VerificationEntry{},
		&models.ForgotPassword{},
		&models.DeletionConfirmation{},
		&models.PasswordAuth{},
		&models.AuthProvider{},
		&models.Team{},
		&models.TeamMember{},
		&models.TeamEntryRequest{},
		&models.Meeting{},
		&models.MeetingAttendance{},
	}
	err := database.DB.AutoMigrate(migrationModels...)
	if err != nil {
		return
	}

	// Remove the 'Password' field from the 'users' table
	// database.DB.Migrator().DropColumn(&models.User{}, "password")
}
