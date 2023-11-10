package repository

import (
	"testing"

	"github.com/GDGVIT/attendance-app-backend/models"
	"github.com/GDGVIT/attendance-app-backend/utils/test_utils"
)

func TestTeamRepository_CreateTeam(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.Team{})

	// Create the Team Repository with the test database
	tr := NewTeamRepository()
	tr.db = db

	// Create a test team
	team := models.Team{
		Name:        "Test Team",
		Description: "Test Team Description",
		// Set other fields as needed
	}

	// Test CreateTeam function
	createdTeam, err := tr.CreateTeam(team)
	if err != nil {
		t.Errorf("CreateTeam returned an error: %v", err)
	}

	// Test if the invite code was generated
	if createdTeam.Invite == "" {
		t.Error("Invite code was not generated")
	}

	// Test GetTeamByID function to retrieve the created team
	retrievedTeam, err := tr.GetTeamByID(createdTeam.ID)
	if err != nil {
		t.Errorf("GetTeamByID returned an error: %v", err)
	}

	// Compare the retrieved team with the created team
	if retrievedTeam.Name != team.Name || retrievedTeam.Description != team.Description {
		t.Errorf("Expected team: %+v, got: %+v", team, retrievedTeam)
	}
}

func TestTeamRepository_GetTeamByInvite(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.Team{})

	// Create the Team Repository with the test database
	tr := NewTeamRepository()
	tr.db = db

	// Create a test team
	team := models.Team{
		Name:   "Test Team",
		Invite: "ABC123", // Set the invite code
		// Set other fields as needed
	}

	// Test CreateTeam function
	_, err = tr.CreateTeam(team)
	if err != nil {
		t.Fatalf("Failed to create a test team: %v", err)
	}

	// Test GetTeamByInvite function to retrieve the created team by invite
	retrievedTeam, err := tr.GetTeamByInvite("ABC123")
	if err != nil {
		t.Errorf("GetTeamByInvite returned an error: %v", err)
	}

	// Compare the retrieved team with the created team
	if retrievedTeam.Name != team.Name {
		t.Errorf("Expected team name: %s, got: %s", team.Name, retrievedTeam.Name)
	}
}

func TestTeamRepository_UpdateTeam(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.Team{})

	// Create the Team Repository with the test database
	tr := NewTeamRepository()
	tr.db = db

	// Create a test team
	team := models.Team{
		Name: "Test Team",
		// Set other fields as needed
	}

	// Test CreateTeam function
	createdTeam, err := tr.CreateTeam(team)
	if err != nil {
		t.Fatalf("Failed to create a test team: %v", err)
	}

	// Modify the team's properties
	createdTeam.Name = "Updated Team Name"

	// Test UpdateTeam function
	updatedTeam, err := tr.UpdateTeam(createdTeam)
	if err != nil {
		t.Errorf("UpdateTeam returned an error: %v", err)
	}

	// Test GetTeamByID function to retrieve the updated team
	retrievedTeam, err := tr.GetTeamByID(updatedTeam.ID)
	if err != nil {
		t.Errorf("GetTeamByID returned an error: %v", err)
	}

	// Compare the retrieved team with the updated team
	if retrievedTeam.Name != updatedTeam.Name {
		t.Errorf("Expected team name: %s, got: %s", updatedTeam.Name, retrievedTeam.Name)
	}
}

func TestTeamRepository_GetTeamByID(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.Team{})

	// Create the Team Repository with the test database
	tr := NewTeamRepository()
	tr.db = db

	// Create a test team
	team := models.Team{
		Name: "Test Team",
		// Set other fields as needed
	}

	// Test CreateTeam function
	createdTeam, err := tr.CreateTeam(team)
	if err != nil {
		t.Fatalf("Failed to create a test team: %v", err)
	}

	// Test GetTeamByID function to retrieve the created team
	retrievedTeam, err := tr.GetTeamByID(createdTeam.ID)
	if err != nil {
		t.Errorf("GetTeamByID returned an error: %v", err)
	}

	// Compare the retrieved team with the created team
	if retrievedTeam.Name != team.Name {
		t.Errorf("Expected team name: %s, got: %s", team.Name, retrievedTeam.Name)
	}
}

