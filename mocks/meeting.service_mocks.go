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
