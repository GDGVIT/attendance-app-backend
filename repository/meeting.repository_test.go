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
		StartTime: time.Now().Add(time.Hour * 24),
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
		StartTime: time.Now().Add(time.Hour * 24),
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

func TestMeetingRepository_GetMeetingByID(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.Meeting{})

	// Create the Meeting Repository with the test database
	mr := NewMeetingRepository()
	mr.db = db

	// Create a test meeting
	meeting := models.Meeting{
		Title:       "Test Meeting",
		Description: "Test Meeting Description",
		TeamID:      1,
		StartTime:   time.Now().Add(time.Hour * 24),
		Venue:       "Test Venue",
		Location: models.Location{
			Latitude:  12.345678,
			Longitude: 98.765432,
			Altitude:  0,
		},
	}

	// Test GetMeetingByID function
	createdMeeting, err := mr.CreateMeeting(meeting)
	if err != nil {
		t.Errorf("CreateMeeting returned an error: %v", err)
	}

	// Test GetMeetingByID with a valid ID
	retrievedMeeting, err := mr.GetMeetingByID(createdMeeting.ID)
	if err != nil {
		t.Errorf("GetMeetingByID returned an error: %v", err)
	}

	// Check that retrievedMeeting matches the createdMeeting
	if retrievedMeeting.ID != createdMeeting.ID {
		t.Errorf("Expected ID to be %v, got %v", createdMeeting.ID, retrievedMeeting.ID)
	}

	// Test GetMeetingByID with an invalid ID
	_, err = mr.GetMeetingByID(999) // Non-existent ID
	if err == nil {
		t.Errorf("GetMeetingByID should have returned an error for a non-existent meeting")
	}
}

func TestMeetingRepository_UpdateMeeting(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.Meeting{})

	// Create the Meeting Repository with the test database
	mr := NewMeetingRepository()
	mr.db = db

	// Create a test meeting
	meeting := models.Meeting{
		Title:       "Test Meeting",
		Description: "Test Meeting Description",
		TeamID:      1,
		StartTime:   time.Now().Add(time.Hour * 24),
		Venue:       "Test Venue",
		Location: models.Location{
			Latitude:  12.345678,
			Longitude: 98.765432,
			Altitude:  0,
		},
	}

	// Test UpdateMeeting function
	createdMeeting, err := mr.CreateMeeting(meeting)
	if err != nil {
		t.Errorf("CreateMeeting returned an error: %v", err)
	}

	// Update the meeting
	createdMeeting.Title = "Updated Meeting Title"
	updatedMeeting, err := mr.UpdateMeeting(createdMeeting)
	if err != nil {
		t.Errorf("UpdateMeeting returned an error: %v", err)
	}

	// Check that updatedMeeting matches the changes
	if updatedMeeting.Title != "Updated Meeting Title" {
		t.Errorf("Expected Title to be 'Updated Meeting Title', got %v", updatedMeeting.Title)
	}
}

func TestMeetingRepository_DeleteMeetingByID(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.Meeting{})

	// Create the Meeting Repository with the test database
	mr := NewMeetingRepository()
	mr.db = db

	// Create a test meeting
	meeting := models.Meeting{
		Title:       "Test Meeting",
		Description: "Test Meeting Description",
		TeamID:      1,
		StartTime:   time.Now().Add(time.Hour * 24),
		Venue:       "Test Venue",
		Location: models.Location{
			Latitude:  12.345678,
			Longitude: 98.765432,
			Altitude:  0,
		},
	}

	// Test DeleteMeetingByID function
	createdMeeting, err := mr.CreateMeeting(meeting)
	if err != nil {
		t.Errorf("CreateMeeting returned an error: %v", err)
	}

	// Test DeleteMeetingByID with a valid ID
	err = mr.DeleteMeetingByID(createdMeeting.ID)
	if err != nil {
		t.Errorf("DeleteMeetingByID returned an error: %v", err)
	}
}

