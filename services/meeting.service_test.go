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

func TestMeetingService_StartMeeting(t *testing.T) {
	// Initialize the mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock MeetingRepository
	mockRepo := mocks.NewMockMeetingRepository(ctrl)

	// Create a new MeetingService with the mock repository
	meetingService := NewMeetingService(mockRepo)

	// TC1

	// Define a test meeting ID
	meetingID := uint(1)

	// Define a mock meeting with MeetingOver = true
	mockMeeting := models.Meeting{
		MeetingOver: true,
		TeamID:      1,
		// Add other fields as needed
	}

	// Mock the GetMeetingByID function to return the mock meeting
	mockRepo.EXPECT().GetMeetingByID(meetingID).Return(mockMeeting, nil)

	// Call the StartMeeting function
	_, err := meetingService.StartMeeting(meetingID, 1)

	// Assert that an error is returned since MeetingOver is true
	if err == nil {
		t.Errorf("StartMeeting should have returned an error, but got nil")
	}

	// TC2

	// Define another test meeting ID
	meetingID2 := uint(2)

	// Define a mock meeting with MeetingOver = false
	mockMeeting2 := models.Meeting{
		MeetingOver: false,
		TeamID:      1,
		// Add other fields as needed
	}

	// Mock the GetMeetingByID function to return the mock meeting
	mockRepo.EXPECT().GetMeetingByID(meetingID2).Return(mockMeeting2, nil)

	mockMeeting2.MeetingPeriod = true

	// Mock the UpdateMeeting function to return the updated meeting
	mockRepo.EXPECT().UpdateMeeting(mockMeeting2).Return(mockMeeting2, nil)

	// Call the StartMeeting function
	startedMeeting, err := meetingService.StartMeeting(meetingID2, 1)

	// Assert that no error is returned since MeetingOver is false
	if err != nil {
		t.Errorf("StartMeeting returned an unexpected error: %v", err)
	}

	// Assert that MeetingPeriod is true after starting the meeting
	if !startedMeeting.MeetingPeriod {
		t.Errorf("StartMeeting did not set MeetingPeriod to true")
	}
}

func TestMeetingService_StartAttendance(t *testing.T) {
	// Initialize the mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock MeetingRepository
	mockRepo := mocks.NewMockMeetingRepository(ctrl)

	// Create a new MeetingService with the mock repository
	meetingService := NewMeetingService(mockRepo)

	testCases := []struct {
		name          string
		meetingID     uint
		mockMeeting   models.Meeting
		expectedError bool
	}{
		{
			name:      "MeetingOver_is_true",
			meetingID: 1,
			mockMeeting: models.Meeting{
				MeetingOver: true,
				TeamID:      1,
				// Add other fields as needed
			},
			expectedError: true,
		},
		{
			name:      "MeetingPeriod_is_false",
			meetingID: 2,
			mockMeeting: models.Meeting{
				MeetingOver:   false,
				MeetingPeriod: false,
				TeamID:        1,
			},
			expectedError: true,
		},
		{
			name:      "Start_attendance_successfully",
			meetingID: 3,
			mockMeeting: models.Meeting{
				MeetingOver:      false,
				MeetingPeriod:    true,
				AttendancePeriod: false,
				AttendanceOver:   false,
				TeamID:           1,
			},
			expectedError: false,
		},
		{
			name:      "Restart_attendance_after_end_attendance",
			meetingID: 4,
			mockMeeting: models.Meeting{
				MeetingOver:      false,
				MeetingPeriod:    true,
				AttendancePeriod: false,
				AttendanceOver:   true,
				TeamID:           1,
			},
			expectedError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Mock the GetMeetingByID function to return the mock meeting
			mockRepo.EXPECT().GetMeetingByID(tc.meetingID).Return(tc.mockMeeting, nil)

			if !tc.expectedError {

				tc.mockMeeting.AttendancePeriod = true
				tc.mockMeeting.AttendanceOver = false

				// Mock the UpdateMeeting function to return the mock meeting
				mockRepo.EXPECT().UpdateMeeting(tc.mockMeeting).Return(tc.mockMeeting, nil)

			}

			// Call the StartMeeting function
			returnedMockMeeting, err := meetingService.StartAttendance(tc.meetingID, 1)

			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, returnedMockMeeting.AttendancePeriod, tc.mockMeeting.AttendancePeriod)
			}
		})
	}
}

