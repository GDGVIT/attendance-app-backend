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
func (m *MockMeetingService) GetMeetingByID(id uint) (models.Meeting, error) {
	ret := m.ctrl.Call(m, "GetMeetingByID", id)
	ret0, _ := ret[0].(models.Meeting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartMeeting mocks the StartMeeting method.
func (m *MockMeetingService) StartMeeting(meetingID uint) (models.Meeting, error) {
	ret := m.ctrl.Call(m, "StartMeeting", meetingID)
	ret0, _ := ret[0].(models.Meeting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EndMeeting mocks the EndMeeting method.
func (m *MockMeetingService) EndMeeting(meetingID uint) (models.Meeting, error) {
	ret := m.ctrl.Call(m, "EndMeeting", meetingID)
	ret0, _ := ret[0].(models.Meeting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartAttendance mocks the StartAttendance method.
func (m *MockMeetingService) StartAttendance(meetingID uint) (models.Meeting, error) {
	ret := m.ctrl.Call(m, "StartAttendance", meetingID)
	ret0, _ := ret[0].(models.Meeting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EndAttendance mocks the EndAttendance method.
func (m *MockMeetingService) EndAttendance(meetingID uint) (models.Meeting, error) {
	ret := m.ctrl.Call(m, "EndAttendance", meetingID)
	ret0, _ := ret[0].(models.Meeting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMeetingByID mocks the DeleteMeetingByID method.
func (m *MockMeetingService) DeleteMeetingByID(meetingID uint) error {
	ret := m.ctrl.Call(m, "DeleteMeetingByID", meetingID)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkAttendanceForUserInMeeting mocks the MarkAttendanceForUserInMeeting method.
func (m *MockMeetingService) MarkAttendanceForUserInMeeting(userID, meetingID uint, attendanceTime time.Time) error {
	ret := m.ctrl.Call(m, "MarkAttendanceForUserInMeeting", userID, meetingID, attendanceTime)
	ret0, _ := ret[0].(error)
	return ret0
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

func (mr *MockMeetingServiceMockRecorder) GetMeetingByID(id uint) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeetingByID", reflect.TypeOf((*MockMeetingService)(nil).GetMeetingByID), id)
}

// StartMeeting mocks the StartMeeting method.
func (mr *MockMeetingServiceMockRecorder) StartMeeting(meetingID uint) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartMeeting", reflect.TypeOf((*MockMeetingService)(nil).StartMeeting), meetingID)
}

// EndMeeting mocks the EndMeeting method.
func (mr *MockMeetingServiceMockRecorder) EndMeeting(meetingID uint) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndMeeting", reflect.TypeOf((*MockMeetingService)(nil).EndMeeting), meetingID)
}

// StartAttendance mocks the StartAttendance method.
func (mr *MockMeetingServiceMockRecorder) StartAttendance(meetingID uint) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartAttendance", reflect.TypeOf((*MockMeetingService)(nil).StartAttendance), meetingID)
}

// EndAttendance mocks the EndAttendance method.
func (mr *MockMeetingServiceMockRecorder) EndAttendance(meetingID uint) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndAttendance", reflect.TypeOf((*MockMeetingService)(nil).EndAttendance), meetingID)
}

// DeleteMeetingByID mocks the DeleteMeetingByID method.
func (mr *MockMeetingServiceMockRecorder) DeleteMeetingByID(meetingID uint) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMeetingByID", reflect.TypeOf((*MockMeetingService)(nil).DeleteMeetingByID), meetingID)
}

// MarkAttendanceForUserInMeeting mocks the MarkAttendanceForUserInMeeting method.
func (mr *MockMeetingServiceMockRecorder) MarkAttendanceForUserInMeeting(userID, meetingID uint, attendanceTime time.Time) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkAttendanceForUserInMeeting", reflect.TypeOf((*MockMeetingService)(nil).MarkAttendanceForUserInMeeting), userID, meetingID, attendanceTime)
}