func TestMeetingRepository_AddMeetingAttendance(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.Meeting{}, &models.MeetingAttendance{})

	// Create the Meeting Repository with the test database
	mr := NewMeetingRepository()
	mr.db = db

	// Create a test meeting
	meeting := models.Meeting{
		Title:       "Test Meeting",
		Description: "Test Meeting Description",
		TeamID:      1,
		StartTime:   time.Now().Add(time.Hour * 24),
		Venue:       "Test Venue",
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

	// Create a test meeting attendance
	meetingAttendance := models.MeetingAttendance{
		MeetingID:          createdMeeting.ID,
		UserID:             1,
		AttendanceMarkedAt: time.Now(),
		OnTime:             true,
	}

	// Test AddMeetingAttendance function
	err = mr.AddMeetingAttendance(meetingAttendance)
	if err != nil {
		t.Errorf("AddMeetingAttendance returned an error: %v", err)
	}

	// Test AddMeetingAttendance with an invalid meeting attendance (null value for time.Time)
	meetingAttendance.AttendanceMarkedAt = time.Time{}
	err = mr.AddMeetingAttendance(meetingAttendance)
	if err == nil {
		t.Errorf("AddMeetingAttendance should have returned an error for a non-existent meeting")
	}
}

func TestMeetingRepository_GetMeetingAttendanceByMeetingID(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.Meeting{}, &models.MeetingAttendance{})

	// Create the Meeting Repository with the test database
	mr := NewMeetingRepository()
	mr.db = db

	// Create a test meeting
	meeting := models.Meeting{
		Title:       "Test Meeting",
		Description: "Test Meeting Description",
		TeamID:      1,
		StartTime:   time.Now().Add(time.Hour * 24),
		Venue:       "Test Venue",
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

	// Create a test meeting attendance
	meetingAttendance := models.MeetingAttendance{
		MeetingID:          createdMeeting.ID,
		UserID:             1,
		AttendanceMarkedAt: time.Now(),
		OnTime:             true,
	}

	// Test AddMeetingAttendance function
	err = mr.AddMeetingAttendance(meetingAttendance)
	if err != nil {
		t.Errorf("AddMeetingAttendance returned an error: %v", err)
	}

	// Test GetMeetingAttendanceByMeetingID with a valid meeting ID
	_, err = mr.GetMeetingAttendanceByMeetingID(createdMeeting.ID)
	if err != nil {
		t.Errorf("GetMeetingAttendanceByMeetingID returned an error: %v", err)
	}

	// Test GetMeetingAttendanceByMeetingID with an invalid meeting ID
	atts, _ := mr.GetMeetingAttendanceByMeetingID(999) // Non-existent meeting ID
	if len(atts) != 0 {
		t.Errorf("GetMeetingAttendanceByMeetingID should have returned an empty slice")
	}
}

func TestMeetingRepository_GetMeetingAttendanceByMeetingIDAndOnTime(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.Meeting{}, &models.MeetingAttendance{})

	// Create the Meeting Repository with the test database
	mr := NewMeetingRepository()
	mr.db = db

	// Create a test meeting
	meeting := models.Meeting{
		Title:       "Test Meeting",
		Description: "Test Meeting Description",
		TeamID:      1,
		StartTime:   time.Now().Add(time.Hour * 24),
		Venue:       "Test Venue",
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

	// Create a test meeting attendance
	meetingAttendance := models.MeetingAttendance{
		MeetingID:          createdMeeting.ID,
		UserID:             1,
		AttendanceMarkedAt: time.Now(),
		OnTime:             true,
	}

	// Test AddMeetingAttendance function
	err = mr.AddMeetingAttendance(meetingAttendance)
	if err != nil {
		t.Errorf("AddMeetingAttendance returned an error: %v", err)
	}

	// Test GetMeetingAttendanceByMeetingIDAndOnTime with a valid meeting ID and onTime
	_, err = mr.GetMeetingAttendanceByMeetingIDAndOnTime(createdMeeting.ID, true)
	if err != nil {
		t.Errorf("GetMeetingAttendanceByMeetingIDAndOnTime returned an error: %v", err)
	}

	// Test GetMeetingAttendanceByMeetingIDAndOnTime with an invalid meeting ID
	atts, _ := mr.GetMeetingAttendanceByMeetingIDAndOnTime(999, true) // Non-existent meeting ID
	if len(atts) != 0 {
		t.Errorf("GetMeetingAttendanceByMeetingIDAndOnTime should have returned an empty slice")
	}
}

