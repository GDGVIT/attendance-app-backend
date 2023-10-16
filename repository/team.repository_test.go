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
	// Add more comparisons for other fields as needed
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
	// Add more comparisons for other fields as needed
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
	// Add more comparisons for other fields as needed
}