func TestMeetingService_EndAttendance(t *testing.T) {
	// Initialize the mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock MeetingRepository
	mockRepo := mocks.NewMockMeetingRepository(ctrl)

	// Create a new MeetingService with the mock repository
	meetingService := NewMeetingService(mockRepo)

	testCases := []struct {
		name          string
		meetingID     uint
		mockMeeting   models.Meeting
		expectedError bool
	}{
		{
			name:      "End_Att_Success",
			meetingID: 1,
			mockMeeting: models.Meeting{
				MeetingOver:      false,
				MeetingPeriod:    true,
				AttendancePeriod: true,
				AttendanceOver:   false,
				TeamID:           1,
			},
			expectedError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Mock the GetMeetingByID function to return the mock meeting
			mockRepo.EXPECT().GetMeetingByID(tc.meetingID).Return(tc.mockMeeting, nil)

			tc.mockMeeting.AttendancePeriod = false
			tc.mockMeeting.AttendanceOver = true

			// Mock the UpdateMeeting function to return the mock meeting
			mockRepo.EXPECT().UpdateMeeting(tc.mockMeeting).Return(tc.mockMeeting, nil)

			// Call the StartMeeting function
			returnedMockMeeting, err := meetingService.EndAttendance(tc.meetingID, 1)

			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, returnedMockMeeting.AttendancePeriod, tc.mockMeeting.AttendancePeriod)
			}
		})
	}
}

func TestMeetingService_EndMeeting(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMeetingRepository(ctrl)
	meetingService := NewMeetingService(mockRepo)

	testCases := []struct {
		name              string
		meetingID         uint
		mockMeeting       models.Meeting
		mockUpdateMeeting models.Meeting
		expectedError     bool
	}{
		{
			name:      "End_meeting_successfully",
			meetingID: 1,
			mockMeeting: models.Meeting{
				MeetingPeriod:    true,
				MeetingOver:      false,
				AttendancePeriod: true,
				AttendanceOver:   false,
				TeamID:           1,
			},
			mockUpdateMeeting: models.Meeting{
				MeetingPeriod:    false,
				MeetingOver:      true,
				AttendancePeriod: false,
				AttendanceOver:   true,
				TeamID:           1,
			},
			expectedError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo.EXPECT().GetMeetingByID(tc.meetingID).Return(tc.mockMeeting, nil)
			mockRepo.EXPECT().UpdateMeeting(tc.mockUpdateMeeting).Return(tc.mockUpdateMeeting, nil)

			updatedMeeting, err := meetingService.EndMeeting(tc.meetingID, 1)

			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.mockUpdateMeeting, updatedMeeting)
			}
		})
	}
}

