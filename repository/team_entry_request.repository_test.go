package repository

import (
	"testing"

	"github.com/GDGVIT/attendance-app-backend/models"
	"github.com/GDGVIT/attendance-app-backend/utils/test_utils"
)

func TestTeamEntryRequestRepository_CreateTeamEntryRequest(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.TeamEntryRequest{})

	repository := NewTeamEntryRequestRepository()
	repository.db = db

	request := models.TeamEntryRequest{
		TeamID: 1,
		UserID: 1,
	}

	createdRequest, err := repository.CreateTeamEntryRequest(request)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if createdRequest.ID == 0 {
		t.Errorf("Expected the request to be created with a valid ID")
	}
}

func TestTeamEntryRequestRepository_GetTeamEntryRequestsByTeamID(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.TeamEntryRequest{})

	repository := NewTeamEntryRequestRepository()
	repository.db = db

	requests := []models.TeamEntryRequest{
		{TeamID: 1, UserID: 1},
		{TeamID: 1, UserID: 2},
		{TeamID: 2, UserID: 3},
	}

	for _, request := range requests {
		repository.CreateTeamEntryRequest(request)
	}

	retrievedRequests, err := repository.GetTeamEntryRequestsByTeamID(1)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if len(retrievedRequests) != 2 {
		t.Errorf("Expected 2 requests, but got %d", len(retrievedRequests))
	}
}

func TestTeamEntryRequestRepository_UpdateTeamEntryRequestStatus(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.TeamEntryRequest{})

	repository := NewTeamEntryRequestRepository()
	repository.db = db

	request := models.TeamEntryRequest{
		TeamID: 1,
		UserID: 1,
		Status: models.TeamEntryRequestPending,
	}

	createdRequest, err := repository.CreateTeamEntryRequest(request)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	updatedRequest, err := repository.UpdateTeamEntryRequestStatus(createdRequest.ID, models.TeamEntryRequestApproved)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if updatedRequest.Status != models.TeamEntryRequestApproved {
		t.Errorf("Expected status to be %s, but got %s", models.TeamEntryRequestApproved, updatedRequest.Status)
	}
}

func TestTeamEntryRequestRepository_GetTeamEntryRequestsByUserID(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.TeamEntryRequest{})

	repository := NewTeamEntryRequestRepository()
	repository.db = db

	requests := []models.TeamEntryRequest{
		{TeamID: 1, UserID: 1},
		{TeamID: 1, UserID: 2},
		{TeamID: 2, UserID: 3},
	}

	for _, request := range requests {
		repository.CreateTeamEntryRequest(request)
	}

	retrievedRequests, err := repository.GetTeamEntryRequestsByUserID(1)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if len(retrievedRequests) != 1 {
		t.Errorf("Expected 1 request, but got %d", len(retrievedRequests))
	}
}

func TestTeamEntryRequestRepository_GetTeamEntryRequestsByUserIDAndStatus(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.TeamEntryRequest{})

	repository := NewTeamEntryRequestRepository()
	repository.db = db

	requests := []models.TeamEntryRequest{
		{TeamID: 1, UserID: 1, Status: models.TeamEntryRequestApproved},
		{TeamID: 1, UserID: 2, Status: models.TeamEntryRequestPending},
		{TeamID: 2, UserID: 3, Status: models.TeamEntryRequestApproved},
	}

	for _, request := range requests {
		repository.CreateTeamEntryRequest(request)
	}

	retrievedRequests, err := repository.GetTeamEntryRequestsByUserIDAndStatus(1, models.TeamEntryRequestApproved)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if len(retrievedRequests) != 1 {
		t.Errorf("Expected 1 request, but got %d", len(retrievedRequests))
	}
}

func TestTeamEntryRequestRepository_GetTeamEntryRequestsByTeamIDAndStatus(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.TeamEntryRequest{})

	repository := NewTeamEntryRequestRepository()
	repository.db = db

	requests := []models.TeamEntryRequest{
		{TeamID: 1, UserID: 1, Status: models.TeamEntryRequestApproved},
		{TeamID: 1, UserID: 2, Status: models.TeamEntryRequestPending},
		{TeamID: 2, UserID: 3, Status: models.TeamEntryRequestApproved},
	}

	for _, request := range requests {
		repository.CreateTeamEntryRequest(request)
	}

	retrievedRequests, err := repository.GetTeamEntryRequestsByTeamIDAndStatus(1, models.TeamEntryRequestApproved)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if len(retrievedRequests) != 1 {
		t.Errorf("Expected 1 request, but got %d", len(retrievedRequests))
	}
}
