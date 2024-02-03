package mocks

import (
	"reflect"

	"github.com/GDGVIT/attendance-app-backend/models"
	"github.com/golang/mock/gomock"
)

// MockMeetingRepository is a mock implementation of MeetingRepositoryInterface.
type MockMeetingRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMeetingRepositoryMockRecorder
}

// NewMockMeetingRepository creates a new mock repository.
func NewMockMeetingRepository(ctrl *gomock.Controller) *MockMeetingRepository {
	mock := &MockMeetingRepository{ctrl: ctrl}
	mock.recorder = &MockMeetingRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows expected calls to be set.
func (m *MockMeetingRepository) EXPECT() *MockMeetingRepositoryMockRecorder {
	return m.recorder
}

// CreateMeeting mocks the CreateMeeting method.
func (m *MockMeetingRepository) CreateMeeting(meeting models.Meeting) (models.Meeting, error) {
	ret := m.ctrl.Call(m, "CreateMeeting", meeting)
	ret0, _ := ret[0].(models.Meeting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeetingByID mocks the GetMeetingByID method.
func (m *MockMeetingRepository) GetMeetingByID(id uint) (models.Meeting, error) {
	ret := m.ctrl.Call(m, "GetMeetingByID", id)
	ret0, _ := ret[0].(models.Meeting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMeeting mocks the UpdateMeeting method.
func (m *MockMeetingRepository) UpdateMeeting(meeting models.Meeting) (models.Meeting, error) {
	ret := m.ctrl.Call(m, "UpdateMeeting", meeting)
	ret0, _ := ret[0].(models.Meeting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMeetingByID mocks the DeleteMeetingByID method.
func (m *MockMeetingRepository) DeleteMeetingByID(id uint) error {
	ret := m.ctrl.Call(m, "DeleteMeetingByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddMeetingAttendance mocks the AddMeetingAttendance method.
func (m *MockMeetingRepository) AddMeetingAttendance(attendance models.MeetingAttendance) error {
	ret := m.ctrl.Call(m, "AddMeetingAttendance", attendance)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetMeetingsByTeamID mocks the GetMeetingsByTeamID method.
func (m *MockMeetingRepository) GetMeetingsByTeamID(teamID uint) ([]models.Meeting, error) {
	ret := m.ctrl.Call(m, "GetMeetingsByTeamID", teamID)
	ret0, _ := ret[0].([]models.Meeting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeetingsByTeamIDAndMeetingOver mocks the GetMeetingsByTeamIDAndMeetingOver method.
func (m *MockMeetingRepository) GetMeetingsByTeamIDAndMeetingOver(teamID uint, meetingOver bool) ([]models.Meeting, error) {
	ret := m.ctrl.Call(m, "GetMeetingsByTeamIDAndMeetingOver", teamID, meetingOver)
	ret0, _ := ret[0].([]models.Meeting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeetingAttendanceByMeetingID mocks the GetMeetingAttendanceByMeetingID method.
func (m *MockMeetingRepository) GetMeetingAttendanceByMeetingID(meetingID uint) ([]models.MeetingAttendance, error) {
	ret := m.ctrl.Call(m, "GetMeetingAttendanceByMeetingID", meetingID)
	ret0, _ := ret[0].([]models.MeetingAttendance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeetingAttendanceByMeetingIDAndOnTime mocks the GetMeetingAttendanceByMeetingIDAndOnTime method.
func (m *MockMeetingRepository) GetMeetingAttendanceByMeetingIDAndOnTime(meetingID uint, onTime bool) ([]models.MeetingAttendance, error) {
	ret := m.ctrl.Call(m, "GetMeetingAttendanceByMeetingIDAndOnTime", meetingID, onTime)
	ret0, _ := ret[0].([]models.MeetingAttendance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeetingAttendanceByUserIDAndMeetingID mocks the GetMeetingAttendanceByUserIDAndMeetingID method.
func (m *MockMeetingRepository) GetMeetingAttendanceByUserIDAndMeetingID(userID uint, meetingID uint) (models.MeetingAttendance, error) {
	ret := m.ctrl.Call(m, "GetMeetingAttendanceByUserIDAndMeetingID", userID, meetingID)
	ret0, _ := ret[0].(models.MeetingAttendance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeetingAttendancesByUserID mocks the GetMeetingAttendancesByUserID method.
func (m *MockMeetingRepository) GetMeetingAttendancesByUserID(userID uint) ([]models.MeetingAttendance, error) {
	ret := m.ctrl.Call(m, "GetMeetingAttendancesByUserID", userID)
	ret0, _ := ret[0].([]models.MeetingAttendance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MockMeetingRepositoryMockRecorder is a mock recorder for MockMeetingRepository.
type MockMeetingRepositoryMockRecorder struct {
	mock *MockMeetingRepository
}

// CreateMeeting mocks the CreateMeeting method.
func (mr *MockMeetingRepositoryMockRecorder) CreateMeeting(meeting models.Meeting) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMeeting", reflect.TypeOf((*MockMeetingRepository)(nil).CreateMeeting), meeting)
}

// GetMeetingByID mocks the GetMeetingByID method.
func (mr *MockMeetingRepositoryMockRecorder) GetMeetingByID(id uint) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeetingByID", reflect.TypeOf((*MockMeetingRepository)(nil).GetMeetingByID), id)
}

// UpdateMeeting mocks the UpdateMeeting method.
func (mr *MockMeetingRepositoryMockRecorder) UpdateMeeting(meeting models.Meeting) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMeeting", reflect.TypeOf((*MockMeetingRepository)(nil).UpdateMeeting), meeting)
}

// DeleteMeetingByID mocks the DeleteMeetingByID method.
func (mr *MockMeetingRepositoryMockRecorder) DeleteMeetingByID(id uint) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMeetingByID", reflect.TypeOf((*MockMeetingRepository)(nil).DeleteMeetingByID), id)
}

// AddMeetingAttendance mocks the AddMeetingAttendance method.
func (mr *MockMeetingRepositoryMockRecorder) AddMeetingAttendance(attendance models.MeetingAttendance) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMeetingAttendance", reflect.TypeOf((*MockMeetingRepository)(nil).AddMeetingAttendance), attendance)
}

// GetMeetingsByTeamID mocks the GetMeetingsByTeamID method.
func (mr *MockMeetingRepositoryMockRecorder) GetMeetingsByTeamID(teamID uint) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeetingsByTeamID", reflect.TypeOf((*MockMeetingRepository)(nil).GetMeetingsByTeamID), teamID)
}

// GetMeetingsByTeamIDAndMeetingOver mocks the GetMeetingsByTeamIDAndMeetingOver method.
func (mr *MockMeetingRepositoryMockRecorder) GetMeetingsByTeamIDAndMeetingOver(teamID uint, meetingOver bool) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeetingsByTeamIDAndMeetingOver", reflect.TypeOf((*MockMeetingRepository)(nil).GetMeetingsByTeamIDAndMeetingOver), teamID, meetingOver)
}

// GetMeetingAttendanceByMeetingID mocks the GetMeetingAttendanceByMeetingID method.
func (mr *MockMeetingRepositoryMockRecorder) GetMeetingAttendanceByMeetingID(meetingID uint) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeetingAttendanceByMeetingID", reflect.TypeOf((*MockMeetingRepository)(nil).GetMeetingAttendanceByMeetingID), meetingID)
}

// GetMeetingAttendanceByMeetingIDAndOnTime mocks the GetMeetingAttendanceByMeetingIDAndOnTime method.
func (mr *MockMeetingRepositoryMockRecorder) GetMeetingAttendanceByMeetingIDAndOnTime(meetingID uint, onTime bool) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeetingAttendanceByMeetingIDAndOnTime", reflect.TypeOf((*MockMeetingRepository)(nil).GetMeetingAttendanceByMeetingIDAndOnTime), meetingID, onTime)
}

// GetMeetingAttendanceByUserIDAndMeetingID mocks the GetMeetingAttendanceByUserIDAndMeetingID method.
func (mr *MockMeetingRepositoryMockRecorder) GetMeetingAttendanceByUserIDAndMeetingID(userID uint, meetingID uint) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeetingAttendanceByUserIDAndMeetingID", reflect.TypeOf((*MockMeetingRepository)(nil).GetMeetingAttendanceByUserIDAndMeetingID), userID, meetingID)
}

// GetMeetingAttendancesByUserID mocks the GetMeetingAttendancesByUserID method.
func (mr *MockMeetingRepositoryMockRecorder) GetMeetingAttendancesByUserID(userID uint) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeetingAttendancesByUserID", reflect.TypeOf((*MockMeetingRepository)(nil).GetMeetingAttendancesByUserID), userID)
}
