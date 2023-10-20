package services

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/GDGVIT/attendance-app-backend/mocks"
	"github.com/GDGVIT/attendance-app-backend/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMeetingService_CreateMeeting(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMeetingRepository(ctrl)
	service := NewMeetingService(mockRepo)

	// Mock Repository Call
	meeting := models.Meeting{
		TeamID:      1,
		Title:       "Sample Meeting",
		Description: "Description",
		Venue:       "Venue",
		Location: models.Location{
			Latitude:  10.0,
			Longitude: 20.0,
			Altitude:  30.0,
		},
		StartTime: time.Now(),
	}

	// continue
	mockRepo.EXPECT().CreateMeeting(meeting).Return(meeting, nil)

	// Call the service
	createdMeeting, err := service.CreateMeeting(meeting.TeamID, meeting.Title, meeting.Description, meeting.Venue, meeting.Location, meeting.StartTime)

	// Assert the response for the passing case
	assert.NoError(t, err)
	assert.Equal(t, meeting, createdMeeting)

	// Failing test case
	failingMeeting := models.Meeting{
		TeamID: 1,
		// Title:       "hello",
		// Description: "Description",
		Venue: "Venue",
		Location: models.Location{
			Latitude:  10.0,
			Longitude: 20.0,
			Altitude:  30.0,
		},
		StartTime: time.Now(),
	}

	// Expect the CreateMeeting method to be called with the failing meeting
	mockRepo.EXPECT().CreateMeeting(failingMeeting).Return(failingMeeting, errors.New("")).Times(1)

	// Call the service for the failing case
	failingCreatedMeeting, failingErr := service.CreateMeeting(failingMeeting.TeamID, failingMeeting.Title, failingMeeting.Description, failingMeeting.Venue, failingMeeting.Location, failingMeeting.StartTime)

	println(failingErr.Error())

	// Assert that the error is not nil and the created meeting is empty for the failing case
	assert.Error(t, failingErr)
	assert.NotEqual(t, failingMeeting, failingCreatedMeeting)

}

func TestMeetingService_GetMeetingsByTeamIDAndMeetingOver(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock MeetingRepository
	mockRepo := mocks.NewMockMeetingRepository(ctrl)

	// Create a test MeetingService with the mock repository
	service := NewMeetingService(mockRepo)

	// Define test data
	teamID := uint(1)
	meetingOver := false

	// Define the expected return value from the repository
	expectedMeetings := []models.Meeting{
		// Create test meetings as needed
		{
			TeamID:      1,
			Title:       "Sample Meeting",
			Description: "Description",
			Venue:       "Venue",
			Location: models.Location{
				Latitude:  10.0,
				Longitude: 20.0,
				Altitude:  30.0,
			},
			StartTime: time.Now().Add(time.Hour * 24),
		},
		{
			TeamID:      1,
			Title:       "Sample Meeting 2",
			Description: "Description",
			Venue:       "Venue",
			Location: models.Location{
				Latitude:  10.0,
				Longitude: 20.0,
				Altitude:  30.0,
			},
			StartTime: time.Now().Add(time.Hour * 24),
		},
	}

	// Mock the GetMeetingsByTeamIDAndMeetingOver function
	mockRepo.EXPECT().GetMeetingsByTeamIDAndMeetingOver(teamID, meetingOver).Return(expectedMeetings, nil)

	// Call the service function
	meetings, err := service.GetMeetingsByTeamID(teamID, "upcoming", "")
	if err != nil {
		t.Errorf("GetMeetingsByTeamIDAndMeetingOver returned an error: %v", err)
	}

	// Assert the result
	if !reflect.DeepEqual(meetings, expectedMeetings) {
		t.Errorf("GetMeetingsByTeamIDAndMeetingOver returned unexpected meetings: got %v, want %v", meetings, expectedMeetings)
	}
}
