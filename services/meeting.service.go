package services

import (
	"errors"
	"sort"
	"time"

	"github.com/GDGVIT/attendance-app-backend/infra/logger"
	"github.com/GDGVIT/attendance-app-backend/models"
	"github.com/GDGVIT/attendance-app-backend/repository"
	"github.com/GDGVIT/attendance-app-backend/utils/email"
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
	GetMeetingByID(id uint, teamid uint) (models.Meeting, error)
	StartMeeting(meetingID uint, teamid uint) (models.Meeting, error)
	EndMeeting(meetingID uint, teamid uint) (models.Meeting, error)
	StartAttendance(meetingID uint, teamid uint) (models.Meeting, error)
	EndAttendance(meetingID uint, teamid uint) (models.Meeting, error)
	DeleteMeetingByID(meetingID uint, teamid uint) error
	MarkAttendanceForUserInMeeting(userID, meetingID uint, attendanceTime time.Time, teamid uint) (bool, error)
	GetAttendanceForMeeting(meetingID, teamID uint) ([]models.MeetingAttendanceListResponse, error)
	UpcomingUserMeetings(userID uint) ([]models.UserUpcomingMeetingsListResponse, error)
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
		logger.Errorf("Error creating meeting: " + err.Error())
		return models.Meeting{}, err
	}

	// get name of team
	teamRepo := repository.NewTeamRepository()
	team, err := teamRepo.GetTeamByID(teamID)
	if err != nil {
		return models.Meeting{}, err
	}

	// get members of team and send email to each member
	teamMemberRepo := repository.NewTeamMemberRepository()
	teamMembers, err := teamMemberRepo.GetTeamMembersByTeamID(teamID)
	if err != nil {
		return models.Meeting{}, err
	}

	teamMemberEmails := []string{}
	userRepo := repository.NewUserRepository()
	for _, teamMember := range teamMembers {
		user, err := userRepo.GetUserByID(teamMember.UserID)
		if err != nil {
			return models.Meeting{}, err
		}
		// add user email to teamMemberEmails
		teamMemberEmails = append(teamMemberEmails, user.Email)
	}

	// send email to each team member
	email.SendMeetingNotifToTeamMembers(teamMemberEmails, team.Name, meeting.Title, meeting.StartTime)

	return createdMeeting, nil
}

// GetMeetingByID retrieves a meeting by its ID.
func (ms *MeetingService) GetMeetingByID(id uint, teamid uint) (models.Meeting, error) {
	meeting, err := ms.meetingRepo.GetMeetingByID(id)
	if err != nil {
		logger.Errorf("Error getting meeting: " + err.Error())
		return models.Meeting{}, err
	}
	// check if meeting teamid is same as teamid
	if meeting.TeamID != teamid {
		return models.Meeting{}, errors.New("meeting not found")
	}
	return meeting, nil
}

