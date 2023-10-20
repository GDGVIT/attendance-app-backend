package services

import (
	"errors"
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
