package repository

// test meeting repository\

import (
	"testing"
	"time"

	"github.com/GDGVIT/attendance-app-backend/models"
	"github.com/GDGVIT/attendance-app-backend/utils/test_utils"
)

// TestMeetingRepository_CreateMeeting tests the CreateMeeting function.
func TestMeetingRepository_CreateMeeting(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.Meeting{})

	// Create the Meeting Repository with the test database
	mr := NewMeetingRepository()
	mr.db = db

	// Create a test meeting with Title, Description, TeamID, StartTime, Venue and Location
	meeting := models.Meeting{
		Title:       "Test Meeting",
		Description: "Test Meeting Description",
		TeamID:      1,
		// StartTime as a time.Time value
		StartTime: time.Date(2024, 1, 1, 0, 0, 5, 0, time.UTC),
		Venue:     "Test Venue",
		// Location as the models.Location struct
		Location: models.Location{
			Latitude:  12.345678,
			Longitude: 98.765432,
			Altitude:  0,
		},
	}

	// Test CreateMeeting function
	createdMeeting, err := mr.CreateMeeting(meeting)
	if err != nil {
		t.Errorf("CreateMeeting returned an error: %v", err)
	}

	// check createdMeeting is equal to meeting
	if createdMeeting.Title != meeting.Title {
		t.Errorf("Expected Title to be %v, got %v", meeting.Title, createdMeeting.Title)
	}
	if createdMeeting.Description != meeting.Description {
		t.Errorf("Expected Description to be %v, got %v", meeting.Description, createdMeeting.Description)
	}
	if createdMeeting.TeamID != meeting.TeamID {
		t.Errorf("Expected TeamID to be %v, got %v", meeting.TeamID, createdMeeting.TeamID)
	}
	if createdMeeting.StartTime != meeting.StartTime {
		t.Errorf("Expected StartTime to be %v, got %v", meeting.StartTime, createdMeeting.StartTime)
	}
	if createdMeeting.Venue != meeting.Venue {
		t.Errorf("Expected Venue to be %v, got %v", meeting.Venue, createdMeeting.Venue)
	}
	if createdMeeting.Location.Latitude != meeting.Location.Latitude {
		t.Errorf("Expected Location.Latitude to be %v, got %v", meeting.Location.Latitude, createdMeeting.Location.Latitude)
	}
	if createdMeeting.Location.Longitude != meeting.Location.Longitude {
		t.Errorf("Expected Location.Longitude to be %v, got %v", meeting.Location.Longitude, createdMeeting.Location.Longitude)
	}
	if createdMeeting.Location.Altitude != meeting.Location.Altitude {
		t.Errorf("Expected Location.Altitude to be %v, got %v", meeting.Location.Altitude, createdMeeting.Location.Altitude)
	}

	// Test CreateMeeting with invalid meeting (missing Title)
	meeting = models.Meeting{
		Description: "Test Meeting Description",
		TeamID:      1,
		// StartTime as a time.Time value
		StartTime: time.Date(2024, 1, 1, 0, 0, 5, 0, time.UTC),
		Venue:     "Test Venue",
		// Location as the models.Location struct
		Location: models.Location{
			Latitude:  12.345678,
			Longitude: 98.765432,
			Altitude:  0,
		},
	}

	// Test CreateMeeting function
	_, err = mr.CreateMeeting(meeting)
	if err == nil {
		t.Errorf("CreateMeeting should have returned an error")
	}
}
