package mocks

import (
	"github.com/GDGVIT/attendance-app-backend/models"
	"github.com/golang/mock/gomock"
)

// MockTeamMemberRepository is a mock for TeamMemberRepositoryInterface.
type MockTeamMemberRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTeamMemberRepositoryMockRecorder
}

// NewMockTeamMemberRepository creates a new mock for TeamMemberRepositoryInterface.
func NewMockTeamMemberRepository(ctrl *gomock.Controller) *MockTeamMemberRepository {
	mock := &MockTeamMemberRepository{ctrl: ctrl}
	mock.recorder = &MockTeamMemberRepositoryMockRecorder{mock}
	return mock
}

// EXPECT methods for expected calls with return values
func (m *MockTeamMemberRepository) EXPECT() *MockTeamMemberRepositoryMockRecorder {
	return m.recorder
}

// CreateTeamMember mocks the CreateTeamMember method.
func (m *MockTeamMemberRepository) CreateTeamMember(teamMember models.TeamMember) (models.TeamMember, error) {
	ret := m.ctrl.Call(m, "CreateTeamMember", teamMember)
	ret0, _ := ret[0].(models.TeamMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTeamMemberByID mocks the GetTeamMemberByID method.
func (m *MockTeamMemberRepository) GetTeamMemberByID(teamID, userID uint) (models.TeamMember, error) {
	ret := m.ctrl.Call(m, "GetTeamMemberByID", teamID, userID)
	ret0, _ := ret[0].(models.TeamMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTeamMember mocks the UpdateTeamMember method.
func (m *MockTeamMemberRepository) UpdateTeamMember(teamMember models.TeamMember) (models.TeamMember, error) {
	ret := m.ctrl.Call(m, "UpdateTeamMember", teamMember)
	ret0, _ := ret[0].(models.TeamMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTeamMember mocks the DeleteTeamMember method.
func (m *MockTeamMemberRepository) DeleteTeamMember(teamID, userID uint) error {
	ret := m.ctrl.Call(m, "DeleteTeamMember", teamID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetTeamMembersByTeamID mocks the GetTeamMembersByTeamID method.
func (m *MockTeamMemberRepository) GetTeamMembersByTeamID(teamID uint) ([]models.TeamMember, error) {
	ret := m.ctrl.Call(m, "GetTeamMembersByTeamID", teamID)
	ret0, _ := ret[0].([]models.TeamMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTeamMembersByUserID mocks the GetTeamMembersByUserID method.
func (m *MockTeamMemberRepository) GetTeamMembersByUserID(userID uint) ([]models.TeamMember, error) {
	ret := m.ctrl.Call(m, "GetTeamMembersByUserID", userID)
	ret0, _ := ret[0].([]models.TeamMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTeamMembersByUserAndRole mocks the GetTeamMembersByUserAndRole method.
func (m *MockTeamMemberRepository) GetTeamMembersByUserAndRole(userID uint, role string) ([]models.TeamMember, error) {
	ret := m.ctrl.Call(m, "GetTeamMembersByUserAndRole", userID, role)
	ret0, _ := ret[0].([]models.TeamMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTeamMembersByTeamAndRole mocks the GetTeamMembersByTeamAndRole method.
func (m *MockTeamMemberRepository) GetTeamMembersByTeamAndRole(teamID uint, role string) ([]models.TeamMember, error) {
	ret := m.ctrl.Call(m, "GetTeamMembersByTeamAndRole", teamID, role)
	ret0, _ := ret[0].([]models.TeamMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdminTeamByTeamID mocks the GetAdminTeamByTeamID method.
func (m *MockTeamMemberRepository) GetAdminTeamByTeamID(teamID uint) ([]models.TeamMember, error) {
	ret := m.ctrl.Call(m, "GetAdminTeamByTeamID", teamID)
	ret0, _ := ret[0].([]models.TeamMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTeamMemberRole mocks the UpdateTeamMemberRole method.
func (m *MockTeamMemberRepository) UpdateTeamMemberRole(teamMemberID uint, role string) (models.TeamMember, error) {
	ret := m.ctrl.Call(m, "UpdateTeamMemberRole", teamMemberID, role)
	ret0, _ := ret[0].(models.TeamMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTeamMemberByUserID mocks the DeleteTeamMemberByUserID method.
func (m *MockTeamMemberRepository) DeleteTeamMemberByUserID(userID uint) error {
	ret := m.ctrl.Call(m, "DeleteTeamMemberByUserID", userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// MockTeamMemberRepositoryMockRecorder is a recorder for the MockTeamMemberRepository.
type MockTeamMemberRepositoryMockRecorder struct {
	mock *MockTeamMemberRepository
}

// CreateTeamMember mocks the CreateTeamMember method.
func (m *MockTeamMemberRepositoryMockRecorder) CreateTeamMember(teamMember models.TeamMember) *gomock.Call {
	return m.mock.ctrl.RecordCall(m.mock, "CreateTeamMember", teamMember)
}

// GetTeamMemberByID mocks the GetTeamMemberByID method.
func (m *MockTeamMemberRepositoryMockRecorder) GetTeamMemberByID(teamID, userID uint) *gomock.Call {
	return m.mock.ctrl.RecordCall(m.mock, "GetTeamMemberByID", teamID, userID)
}

// UpdateTeamMember mocks the UpdateTeamMember method.
func (m *MockTeamMemberRepositoryMockRecorder) UpdateTeamMember(teamMember models.TeamMember) *gomock.Call {
	return m.mock.ctrl.RecordCall(m.mock, "UpdateTeamMember", teamMember)
}

// DeleteTeamMember mocks the DeleteTeamMember method.
func (m *MockTeamMemberRepositoryMockRecorder) DeleteTeamMember(teamID, userID uint) *gomock.Call {
	return m.mock.ctrl.RecordCall(m.mock, "DeleteTeamMember", teamID, userID)
}

// GetTeamMembersByTeamID mocks the GetTeamMembersByTeamID method.
func (m *MockTeamMemberRepositoryMockRecorder) GetTeamMembersByTeamID(teamID uint) *gomock.Call {
	return m.mock.ctrl.RecordCall(m.mock, "GetTeamMembersByTeamID", teamID)
}

// GetTeamMembersByUserID mocks the GetTeamMembersByUserID method.
func (m *MockTeamMemberRepositoryMockRecorder) GetTeamMembersByUserID(userID uint) *gomock.Call {
	return m.mock.ctrl.RecordCall(m.mock, "GetTeamMembersByUserID", userID)
}

// GetTeamMembersByUserAndRole mocks the GetTeamMembersByUserAndRole method.
func (m *MockTeamMemberRepositoryMockRecorder) GetTeamMembersByUserAndRole(userID uint, role string) *gomock.Call {
	return m.mock.ctrl.RecordCall(m.mock, "GetTeamMembersByUserAndRole", userID, role)
}

// GetTeamMembersByTeamAndRole mocks the GetTeamMembersByTeamAndRole method.
func (m *MockTeamMemberRepositoryMockRecorder) GetTeamMembersByTeamAndRole(teamID uint, role string) *gomock.Call {
	return m.mock.ctrl.RecordCall(m.mock, "GetTeamMembersByTeamAndRole", teamID, role)
}

// GetAdminTeamByTeamID mocks the GetAdminTeamByTeamID method.
func (m *MockTeamMemberRepositoryMockRecorder) GetAdminTeamByTeamID(teamID uint) *gomock.Call {
	return m.mock.ctrl.RecordCall(m.mock, "GetAdminTeamByTeamID", teamID)
}

// UpdateTeamMemberRole mocks the UpdateTeamMemberRole method.
func (m *MockTeamMemberRepositoryMockRecorder) UpdateTeamMemberRole(teamMemberID uint, role string) *gomock.Call {
	return m.mock.ctrl.RecordCall(m.mock, "UpdateTeamMemberRole", teamMemberID, role)
}

// DeleteTeamMemberByUserID mocks the DeleteTeamMemberByUserID method.
func (m *MockTeamMemberRepositoryMockRecorder) DeleteTeamMemberByUserID(userID uint) *gomock.Call {
	return m.mock.ctrl.RecordCall(m.mock, "DeleteTeamMemberByUserID", userID)
}
