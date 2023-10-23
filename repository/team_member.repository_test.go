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

// Test GetTeamMemberByID
func TestTeamMemberRepository_GetTeamMemberByID(t *testing.T) {
	// Setup the test database
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}

	// Drop the table after testing
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

func TestTeamMemberRepository_GetTeamMembersByUserID(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.TeamMember{})

	// Create the TeamMember Repository with the test database
	tmr := NewTeamMemberRepository()
	tmr.db = db

	// Create test team members with the same UserID
	userID := uint(1)
	teamMembers := []models.TeamMember{
		{TeamID: 1, UserID: userID, Role: models.MemberRole},
		{TeamID: 2, UserID: userID, Role: models.AdminRole},
		{TeamID: 3, UserID: userID, Role: models.MemberRole},
	}

	for _, teamMember := range teamMembers {
		_, err = tmr.CreateTeamMember(teamMember)
		if err != nil {
			t.Fatalf("Failed to create a test team member: %v", err)
		}
	}

	// Test GetTeamMembersByUserID function to retrieve team members by UserID
	retrievedTeamMembers, err := tmr.GetTeamMembersByUserID(userID)
	if err != nil {
		t.Errorf("GetTeamMembersByUserID returned an error: %v", err)
	}

	// Check the number of retrieved team members
	expectedCount := len(teamMembers)
	if len(retrievedTeamMembers) != expectedCount {
		t.Errorf("Expected %d team members, got %d", expectedCount, len(retrievedTeamMembers))
	}

	// Verify the correctness of the retrieved team members
	for i, retrieved := range retrievedTeamMembers {
		if retrieved.TeamID != teamMembers[i].TeamID {
			t.Errorf("Mismatch in TeamID. Expected: %d, Got: %d", teamMembers[i].TeamID, retrieved.TeamID)
		}
		if retrieved.Role != teamMembers[i].Role {
			t.Errorf("Mismatch in Role. Expected: %s, Got: %s", teamMembers[i].Role, retrieved.Role)
		}
		// Add more comparisons for other fields as needed
	}
}

func TestTeamMemberRepository_GetTeamMembersByUserAndRole(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.TeamMember{})

	// Create the TeamMember Repository with the test database
	tmr := NewTeamMemberRepository()
	tmr.db = db

	// Create test team members with the same UserID and Role
	userID := uint(1)
	role := models.MemberRole
	teamMembers := []models.TeamMember{
		{TeamID: 1, UserID: userID, Role: role},
		{TeamID: 2, UserID: userID, Role: role},
		{TeamID: 3, UserID: userID, Role: role},
	}

	for _, teamMember := range teamMembers {
		_, err = tmr.CreateTeamMember(teamMember)
		if err != nil {
			t.Fatalf("Failed to create a test team member: %v", err)
		}
	}

	// Test GetTeamMembersByUserAndRole function to retrieve team members by UserID and Role
	retrievedTeamMembers, err := tmr.GetTeamMembersByUserAndRole(userID, role)
	if err != nil {
		t.Errorf("GetTeamMembersByUserAndRole returned an error: %v", err)
	}

	// Check the number of retrieved team members
	expectedCount := len(teamMembers)
	if len(retrievedTeamMembers) != expectedCount {
		t.Errorf("Expected %d team members, got %d", expectedCount, len(retrievedTeamMembers))
	}

	// Verify the correctness of the retrieved team members
	for i, retrieved := range retrievedTeamMembers {
		if retrieved.TeamID != teamMembers[i].TeamID {
			t.Errorf("Mismatch in TeamID. Expected: %d, Got: %d", teamMembers[i].TeamID, retrieved.TeamID)
		}
		if retrieved.UserID != teamMembers[i].UserID {
			t.Errorf("Mismatch in UserID. Expected: %d, Got: %d", teamMembers[i].UserID, retrieved.UserID)
		}
		// Add more comparisons for other fields as needed
	}
}

func TestTeamMemberRepository_GetTeamMembersByTeamAndRole(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.TeamMember{})

	// Create the TeamMember Repository with the test database
	tmr := NewTeamMemberRepository()
	tmr.db = db

	// Create test team members with the same TeamID and Role
	teamID := uint(1)
	role := models.MemberRole
	teamMembers := []models.TeamMember{
		{TeamID: teamID, UserID: 1, Role: role},
		{TeamID: teamID, UserID: 2, Role: role},
		{TeamID: teamID, UserID: 3, Role: role},
		{TeamID: teamID, UserID: 4, Role: models.AdminRole},
	}

	for _, teamMember := range teamMembers {
		_, err = tmr.CreateTeamMember(teamMember)
		if err != nil {
			t.Fatalf("Failed to create a test team member: %v", err)
		}
	}

	// Test GetTeamMembersByTeamAndRole function to retrieve team members by TeamID and Role
	retrievedTeamMembers, err := tmr.GetTeamMembersByTeamAndRole(teamID, role)
	if err != nil {
		t.Errorf("GetTeamMembersByTeamAndRole returned an error: %v", err)
	}

	// Check the number of retrieved team members
	expectedCount := 3
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

// test GetAdminTeamByTeamID, add both admin and non admin users to test. test this function
func TestTeamMemberRepository_GetAdminTeamByTeamID(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.TeamMember{})

	// Create the TeamMember Repository with the test database
	tmr := NewTeamMemberRepository()
	tmr.db = db

	// Create test team members with the same TeamID and Role
	teamID := uint(1)
	role := models.AdminRole
	teamMembers := []models.TeamMember{
		{TeamID: teamID, UserID: 1, Role: role},
		{TeamID: teamID, UserID: 2, Role: role},
		{TeamID: teamID, UserID: 3, Role: role},
		{TeamID: teamID, UserID: 4, Role: models.MemberRole},
	}

	for _, teamMember := range teamMembers {
		_, err = tmr.CreateTeamMember(teamMember)
		if err != nil {
			t.Fatalf("Failed to create a test team member: %v", err)
		}
	}

	// Test GetTeamMembersByTeamAndRole function to retrieve team members by TeamID and Role
	retrievedTeamMembers, err := tmr.GetAdminTeamByTeamID(teamID)
	if err != nil {
		t.Errorf("GetTeamMembersByTeamAndRole returned an error: %v", err)
	}

	// Check the number of retrieved team members
	expectedCount := 3
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

func TestTeamMemberRepository_UpdateTeamMemberRole(t *testing.T) {
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
}

func TestTeamMemberRepository_DeleteTeamMemberByUserID(t *testing.T) {
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

	// Test DeleteTeamMemberByUserID function to delete the team member by TeamID and UserID
	err = tmr.DeleteTeamMemberByUserID(createdTeamMember.UserID)
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
