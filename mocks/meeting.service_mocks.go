package mocks

import (
	"reflect"
	"time"

	"github.com/GDGVIT/attendance-app-backend/models"
	"github.com/golang/mock/gomock"
)

// MockMeetingService is a mock implementation of MeetingServiceInterface.
type MockMeetingService struct {
	ctrl     *gomock.Controller
	recorder *MockMeetingServiceMockRecorder
}

// NewMockMeetingService creates a new mock service.
func NewMockMeetingService(ctrl *gomock.Controller) *MockMeetingService {
	mock := &MockMeetingService{ctrl: ctrl}
	mock.recorder = &MockMeetingServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows expected calls to be set.
func (m *MockMeetingService) EXPECT() *MockMeetingServiceMockRecorder {
	return m.recorder
}

// CreateMeeting mocks the CreateMeeting method.
func (m *MockMeetingService) CreateMeeting(teamID uint, title, description, venue string, location models.Location, startTime time.Time) (models.Meeting, error) {
	ret := m.ctrl.Call(m, "CreateMeeting", teamID, title, description, venue, location, startTime)
	ret0, _ := ret[0].(models.Meeting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeetingsByTeamID mocks the GetMeetingsByTeamID method.
func (m *MockMeetingService) GetMeetingsByTeamID(teamID uint, filterBy string, orderBy string) ([]models.Meeting, error) {
	ret := m.ctrl.Call(m, "GetMeetingsByTeamID", teamID, filterBy, orderBy)
	ret0, _ := ret[0].([]models.Meeting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeetingByID mocks the GetMeetingByID method.
func (m *MockMeetingService) GetMeetingByID(id uint, teamid uint) (models.Meeting, error) {
	ret := m.ctrl.Call(m, "GetMeetingByID", id, teamid)
	ret0, _ := ret[0].(models.Meeting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartMeeting mocks the StartMeeting method.
func (m *MockMeetingService) StartMeeting(meetingID uint, teamid uint) (models.Meeting, error) {
	ret := m.ctrl.Call(m, "StartMeeting", meetingID, teamid)
	ret0, _ := ret[0].(models.Meeting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EndMeeting mocks the EndMeeting method.
func (m *MockMeetingService) EndMeeting(meetingID, teamid uint) (models.Meeting, error) {
	ret := m.ctrl.Call(m, "EndMeeting", meetingID, teamid)
	ret0, _ := ret[0].(models.Meeting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartAttendance mocks the StartAttendance method.
func (m *MockMeetingService) StartAttendance(meetingID, teamid uint) (models.Meeting, error) {
	ret := m.ctrl.Call(m, "StartAttendance", meetingID, teamid)
	ret0, _ := ret[0].(models.Meeting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EndAttendance mocks the EndAttendance method.
func (m *MockMeetingService) EndAttendance(meetingID, teamid uint) (models.Meeting, error) {
	ret := m.ctrl.Call(m, "EndAttendance", meetingID, teamid)
	ret0, _ := ret[0].(models.Meeting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMeetingByID mocks the DeleteMeetingByID method.
func (m *MockMeetingService) DeleteMeetingByID(meetingID uint, teamid uint) error {
	ret := m.ctrl.Call(m, "DeleteMeetingByID", meetingID, teamid)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkAttendanceForUserInMeeting mocks the MarkAttendanceForUserInMeeting method.
func (m *MockMeetingService) MarkAttendanceForUserInMeeting(userID, meetingID uint, attendanceTime time.Time, teamid uint) (bool, error) {
	ret := m.ctrl.Call(m, "MarkAttendanceForUserInMeeting", userID, meetingID, attendanceTime, teamid)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[0].(error)
	return ret0, ret1
}

// GetAttendanceForMeeting mocks the GetAttendanceForMeeting method.
func (m *MockMeetingService) GetAttendanceForMeeting(meetingID, teamID uint) ([]models.MeetingAttendanceListResponse, error) {
	ret := m.ctrl.Call(m, "GetAttendanceForMeeting", meetingID, teamID)
	ret0, _ := ret[0].([]models.MeetingAttendanceListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpcomingUserMeetings mocks the UpcomingUserMeetings method.
func (m *MockMeetingService) UpcomingUserMeetings(userID uint) ([]models.UserUpcomingMeetingsListResponse, error) {
	ret := m.ctrl.Call(m, "UpcomingUserMeetings", userID)
	ret0, _ := ret[0].([]models.UserUpcomingMeetingsListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFullUserAttendanceRecord mocks the GetFullUserAttendanceRecord method.
func (m *MockMeetingService) GetFullUserAttendanceRecord(userID uint) ([]models.MeetingAttendanceListResponse, error) {
	ret := m.ctrl.Call(m, "GetFullUserAttendanceRecord", userID)
	ret0, _ := ret[0].([]models.MeetingAttendanceListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MockMeetingServiceMockRecorder is a mock recorder for MockMeetingService.
type MockMeetingServiceMockRecorder struct {
	mock *MockMeetingService
}

// CreateMeeting mocks the CreateMeeting method.
func (mr *MockMeetingServiceMockRecorder) CreateMeeting(teamID uint, title, description, venue string, location models.Location, startTime time.Time) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMeeting", reflect.TypeOf((*MockMeetingService)(nil).CreateMeeting), teamID, title, description, venue, location, startTime)
}

// GetMeetingsByTeamID mocks the GetMeetingsByTeamID method.
func (mr *MockMeetingServiceMockRecorder) GetMeetingsByTeamID(teamID uint, filterBy string, orderBy string) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeetingsByTeamID", reflect.TypeOf((*MockMeetingService)(nil).GetMeetingsByTeamID), teamID, filterBy, orderBy)
}

func (mr *MockMeetingServiceMockRecorder) GetMeetingByID(id, teamid uint) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeetingByID", reflect.TypeOf((*MockMeetingService)(nil).GetMeetingByID), id, teamid)
}

// StartMeeting mocks the StartMeeting method.
func (mr *MockMeetingServiceMockRecorder) StartMeeting(meetingID, teamid uint) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartMeeting", reflect.TypeOf((*MockMeetingService)(nil).StartMeeting), meetingID, teamid)
}

// EndMeeting mocks the EndMeeting method.
func (mr *MockMeetingServiceMockRecorder) EndMeeting(meetingID, teamid uint) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndMeeting", reflect.TypeOf((*MockMeetingService)(nil).EndMeeting), meetingID, teamid)
}

// StartAttendance mocks the StartAttendance method.
func (mr *MockMeetingServiceMockRecorder) StartAttendance(meetingID, teamid uint) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartAttendance", reflect.TypeOf((*MockMeetingService)(nil).StartAttendance), meetingID, teamid)
}

// EndAttendance mocks the EndAttendance method.
func (mr *MockMeetingServiceMockRecorder) EndAttendance(meetingID, teamid uint) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndAttendance", reflect.TypeOf((*MockMeetingService)(nil).EndAttendance), meetingID, teamid)
}

// DeleteMeetingByID mocks the DeleteMeetingByID method.
func (mr *MockMeetingServiceMockRecorder) DeleteMeetingByID(meetingID, teamid uint) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMeetingByID", reflect.TypeOf((*MockMeetingService)(nil).DeleteMeetingByID), meetingID, teamid)
}

// MarkAttendanceForUserInMeeting mocks the MarkAttendanceForUserInMeeting method.
func (mr *MockMeetingServiceMockRecorder) MarkAttendanceForUserInMeeting(userID, meetingID uint, attendanceTime time.Time, teamid uint) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkAttendanceForUserInMeeting", reflect.TypeOf((*MockMeetingService)(nil).MarkAttendanceForUserInMeeting), userID, meetingID, attendanceTime, teamid)
}

// GetAttendanceForMeeting mocks the GetAttendanceForMeeting method.
func (mr *MockMeetingServiceMockRecorder) GetAttendanceForMeeting(meetingID, teamID uint) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttendanceForMeeting", reflect.TypeOf((*MockMeetingService)(nil).GetAttendanceForMeeting), meetingID, teamID)
}

// UpcomingUserMeetings mocks the UpcomingUserMeetings method.
func (mr *MockMeetingServiceMockRecorder) UpcomingUserMeetings(userID uint) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpcomingUserMeetings", reflect.TypeOf((*MockMeetingService)(nil).UpcomingUserMeetings), userID)
}

// GetFullUserAttendanceRecord mocks the GetFullUserAttendanceRecord method.
func (mr *MockMeetingServiceMockRecorder) GetFullUserAttendanceRecord(userID uint) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFullUserAttendanceRecord", reflect.TypeOf((*MockMeetingService)(nil).GetFullUserAttendanceRecord), userID)
}