// StartMeeting starts a meeting by setting MeetingPeriod to true, if not MeetingOver.
func (ms *MeetingService) StartMeeting(meetingID uint, teamid uint) (models.Meeting, error) {
	meeting, err := ms.GetMeetingByID(meetingID, teamid)
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
func (ms *MeetingService) StartAttendance(meetingID uint, teamID uint) (models.Meeting, error) {
	meeting, err := ms.GetMeetingByID(meetingID, teamID)
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
	meeting.AttendanceOver = false

	// Update the meeting in the database
	updatedMeeting, err := ms.meetingRepo.UpdateMeeting(meeting)
	if err != nil {
		return models.Meeting{}, err
	}

	return updatedMeeting, nil
}

// EndAttendance ends attendance for a meeting by setting AttendancePeriod to false.
func (ms *MeetingService) EndAttendance(meetingID uint, teamID uint) (models.Meeting, error) {
	meeting, err := ms.GetMeetingByID(meetingID, teamID)
	if err != nil {
		return models.Meeting{}, err
	}

	// // cannot end attendance period before starting it
	// if !meeting.AttendancePeriod {
	// 	return models.Meeting{}, errors.New("attendance cannot be ended before starting it")
	// }
	// Should user be able to end meeting without taking attendance?

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
func (ms *MeetingService) EndMeeting(meetingID uint, teamID uint) (models.Meeting, error) {
	meeting, err := ms.GetMeetingByID(meetingID, teamID)
	if err != nil {
		return models.Meeting{}, err
	}

	// If attendance period is still open, close it
	meeting.AttendancePeriod = false
	meeting.AttendanceOver = true

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
func (ms *MeetingService) DeleteMeetingByID(meetingID uint, teamID uint) error {
	// A meeting can only be deleted if MeetingPeriod = false and AttendancePeriod = false and MeetingOver = false. I.e., meeting hasn't started yet.

	meeting, err := ms.GetMeetingByID(meetingID, teamID)
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

// MarkAttendaceForUserInMeeting marks attendance for a user in a meeting. Returns onTime bool.
func (ms *MeetingService) MarkAttendanceForUserInMeeting(userID, meetingID uint, attendanceTime time.Time, teamID uint) (bool, error) {
	// If meeting not started or meeting over, return error
	meeting, err := ms.GetMeetingByID(meetingID, teamID)
	if err != nil {
		return false, err
	}

	if !meeting.MeetingPeriod || meeting.MeetingOver {
		return false, errors.New("meeting not started or meeting over")
	}

	// If meeting started but attendance not started (ie, not attendance period, and not attendance ended), return error
	if !meeting.AttendancePeriod && meeting.MeetingPeriod && !meeting.AttendanceOver {
		return false, errors.New("attendance not started")
	}

	// check if attendance record for user and meeting exists. If it does, return error.
	_, err = ms.meetingRepo.GetMeetingAttendanceByUserIDAndMeetingID(userID, meetingID)
	if err == nil {
		// attendance record exists
		return false, errors.New("attendance already marked")
	}

	var onTime bool
	// if attendance period ended (but meeting period still on), mark attendance as late
	if meeting.AttendanceOver && meeting.MeetingPeriod {
		onTime = false
	} else {
		onTime = true
	}

	meetingAttendance := models.MeetingAttendance{
		UserID:             userID,
		MeetingID:          meetingID,
		AttendanceMarkedAt: attendanceTime,
		OnTime:             onTime,
	}

	if err := ms.meetingRepo.AddMeetingAttendance(meetingAttendance); err != nil {
		return false, err
	}

	return meetingAttendance.OnTime, nil
}

// GetAttendanceForMeeting retrieves attendance for a meeting.
func (ms *MeetingService) GetAttendanceForMeeting(meetingID, teamID uint) ([]models.MeetingAttendanceListResponse, error) {
	_, err := ms.GetMeetingByID(meetingID, teamID)
	if err != nil {
		return []models.MeetingAttendanceListResponse{}, err
	}
	attendance, err := ms.meetingRepo.GetMeetingAttendanceByMeetingID(meetingID)
	if err != nil {
		return nil, err
	}

	userRepo := repository.NewUserRepository()
	// Get user details for each attendance record, and make array of MeetingAttendanceResponse
	var attendanceResponse []models.MeetingAttendanceListResponse
	for _, attendanceRecord := range attendance {
		user, err := userRepo.GetUserByID(attendanceRecord.UserID)
		if err != nil {
			return nil, err
		}
		attendanceResponse = append(attendanceResponse, models.MeetingAttendanceListResponse{
			ID:                 attendanceRecord.ID,
			MeetingID:          attendanceRecord.MeetingID,
			AttendanceMarkedAt: attendanceRecord.AttendanceMarkedAt,
			OnTime:             attendanceRecord.OnTime,
			User:               user,
		})
	}

	return attendanceResponse, nil
}

// UpcomingUserMeetings retrieves all upcoming meetings for a user.
func (ms *MeetingService) UpcomingUserMeetings(userID uint) ([]models.UserUpcomingMeetingsListResponse, error) {
	// Get all teams for the user
	teamMemberRepo := repository.NewTeamMemberRepository()
	teamRepo := repository.NewTeamRepository()
	teams, err := teamMemberRepo.GetTeamMembersByUserID(userID)
	if err != nil {
		return nil, err
	}

	// Get all meetings for each team, overall ascending order by start time
	var meetings []models.UserUpcomingMeetingsListResponse
	for _, team := range teams {
		teamMeetings, err := ms.GetMeetingsByTeamID(team.TeamID, "upcoming", "asc")
		if err != nil {
			return nil, err
		}
		teamDetails, err := teamRepo.GetTeamByID(team.TeamID)
		if err != nil {
			return nil, err
		}
		for _, meeting := range teamMeetings {
			meetings = append(meetings, models.UserUpcomingMeetingsListResponse{
				Meeting: meeting,
				Team:    teamDetails,
			})
		}
	}

	// order meetings by start time
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].Meeting.StartTime.Before(meetings[j].Meeting.StartTime)
	})

	if len(meetings) <= 0 {
		meetings = []models.UserUpcomingMeetingsListResponse{}
	}

	return meetings, nil
}

// GetMeetingStatsByMeetingID retrieves meeting stats for a given meeting ID. Stats: total attendance, on time attendance, late attendance.
