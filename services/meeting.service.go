package services

import "github.com/GDGVIT/attendance-app-backend/repository"

// Meeting service
type MeetingService struct {
	meetingRepo *repository.MeetingRepository
}

// NewMeetingService creates a new meeting service
func NewMeetingService() *MeetingService {
	meetingRepo := repository.NewMeetingRepository()
	return &MeetingService{meetingRepo}
}

// CreateMeeting creates a new meeting by calling the CreateMeeting function of the meeting repository
// Takes meetingName, meetingDescription, teamID, startTime, venue, latitude, longitude, altitude as parameters
// write the above function
func (ms *MeetingService) CreateMeeting() {

}

// GetMeetingByID retrieves a meeting by its ID
func (ms *MeetingService) GetMeetingByID() {

}

// StartMeeting starts a meeting
func (ms *MeetingService) StartMeeting() {

}

// StartAttendance starts attendance for a meeting
func (ms *MeetingService) StartAttendance() {

}

// EndAttendance ends attendance for a meeting
func (ms *MeetingService) EndAttendance() {

}

// EndMeeting ends a meeting
func (ms *MeetingService) EndMeeting() {

}

// DeleteMeetingByID deletes a meeting by its ID
func (ms *MeetingService) DeleteMeetingByID() {

}

// GetMeetingsByTeamID retrieves all meetings for a given team ID
func (ms *MeetingService) GetMeetingsByTeamID() {

}
