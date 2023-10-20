package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/GDGVIT/attendance-app-backend/mocks" // Import your mock package
	"github.com/GDGVIT/attendance-app-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMeetingController_CreateMeeting(t *testing.T) {
	// Initialize the mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock MeetingService
	mockService := mocks.NewMockMeetingService(ctrl)

	// Create a new Gin router
	r := gin.Default()

	// Initialize the meeting controller with the mock service
	meetingController := NewMeetingController(mockService)

	// Register routes for testing
	r.POST("/teams/:teamID/meetings", meetingController.CreateMeeting)

	// time variable to new years of 2022
	time := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)

	// Create a test meeting input
	meeting := models.Meeting{
		TeamID:      1,
		Title:       "Test Meeting",
		Description: "Test Description",
		Venue:       "Test Venue",
		Location: models.Location{
			Latitude:  12.345678,
			Longitude: 98.765432,
			Altitude:  0,
		},
		StartTime: time,
	}

	// Mock the service's CreateMeeting function
	mockService.EXPECT().CreateMeeting(
		uint(1),
		meeting.Title,
		meeting.Description,
		meeting.Venue,
		meeting.Location,
		meeting.StartTime,
	).Return(meeting, nil)

	// Create a test request
	reqBody, err := json.Marshal(meeting)
	assert.NoError(t, err)
	req, _ := http.NewRequest("POST", "/teams/1/meetings", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusCreated, w.Code)
	// You can also parse the response body to validate the result
	var responseMeeting models.Meeting
	err = json.NewDecoder(w.Body).Decode(&responseMeeting)
	assert.NoError(t, err)
	assert.Equal(t, meeting, responseMeeting)
}

func TestMeetingController_GetMeetingsByTeamIDAndMeetingOver(t *testing.T) {
	// Initialize the mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock MeetingService
	mockService := mocks.NewMockMeetingService(ctrl)

	// Create a new Gin router
	r := gin.Default()

	// Initialize the meeting controller with the mock service
	meetingController := NewMeetingController(mockService)

	// Register routes for testing
	r.GET("/teams/:teamID/meetings", meetingController.GetMeetingsByTeamIDAndMeetingOver)

	// Set up route parameters
	teamID := "1"
	meetingOver := "true"

	// Mock the service's GetMeetingsByTeamIDAndMeetingOver function
	mockService.EXPECT().GetMeetingsByTeamIDAndMeetingOver(uint(1), true).Return([]models.Meeting{}, nil)

	// Perform the request
	req, _ := http.NewRequest("GET", "/teams/"+teamID+"/meetings?meetingOver="+meetingOver, nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	// You can also parse the response body to validate the result
	var responseMeetings []models.Meeting
	err := json.NewDecoder(w.Body).Decode(&responseMeetings)
	assert.NoError(t, err)
	// Add more assertions based on your test case
}
