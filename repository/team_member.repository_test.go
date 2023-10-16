package repository

import (
	"testing"

	"github.com/GDGVIT/attendance-app-backend/models"
	"github.com/GDGVIT/attendance-app-backend/utils/test_utils"
)

func TestTeamMemberRepository_CreateTeamMember(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.TeamMember{})

	// Create the TeamMember Repository with the test database
	tmr := NewTeamMemberRepository()
	tmr.db = db

	// Create a test team member
	teamMember := models.TeamMember{
		TeamID: 1,
		UserID: 1,
		Role:   models.MemberRole,
	}

	// Test CreateTeamMember function
	_, err = tmr.CreateTeamMember(teamMember)
	if err != nil {
		t.Fatalf("Failed to create a test team member: %v", err)
	}

	// Test GetTeamMemberByID function to retrieve the created team member
	retrievedTeamMember, err := tmr.GetTeamMemberByID(teamMember.TeamID, teamMember.UserID)
	if err != nil {
		t.Errorf("GetTeamMemberByID returned an error: %v", err)
	}

	// Compare the retrieved team member with the created team member
	if retrievedTeamMember.TeamID != teamMember.TeamID {
		t.Errorf("Expected team ID: %d, got: %d", teamMember.TeamID, retrievedTeamMember.TeamID)
	}
	if retrievedTeamMember.UserID != teamMember.UserID {
		t.Errorf("Expected user ID: %d, got: %d", teamMember.UserID, retrievedTeamMember.UserID)
	}
	// Add more comparisons for other fields as needed
}

func TestTeamMemberRepository_UpdateTeamMember(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.TeamMember{})

	// Create the TeamMember Repository with the test database
	tmr := NewTeamMemberRepository()
	tmr.db = db

	// Create a test team member
	teamMember := models.TeamMember{
		TeamID: 1,
		UserID: 1,
		Role:   models.MemberRole,
	}

	// Test CreateTeamMember function
	createdTeamMember, err := tmr.CreateTeamMember(teamMember)
	if err != nil {
		t.Fatalf("Failed to create a test team member: %v", err)
	}

	// Modify the team member's properties
	createdTeamMember.Role = models.AdminRole

	// Test UpdateTeamMember function
	updatedTeamMember, err := tmr.UpdateTeamMember(createdTeamMember)
	if err != nil {
		t.Errorf("UpdateTeamMember returned an error: %v", err)
	}

	// Test GetTeamMemberByID function to retrieve the updated team member
	retrievedTeamMember, err := tmr.GetTeamMemberByID(updatedTeamMember.TeamID, updatedTeamMember.UserID)
	if err != nil {
		t.Errorf("GetTeamMemberByID returned an error: %v", err)
	}

	// Compare the retrieved team member with the updated team member
	if retrievedTeamMember.Role != updatedTeamMember.Role {
		t.Errorf("Expected role: %s, got: %s", updatedTeamMember.Role, retrievedTeamMember.Role)
	}
	// Add more comparisons for other fields as needed
}

func TestTeamMemberRepository_DeleteTeamMember(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.TeamMember{})

	// Create the TeamMember Repository with the test database
	tmr := NewTeamMemberRepository()
	tmr.db = db

	// Create a test team member
	teamMember := models.TeamMember{
		TeamID: 1,
		UserID: 1,
		Role:   models.MemberRole,
	}

	// Test CreateTeamMember function
	createdTeamMember, err := tmr.CreateTeamMember(teamMember)
	if err != nil {
		t.Fatalf("Failed to create a test team member: %v", err)
	}

	// Test DeleteTeamMember function to delete the team member by TeamID and UserID
	err = tmr.DeleteTeamMember(createdTeamMember.TeamID, createdTeamMember.UserID)
	if err != nil {
		t.Errorf("DeleteTeamMember returned an error: %v", err)
	}

	// Test GetTeamMemberByID function to check if the team member is deleted
	_, err = tmr.GetTeamMemberByID(createdTeamMember.TeamID, createdTeamMember.UserID)
	if err == nil {
		t.Error("Team member was not deleted, GetTeamMemberByID returned no error")
	} else if err.Error() != "record not found" {
		t.Errorf("Expected 'record not found' error, got: %v", err)
	}
}

func TestTeamMemberRepository_GetTeamMembersByTeamID(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.TeamMember{})

	// Create the TeamMember Repository with the test database
	tmr := NewTeamMemberRepository()
	tmr.db = db

	// Create test team members with the same TeamID
	teamID := uint(1)
	teamMembers := []models.TeamMember{
		{TeamID: teamID, UserID: 1, Role: models.MemberRole},
		{TeamID: teamID, UserID: 2, Role: models.AdminRole},
		{TeamID: teamID, UserID: 3, Role: models.MemberRole},
	}

	for _, teamMember := range teamMembers {
		_, err = tmr.CreateTeamMember(teamMember)
		if err != nil {
			t.Fatalf("Failed to create a test team member: %v", err)
		}
	}

	// Test GetTeamMembersByTeamID function to retrieve team members by TeamID
	retrievedTeamMembers, err := tmr.GetTeamMembersByTeamID(teamID)
	if err != nil {
		t.Errorf("GetTeamMembersByTeamID returned an error: %v", err)
	}

	// Check the number of retrieved team members
	expectedCount := len(teamMembers)
	if len(retrievedTeamMembers) != expectedCount {
		t.Errorf("Expected %d team members, got %d", expectedCount, len(retrievedTeamMembers))
	}

	// Verify the correctness of the retrieved team members
	for i, retrieved := range retrievedTeamMembers {
		if retrieved.UserID != teamMembers[i].UserID {
			t.Errorf("Mismatch in UserID. Expected: %d, Got: %d", teamMembers[i].UserID, retrieved.UserID)
		}
		if retrieved.Role != teamMembers[i].Role {
			t.Errorf("Mismatch in Role. Expected: %s, Got: %s", teamMembers[i].Role, retrieved.Role)
		}
		// Add more comparisons for other fields as needed
	}
}
