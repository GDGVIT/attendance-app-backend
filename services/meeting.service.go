package services

import (
	"errors"
	"sort"
	"time"

	"github.com/GDGVIT/attendance-app-backend/models"
	"github.com/GDGVIT/attendance-app-backend/repository"
)

// MeetingService handles business logic related to meetings.
type MeetingService struct {
	meetingRepo repository.MeetingRepositoryInterface
}

// NewMeetingService creates a new MeetingService.
func NewMeetingService(meetingRepo repository.MeetingRepositoryInterface) *MeetingService {
	return &MeetingService{meetingRepo}
}

type MeetingServiceInterface interface {
	CreateMeeting(teamID uint, title, description, venue string, location models.Location, startTime time.Time) (models.Meeting, error)
	GetMeetingsByTeamID(teamID uint, filterBy string, orderBy string) ([]models.Meeting, error)
}

// CreateMeeting creates a new meeting in the database.
func (ms *MeetingService) CreateMeeting(teamID uint, title, description, venue string, location models.Location, startTime time.Time) (models.Meeting, error) {
	meeting := models.Meeting{
		TeamID:      teamID,
		Title:       title,
		Description: description,
		Venue:       venue,
		Location:    location,
		StartTime:   startTime,
	}

	// Create the meeting in the database
	createdMeeting, err := ms.meetingRepo.CreateMeeting(meeting)
	if err != nil {
		return models.Meeting{}, err
	}

	return createdMeeting, nil
}

// GetMeetingByID retrieves a meeting by its ID.
func (ms *MeetingService) GetMeetingByID(id uint) (models.Meeting, error) {
	meeting, err := ms.meetingRepo.GetMeetingByID(id)
	if err != nil {
		return models.Meeting{}, err
	}
	return meeting, nil
}

// StartMeeting starts a meeting by setting MeetingPeriod to true, if not MeetingOver.
func (ms *MeetingService) StartMeeting(meetingID uint) (models.Meeting, error) {
	meeting, err := ms.GetMeetingByID(meetingID)
	if err != nil {
		return models.Meeting{}, err
	}

	if meeting.MeetingOver {
		return models.Meeting{}, errors.New("meeting cannot be started after it has ended once")
	}

	meeting.MeetingPeriod = true

	// Update the meeting in the database
	updatedMeeting, err := ms.meetingRepo.UpdateMeeting(meeting)
	if err != nil {
		return models.Meeting{}, err
	}

	return updatedMeeting, nil
}

// StartAttendance starts attendance for a meeting by setting AttendancePeriod to true, if meeting in progress, or if not meeting over.
func (ms *MeetingService) StartAttendance(meetingID uint) (models.Meeting, error) {
	meeting, err := ms.GetMeetingByID(meetingID)
	if err != nil {
		return models.Meeting{}, err
	}

	if meeting.MeetingOver {
		return models.Meeting{}, errors.New("attendance cannot be started after meeting has ended")
	}

	if !meeting.MeetingPeriod {
		return models.Meeting{}, errors.New("attendance cannot be started before meeting has started")
	}

	meeting.AttendancePeriod = true

	// Update the meeting in the database
	updatedMeeting, err := ms.meetingRepo.UpdateMeeting(meeting)
	if err != nil {
		return models.Meeting{}, err
	}

	return updatedMeeting, nil
}

// EndAttendance ends attendance for a meeting by setting AttendancePeriod to false.
func (ms *MeetingService) EndAttendance(meetingID uint) (models.Meeting, error) {
	meeting, err := ms.GetMeetingByID(meetingID)
	if err != nil {
		return models.Meeting{}, err
	}

	// // cannot end attendance period before starting it
	// if !meeting.AttendancePeriod {
	// 	return models.Meeting{}, errors.New("attendance cannot be ended before starting it")
	// }

	meeting.AttendancePeriod = false
	meeting.AttendanceOver = true

	// Update the meeting in the database
	updatedMeeting, err := ms.meetingRepo.UpdateMeeting(meeting)
	if err != nil {
		return models.Meeting{}, err
	}

	return updatedMeeting, nil
}