// test GetMeetingAttendanceByUserIDAndMeetingID
func TestMeetingRepository_GetMeetingAttendanceByUserIDAndMeetingID(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.Meeting{}, &models.MeetingAttendance{})

	// Create the Meeting Repository with the test database
	mr := NewMeetingRepository()
	mr.db = db

	// Create a test meeting
	meeting := models.Meeting{
		Title:       "Test Meeting",
		Description: "Test Meeting Description",
		TeamID:      1,
		StartTime:   time.Now().Add(time.Hour * 24),
		Venue:       "Test Venue",
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

	// Create a test meeting attendance
	meetingAttendance := models.MeetingAttendance{
		MeetingID:          createdMeeting.ID,
		UserID:             1,
		AttendanceMarkedAt: time.Now(),
		OnTime:             true,
	}

	// Test AddMeetingAttendance function
	err = mr.AddMeetingAttendance(meetingAttendance)
	if err != nil {
		t.Errorf("AddMeetingAttendance returned an error: %v", err)
	}

	// Test GetMeetingAttendanceByUserIDAndMeetingID with a valid meeting ID and onTime
	_, err = mr.GetMeetingAttendanceByUserIDAndMeetingID(1, createdMeeting.ID)
	if err != nil {
		t.Errorf("GetMeetingAttendanceByUserIDAndMeetingID returned an error: %v", err)
	}

	// Test GetMeetingAttendanceByUserIDAndMeetingID with an invalid meeting ID
	_, err = mr.GetMeetingAttendanceByUserIDAndMeetingID(1, 999) // Non-existent meeting ID
	if err == nil {
		t.Errorf("GetMeetingAttendanceByUserIDAndMeetingID should have returned an error")
	}
}

// test GetMeetingAttendancesByUserID
func TestMeetingRepository_GetMeetingAttendancesByUserID(t *testing.T) {
	db, err := test_utils.SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.Meeting{}, &models.MeetingAttendance{})

	// Create the Meeting Repository with the test database
	mr := NewMeetingRepository()
	mr.db = db

	// Create a test meeting
	meeting := models.Meeting{
		Title:       "Test Meeting",
		Description: "Test Meeting Description",
		TeamID:      1,
		StartTime:   time.Now().Add(time.Hour * 24),
		Venue:       "Test Venue",
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

	// Create a test meeting attendance
	meetingAttendance := models.MeetingAttendance{
		MeetingID:          createdMeeting.ID,
		UserID:             1,
		AttendanceMarkedAt: time.Now(),
		OnTime:             true,
	}

	// Another meeting and attendance
	anotherMeeting := models.Meeting{
		Title:       "Another Test Meeting",
		Description: "Another Test Meeting Description",
		TeamID:      2,
		StartTime:   time.Now().Add(time.Hour * 24),
		Venue:       "Another Test Venue",
		Location: models.Location{
			Latitude:  12.345678,
			Longitude: 98.765432,
			Altitude:  0,
		},
	}
	createdAnotherMeeting, err := mr.CreateMeeting(anotherMeeting)
	if err != nil {
		t.Errorf("CreateMeeting returned an error: %v", err)
	}
	anotherMeetingAttendance := models.MeetingAttendance{
		MeetingID:          createdAnotherMeeting.ID,
		UserID:             1,
		AttendanceMarkedAt: time.Now(),
		OnTime:             true,
	}
	err = mr.AddMeetingAttendance(anotherMeetingAttendance)
	if err != nil {
		t.Errorf("AddMeetingAttendance returned an error: %v", err)
	}

	// Test AddMeetingAttendance function
	err = mr.AddMeetingAttendance(meetingAttendance)
	if err != nil {
		t.Errorf("AddMeetingAttendance returned an error: %v", err)
	}

	// Test GetMeetingAttendancesByUserID with a valid user ID
	_, err = mr.GetMeetingAttendancesByUserID(1)
	if err != nil {
		t.Errorf("GetMeetingAttendancesByUserID returned an error: %v", err)
	}
	// Test GetMeetingAttendancesByUserID returned correct number of attendances
	attendances, _ := mr.GetMeetingAttendancesByUserID(1)
	if len(attendances) != 2 {
		t.Errorf("Expected 2 attendances, got: %d", len(attendances))
	}

	// Test GetMeetingAttendancesByUserID with an invalid user ID
	attendances, _ = mr.GetMeetingAttendancesByUserID(999) // Non-existent user ID
	if len(attendances) != 0 {
		t.Errorf("GetMeetingAttendancesByUserID should have returned an empty slice")
	}
}
