package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
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

func createTestMeeting(startTime time.Time) models.Meeting {
	return models.Meeting{
		TeamID:      1,
		Title:       "Test Meeting",
		Description: "Test Description",
		Venue:       "Test Venue",
		Location: models.Location{
			Latitude:  12.345678,
			Longitude: 98.765432,
			Altitude:  0,
		},
		StartTime:        startTime,
		MeetingPeriod:    false,
		AttendancePeriod: false,
		MeetingOver:      false,
		AttendanceOver:   false,
	}
}

func parseMeetingResponse(w *httptest.ResponseRecorder) *models.Meeting {
	var responseMeeting models.Meeting
	_ = json.NewDecoder(w.Body).Decode(&responseMeeting)
	return &responseMeeting
}

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
	r.GET("/teams/:teamID/meetings", meetingController.GetMeetingsByTeamID)

	// Set up route parameters
	teamID := "1"
	filterBy := "upcoming"
	orderBy := "asc"

	// Mock the service's GetMeetingsByTeamIDAndMeetingOver function
	mockService.EXPECT().GetMeetingsByTeamID(uint(1), "upcoming", "asc").Return([]models.Meeting{}, nil)

	// Perform the request
	req, _ := http.NewRequest("GET", "/teams/"+teamID+"/meetings?filterBy="+filterBy+"&orderBy="+orderBy, nil)

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)

	// You can also parse the response body to validate the result
	var responseMeetings []models.Meeting
	err := json.NewDecoder(w.Body).Decode(&responseMeetings)
	assert.NoError(t, err)
	assert.Equal(t, []models.Meeting{}, responseMeetings)
}

// test getmeetingdetails
func TestMeetingController_GetMeetingDetails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockMeetingService(ctrl)

	r := gin.Default()

	meetingController := NewMeetingController(mockService)

	r.GET("/team/:teamID/meetings/:meetingID", meetingController.GetMeetingDetails)

	// Create a test meeting input
	now := time.Now()
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
		StartTime:        now,
		MeetingPeriod:    true,
		AttendancePeriod: false,
		MeetingOver:      false,
		AttendanceOver:   false,
	}

	// Mock the service's GetMeetingByID function
	mockService.EXPECT().GetMeetingByID(uint(1), uint(1)).Return(meeting, nil)

	// Perform the request
	req, _ := http.NewRequest("GET", "/team/1/meetings/1", nil)

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)

	// You can also parse the response body to validate the result
	var responseMeeting models.Meeting
	err := json.NewDecoder(w.Body).Decode(&responseMeeting)
	assert.NoError(t, err)

	// deep equality check of meeting and responsemeeting structs
	assert.True(t, meeting.StartTime.Equal(responseMeeting.StartTime))
	assert.Equal(t, meeting.Title, responseMeeting.Title)
}