// EndMeeting ends a meeting by setting MeetingOver to true.
func (ms *MeetingService) EndMeeting(meetingID uint) (models.Meeting, error) {
	meeting, err := ms.GetMeetingByID(meetingID)
	if err != nil {
		return models.Meeting{}, err
	}

	// If attendance period is still open, close it
	ms.EndAttendance(meetingID)

	meeting.MeetingPeriod = false
	meeting.MeetingOver = true

	// Update the meeting in the database
	updatedMeeting, err := ms.meetingRepo.UpdateMeeting(meeting)
	if err != nil {
		return models.Meeting{}, err
	}

	return updatedMeeting, nil
}

// DeleteMeetingByID deletes a meeting by its ID.
func (ms *MeetingService) DeleteMeetingByID(meetingID uint) error {
	// A meeting can only be deleted if MeetingPeriod = false and AttendancePeriod = false and MeetingOver = false. I.e., meeting hasn't started yet.

	meeting, err := ms.GetMeetingByID(meetingID)
	if err != nil {
		return err
	}

	if meeting.MeetingPeriod || meeting.AttendancePeriod || meeting.MeetingOver || meeting.AttendanceOver {
		return errors.New("meeting cannot be deleted after it has started or finished")
	}

	return ms.meetingRepo.DeleteMeetingByID(meetingID)
}

// GetMeetingsByTeamID retrieves meetings for a team based on filters.
func (ms *MeetingService) GetMeetingsByTeamID(teamID uint, filterBy string, orderBy string) ([]models.Meeting, error) {
	var meetings []models.Meeting
	var err error

	// Use a switch statement to handle different filtering options
	switch filterBy {
	case "all":
		meetings, err = ms.meetingRepo.GetMeetingsByTeamID(teamID)
	case "upcoming":
		meetings, err = ms.meetingRepo.GetMeetingsByTeamIDAndMeetingOver(teamID, false)
	case "past":
		meetings, err = ms.meetingRepo.GetMeetingsByTeamIDAndMeetingOver(teamID, true)
	default:
		return nil, errors.New("invalid filterBy value")
	}

	if err != nil {
		return nil, err
	}

	// Implement sorting based on the `orderBy` parameter
	switch orderBy {
	case "asc":
		// Sort meetings in ascending order by startTime
		sort.Slice(meetings, func(i, j int) bool {
			return meetings[i].StartTime.Before(meetings[j].StartTime)
		})
	case "desc":
		// Sort meetings in descending order by start date
		sort.Slice(meetings, func(i, j int) bool {
			return meetings[i].StartTime.After(meetings[j].StartTime)
		})
	}

	return meetings, nil
}

// MarkAttendaceForUserInMeeting marks attendance for a user in a meeting.
func (ms *MeetingService) MarkAttendaceForUserInMeeting(userID, meetingID uint) error {
	// If meeting not started or meeting over, return error
	meeting, err := ms.GetMeetingByID(meetingID)
	if err != nil {
		return err
	}

	if !meeting.MeetingPeriod || meeting.MeetingOver {
		return errors.New("meeting not started or meeting over")
	}

	// If meeting started but attendance not started (ie, not attendance period, and not attendance ended), return error
	if !meeting.AttendancePeriod && meeting.MeetingPeriod && !meeting.AttendanceOver {
		return errors.New("attendance not started")
	}

	meetingAttendance := models.MeetingAttendance{
		UserID:             userID,
		MeetingID:          meetingID,
		AttendanceMarkedAt: time.Now(),
	}

	// if attendance period ended (but meeting period still on), mark attendance as late
	if meeting.AttendanceOver && meeting.MeetingPeriod {
		meetingAttendance.OnTime = false
	}

	if err := ms.meetingRepo.AddMeetingAttendance(meetingAttendance); err != nil {
		return err
	}

	return nil
}

// GetMeetingStatsByMeetingID retrieves meeting stats for a given meeting ID. Stats: total attendance, on time attendance, late attendance.