func TestTeamRepository_DeleteTeamByID(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.Team{})

	// Create the Team Repository with the test database
	tr := NewTeamRepository()
	tr.db = db

	// Create a test team
	team := models.Team{
		Name: "Test Team",
		// Set other fields as needed
	}

	// Test CreateTeam function
	createdTeam, err := tr.CreateTeam(team)
	if err != nil {
		t.Fatalf("Failed to create a test team: %v", err)
	}

	// Test DeleteTeamByID function
	err = tr.DeleteTeamByID(createdTeam.ID)
	if err != nil {
		t.Errorf("DeleteTeamByID returned an error: %v", err)
	}

	// Test GetTeamByID function to retrieve the deleted team
	_, err = tr.GetTeamByID(createdTeam.ID)
	if err == nil {
		t.Error("Expected GetTeamByID to return an error, got nil")
	}
}

func TestTeamRepository_GetUnprotectedTeams(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.Team{})

	// Create the Team Repository with the test database
	tr := NewTeamRepository()
	tr.db = db

	// Create a test team
	team := models.Team{
		Name: "Test Team",
		// Set other fields as needed
	}

	// Test CreateTeam function
	_, err = tr.CreateTeam(team)
	if err != nil {
		t.Fatalf("Failed to create a test team: %v", err)
	}

	// Test GetUnprotectedTeams function
	unprotectedTeams, err := tr.GetUnprotectedTeams()
	if err != nil {
		t.Errorf("GetUnprotectedTeams returned an error: %v", err)
	}

	// Compare the retrieved team with the created team
	if len(unprotectedTeams) != 1 {
		t.Errorf("Expected 1 team, got: %d", len(unprotectedTeams))
	}
}

// test UpdateTeamSuperAdmin
func TestTeamRepository_UpdateTeamSuperAdmin(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.Team{}, &models.TeamMember{})

	// Create the Team Repository with the test database
	tr := NewTeamRepository()
	tr.db = db

	// Create a test team
	team := models.Team{
		Name: "Test Team",
		// Set other fields as needed
	}

	// Test CreateTeam function
	createdTeam, err := tr.CreateTeam(team)
	if err != nil {
		t.Fatalf("Failed to create a test team: %v", err)
	}

	// Create a test user
	user := models.User{
		Name:  "Test User",
		Email: "test@gmail.com",
		// Set other fields as needed
	}

	// Test CreateUser function
	createdUser := user

	// Test UpdateTeamSuperAdmin function
	updatedTeam, err := tr.UpdateTeamSuperAdmin(createdTeam.ID, createdUser.ID)
	if err != nil {
		t.Errorf("UpdateTeamSuperAdmin returned an error: %v", err)
	}

	// Test GetTeamByID function to retrieve the updated team
	retrievedTeam, err := tr.GetTeamByID(updatedTeam.ID)
	if err != nil {
		t.Errorf("GetTeamByID returned an error: %v", err)
	}

	// Compare the retrieved team with the updated team
	if retrievedTeam.SuperAdminID != createdUser.ID {
		t.Errorf("Expected team super admin id: %d, got: %d", createdUser.ID, retrievedTeam.SuperAdminID)
	}

	// Test UpdateTeamSuperAdmin function with invalid team id
	_, err = tr.UpdateTeamSuperAdmin(0, createdUser.ID)
	if err == nil {
		t.Error("Expected UpdateTeamSuperAdmin to return an error, got nil")
	}

	// Test UpdateTeamSuperAdmin function with invalid user id
	// _, err = tr.UpdateTeamSuperAdmin(createdTeam.ID, 0)
	// if err == nil {
	// 	t.Error("Expected UpdateTeamSuperAdmin to return an error, got nil")
	// }
}
