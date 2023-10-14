package migrations

import (
	"github.com/GDGVIT/attendance-app-backend/infra/database"
	"github.com/GDGVIT/attendance-app-backend/models"
)

// Migrate Add list of model add for migrations
// TODO later separate migration each models
func Migrate() {
	var migrationModels = []interface{}{
		&models.Example{},
		&models.User{},
	}
	err := database.DB.AutoMigrate(migrationModels...)
	if err != nil {
		return
	}
}