func TestMeetingService_DeleteMeetingByID(t *testing.T) {
	// Initialize the mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock MeetingRepository
	mockRepo := mocks.NewMockMeetingRepository(ctrl)

	// Create a new MeetingService with the mock repository
	meetingService := NewMeetingService(mockRepo)

	testCases := []struct {
		name          string
		meetingID     uint
		mockMeeting   models.Meeting
		expectedError bool
	}{
		{
			name:      "MeetingPeriod_is_true",
			meetingID: 1,
			mockMeeting: models.Meeting{
				MeetingPeriod: true,
				TeamID:        1,
			},
			expectedError: true,
		},
		{
			name:      "AttendancePeriod_is_true",
			meetingID: 2,
			mockMeeting: models.Meeting{
				AttendancePeriod: true,
				TeamID:           1,
			},
			expectedError: true,
		},
		{
			name:      "MeetingOver_is_true",
			meetingID: 3,
			mockMeeting: models.Meeting{
				MeetingOver: true,
				TeamID:      1,
			},
			expectedError: true,
		},
		{
			name:      "AttendanceOver_is_true",
			meetingID: 4,
			mockMeeting: models.Meeting{
				AttendanceOver: true,
				TeamID:         1,
			},
			expectedError: true,
		},
		{
			name:      "Delete_meeting_successfully",
			meetingID: 5,
			mockMeeting: models.Meeting{
				MeetingPeriod:    false,
				AttendancePeriod: false,
				MeetingOver:      false,
				AttendanceOver:   false,
				TeamID:           1,
			},
			expectedError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Mock the GetMeetingByID function to return the mock meeting
			mockRepo.EXPECT().GetMeetingByID(tc.meetingID).Return(tc.mockMeeting, nil)

			if !tc.expectedError {
				mockRepo.EXPECT().DeleteMeetingByID(tc.meetingID).Return(nil)
			}

			err := meetingService.DeleteMeetingByID(tc.meetingID, 1)

			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestMeetingService_MarkAttendanceForUserInMeeting(t *testing.T) {
	// Initialize the mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock MeetingRepository
	mockRepo := mocks.NewMockMeetingRepository(ctrl)

	// Create a new MeetingService with the mock repository
	meetingService := NewMeetingService(mockRepo)

	testCases := []struct {
		name             string
		userID           uint
		meetingID        uint
		attendanceTime   time.Time
		mockMeeting      models.Meeting
		expectedError    bool
		expectedOnTime   bool
		expectedNumCalls int
	}{
		{
			name:           "Meeting_not_started",
			userID:         1,
			meetingID:      1,
			attendanceTime: time.Now(),
			mockMeeting: models.Meeting{
				MeetingPeriod: false,
				MeetingOver:   false,
				TeamID:        1,
			},
			expectedError:    true,
			expectedOnTime:   false,
			expectedNumCalls: 0,
		},
		{
			name:           "Meeting_over",
			userID:         2,
			meetingID:      2,
			attendanceTime: time.Now(),
			mockMeeting: models.Meeting{
				MeetingPeriod: true,
				TeamID:        1,
				MeetingOver:   true,
			},
			expectedError:    true,
			expectedOnTime:   false,
			expectedNumCalls: 0,
		},
		{
			name:           "Attendance_not_started",
			userID:         3,
			meetingID:      3,
			attendanceTime: time.Now(),
			mockMeeting: models.Meeting{
				MeetingPeriod:    true,
				MeetingOver:      false,
				AttendancePeriod: false,
				AttendanceOver:   false,
				TeamID:           1,
			},
			expectedError:    true,
			expectedOnTime:   false,
			expectedNumCalls: 0,
		},
		{
			name:           "Mark_attendance_successfully",
			userID:         4,
			meetingID:      4,
			attendanceTime: time.Now(),
			mockMeeting: models.Meeting{
				MeetingPeriod:    true,
				MeetingOver:      false,
				AttendancePeriod: true,
				AttendanceOver:   false,
				TeamID:           1,
			},
			expectedError:    false,
			expectedOnTime:   true,
			expectedNumCalls: 1,
		},
		{
			name:           "Mark_late_attendance",
			userID:         5,
			meetingID:      5,
			attendanceTime: time.Now(),
			mockMeeting: models.Meeting{
				MeetingPeriod:    true,
				MeetingOver:      false,
				AttendancePeriod: false,
				AttendanceOver:   true,
				TeamID:           1,
			},
			expectedError:    false,
			expectedOnTime:   false,
			expectedNumCalls: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Mock the repos's GetMeetingByID function to return the mock meeting
			mockRepo.EXPECT().GetMeetingByID(tc.meetingID).Return(tc.mockMeeting, nil)

			// If the function is expected to succeed, we should create a mock MeetingAttendance.
			if !tc.expectedError {
				// mock repos get attendance by user id and meeting id
				mockRepo.EXPECT().GetMeetingAttendanceByUserIDAndMeetingID(tc.userID, tc.meetingID).Return(models.MeetingAttendance{}, errors.New("existing attendance not found"))

				mockAttendance := models.MeetingAttendance{
					MeetingID:          tc.meetingID,
					UserID:             tc.userID,
					AttendanceMarkedAt: tc.attendanceTime,
					OnTime:             tc.expectedOnTime,
				}

				// Mock the repos's AddMeetingAttendance function to return the mock MeetingAttendance
				mockRepo.EXPECT().AddMeetingAttendance(mockAttendance).Return(nil)
			}

			_, err := meetingService.MarkAttendanceForUserInMeeting(tc.userID, tc.meetingID, tc.attendanceTime, 1)

			// Assert the error based on the expected result
			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