func TestMeetingController_StartMeeting(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockMeetingService(ctrl)

	r := gin.Default()
	meetingController := NewMeetingController(mockService)

	r.PUT("/team/:teamID/meetings/:meetingID/start", meetingController.StartMeeting)

	// Helper function to send a request and check the response
	sendRequest := func(method, path string) (*httptest.ResponseRecorder, *models.Meeting) {
		req, _ := http.NewRequest(method, path, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		return w, parseMeetingResponse(w)
	}

	// Test case 1: Valid request
	now := time.Now().Add(time.Hour)
	meeting := createTestMeeting(now)
	meeting.MeetingPeriod = true
	mockService.EXPECT().StartMeeting(uint(1), uint(1)).Return(meeting, nil)
	w, responseMeeting := sendRequest("PUT", "/team/1/meetings/1/start")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, meeting.Title, responseMeeting.Title)
	assert.True(t, responseMeeting.MeetingPeriod)

	// Test case 2: Invalid meeting ID
	w, _ = sendRequest("PUT", "/team/1/meetings/jj/start")
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Test case 3: Meeting already ended
	mockService.EXPECT().StartMeeting(uint(1), uint(1)).Return(models.Meeting{}, errors.New("some error"))
	w, _ = sendRequest("PUT", "/team/1/meetings/1/start")
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// test StartAttendance
func TestMeetingController_StartAttendance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockMeetingService(ctrl)

	r := gin.Default()
	meetingController := NewMeetingController(mockService)

	r.PUT("/team/:teamID/meetings/:meetingID/attendance/start", meetingController.StartAttendance)

	// Helper function to send a request and check the response
	sendRequest := func(method, path string) (*httptest.ResponseRecorder, *models.Meeting) {
		req, _ := http.NewRequest(method, path, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		return w, parseMeetingResponse(w)
	}

	// Test case 1: Valid request
	now := time.Now().Add(time.Hour)
	meeting := createTestMeeting(now)
	meeting.AttendancePeriod = true
	mockService.EXPECT().StartAttendance(uint(1), uint(1)).Return(meeting, nil)
	w, responseMeeting := sendRequest("PUT", "/team/1/meetings/1/attendance/start")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, meeting.Title, responseMeeting.Title)
	assert.True(t, responseMeeting.AttendancePeriod)

	// Test case 2: Invalid meeting ID
	w, _ = sendRequest("PUT", "/team/1/meetings/jj/attendance/start")
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Test case 3: Meeting already ended
	mockService.EXPECT().StartAttendance(uint(1), uint(1)).Return(models.Meeting{}, errors.New("some error"))
	w, _ = sendRequest("PUT", "/team/1/meetings/1/attendance/start")
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// test EndAttendance
func TestMeetingController_EndAttendance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockMeetingService(ctrl)

	r := gin.Default()
	meetingController := NewMeetingController(mockService)

	r.PUT("/team/:teamID/meetings/:meetingID/attendance/end", meetingController.EndAttendance)

	// Helper function to send a request and check the response
	sendRequest := func(method, path string) (*httptest.ResponseRecorder, *models.Meeting) {
		req, _ := http.NewRequest(method, path, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		return w, parseMeetingResponse(w)
	}

	// Test case 1: Valid request
	now := time.Now().Add(time.Hour)
	meeting := createTestMeeting(now)
	meeting.AttendanceOver = true
	mockService.EXPECT().EndAttendance(uint(1), uint(1)).Return(meeting, nil)
	w, responseMeeting := sendRequest("PUT", "/team/1/meetings/1/attendance/end")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, meeting.Title, responseMeeting.Title)
	assert.True(t, responseMeeting.AttendanceOver)

	// Test case 2: Invalid meeting ID
	w, _ = sendRequest("PUT", "/team/1/meetings/jj/attendance/end")
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Test case 3: Meeting already ended
	mockService.EXPECT().EndAttendance(uint(1), uint(1)).Return(models.Meeting{}, errors.New("some error"))
	w, _ = sendRequest("PUT", "/team/1/meetings/1/attendance/end")
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestMeetingController_EndMeeting(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockMeetingService(ctrl)

	r := gin.Default()
	meetingController := NewMeetingController(mockService)

	r.PUT("/team/:teamID/meetings/:meetingID/end", meetingController.EndMeeting)

	// Helper function to send a request and check the response
	sendRequest := func(method, path string) (*httptest.ResponseRecorder, *models.Meeting) {
		req, _ := http.NewRequest(method, path, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		return w, parseMeetingResponse(w)
	}

	// Test case 1: Valid request
	now := time.Now().Add(time.Hour)
	meeting := createTestMeeting(now)
	meeting.MeetingOver = true
	mockService.EXPECT().EndMeeting(uint(1), uint(1)).Return(meeting, nil)
	w, responseMeeting := sendRequest("PUT", "/team/1/meetings/1/end")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, meeting.Title, responseMeeting.Title)
	assert.True(t, responseMeeting.MeetingOver)

	// Test case 2: Invalid meeting ID
	w, _ = sendRequest("PUT", "/team/1/meetings/jj/end")
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Test case 3: Meeting already ended
	mockService.EXPECT().EndMeeting(uint(1), uint(1)).Return(models.Meeting{}, errors.New("some error"))
	w, _ = sendRequest("PUT", "/team/1/meetings/1/end")
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// test DeleteMeeting
func TestMeetingController_DeleteMeeting(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockMeetingService(ctrl)

	r := gin.Default()
	meetingController := NewMeetingController(mockService)

	r.DELETE("/team/:teamID/meetings/:meetingID", meetingController.DeleteMeetingByID)

	// Helper function to send a request and check the response
	sendRequest := func(method, path string) (*httptest.ResponseRecorder, *models.Meeting) {
		req, _ := http.NewRequest(method, path, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		return w, parseMeetingResponse(w)
	}

	// Test case 1: Valid request
	mockService.EXPECT().DeleteMeetingByID(uint(1), uint(1)).Return(nil)
	w, _ := sendRequest("DELETE", "/team/1/meetings/1")
	assert.Equal(t, http.StatusOK, w.Code)

	// Test case 2: Invalid meeting ID
	w, _ = sendRequest("DELETE", "/team/1/meetings/jj")
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Test case 3: Meeting already ended
	mockService.EXPECT().DeleteMeetingByID(uint(1), uint(1)).Return(errors.New("some error"))
	w, _ = sendRequest("DELETE", "/team/1/meetings/1")
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// test GetAttendanceForMeeting
func TestMeetingController_GetAttendanceForMeeting(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockMeetingService(ctrl)

	r := gin.Default()
	meetingController := NewMeetingController(mockService)

	r.GET("/team/:teamID/meetings/:meetingID/attendance", meetingController.GetAttendanceForMeeting)

	// parseMeetingAttendanceResponse parses the response body of GetAttendanceForMeeting
	parseMeetingAttendanceResponse := func(w *httptest.ResponseRecorder) []models.MeetingAttendanceListResponse {
		var response []models.MeetingAttendanceListResponse
		_ = json.NewDecoder(w.Body).Decode(&response)
		return response
	}

	// Helper function to send a request and check the response
	sendRequest := func(method, path string) (*httptest.ResponseRecorder, []models.MeetingAttendanceListResponse) {
		req, _ := http.NewRequest(method, path, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		return w, parseMeetingAttendanceResponse(w)
	}

	// Test case 1: Valid request
	mockService.EXPECT().GetAttendanceForMeeting(uint(1), uint(1)).Return([]models.MeetingAttendanceListResponse{}, nil)
	w, _ := sendRequest("GET", "/team/1/meetings/1/attendance")
	assert.Equal(t, http.StatusOK, w.Code)

	// Test case 2: Invalid meeting ID
	w, _ = sendRequest("GET", "/team/1/meetings/jj/attendance")
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Test case 3: Meeting already ended
	mockService.EXPECT().GetAttendanceForMeeting(uint(1), uint(1)).Return([]models.MeetingAttendanceListResponse{}, errors.New("some error"))
	w, _ = sendRequest("GET", "/team/1/meetings/1/attendance")
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// test GetUserAttendanceRecords
func TestMeetingController_GetUserAttendanceRecords(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockMeetingService(ctrl)

	r := gin.Default()
	meetingController := NewMeetingController(mockService)

	r.GET("/user/me/attendance", func(c *gin.Context) {
		// Set up a user in the context
		user := &models.User{}
		user.ID = 1
		c.Set("user", user)

		// Call the controller function
		meetingController.GetUserAttendanceRecords(c)
	})

	// parseMeetingAttendanceResponse parses the response body of GetUserAttendanceRecords
	parseMeetingAttendanceResponse := func(w *httptest.ResponseRecorder) []models.MeetingAttendanceListResponse {
		var response []models.MeetingAttendanceListResponse
		_ = json.NewDecoder(w.Body).Decode(&response)
		return response
	}

	// Helper function to send a request and check the response
	sendRequest := func(method, path string) (*httptest.ResponseRecorder, []models.MeetingAttendanceListResponse) {
		req, _ := http.NewRequest(method, path, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		return w, parseMeetingAttendanceResponse(w)
	}

	mockService.EXPECT().GetFullUserAttendanceRecord(uint(1)).Return([]models.MeetingAttendanceListResponse{}, nil)
	w, _ := sendRequest("GET", "/user/me/attendance")
	assert.Equal(t, http.StatusOK, w.Code)
}
