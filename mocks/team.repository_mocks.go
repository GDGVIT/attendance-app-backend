package mocks

import (
	"github.com/GDGVIT/attendance-app-backend/models"
	"github.com/golang/mock/gomock"
)

// MockTeamRepository is a mock for TeamRepositoryInterface.
type MockTeamRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTeamRepositoryMockRecorder
}

// NewMockTeamRepository creates a new mock for TeamRepositoryInterface.
func NewMockTeamRepository(ctrl *gomock.Controller) *MockTeamRepository {
	mock := &MockTeamRepository{ctrl: ctrl}
	mock.recorder = &MockTeamRepositoryMockRecorder{mock}
	return mock
}

// EXPECT methods for expected calls with return values
func (m *MockTeamRepository) EXPECT() *MockTeamRepositoryMockRecorder {
	return m.recorder
}

// CreateTeam mocks the CreateTeam method.
func (m *MockTeamRepository) CreateTeam(team models.Team) (models.Team, error) {
	ret := m.ctrl.Call(m, "CreateTeam", team)
	ret0, _ := ret[0].(models.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTeamByID mocks the GetTeamByID method.
func (m *MockTeamRepository) GetTeamByID(id uint) (models.Team, error) {
	ret := m.ctrl.Call(m, "GetTeamByID", id)
	ret0, _ := ret[0].(models.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTeamByInvite mocks the GetTeamByInvite method.
func (m *MockTeamRepository) GetTeamByInvite(inviteCode string) (models.Team, error) {
	ret := m.ctrl.Call(m, "GetTeamByInvite", inviteCode)
	ret0, _ := ret[0].(models.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTeam mocks the UpdateTeam method.
func (m *MockTeamRepository) UpdateTeam(team models.Team) (models.Team, error) {
	ret := m.ctrl.Call(m, "UpdateTeam", team)
	ret0, _ := ret[0].(models.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTeamByID mocks the DeleteTeamByID method.
func (m *MockTeamRepository) DeleteTeamByID(id uint) error {
	ret := m.ctrl.Call(m, "DeleteTeamByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetUnprotectedTeams mocks the GetUnprotectedTeams method.
func (m *MockTeamRepository) GetUnprotectedTeams() ([]models.Team, error) {
	ret := m.ctrl.Call(m, "GetUnprotectedTeams")
	ret0, _ := ret[0].([]models.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTeamSuperAdmin mocks the UpdateTeamSuperAdmin method.
func (m *MockTeamRepository) UpdateTeamSuperAdmin(teamID, userID uint) (models.Team, error) {
	ret := m.ctrl.Call(m, "UpdateTeamSuperAdmin", teamID, userID)
	ret0, _ := ret[0].(models.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MockTeamRepositoryMockRecorder is a recorder for the MockTeamRepository.
type MockTeamRepositoryMockRecorder struct {
	mock *MockTeamRepository
}

// CreateTeam mocks the CreateTeam method.
func (m *MockTeamRepositoryMockRecorder) CreateTeam(team models.Team) *gomock.Call {
	return m.mock.ctrl.RecordCall(m.mock, "CreateTeam", team)
}

// GetTeamByID mocks the GetTeamByID method.
func (m *MockTeamRepositoryMockRecorder) GetTeamByID(id uint) *gomock.Call {
	return m.mock.ctrl.RecordCall(m.mock, "GetTeamByID", id)
}

// GetTeamByInvite mocks the GetTeamByInvite method.
func (m *MockTeamRepositoryMockRecorder) GetTeamByInvite(inviteCode string) *gomock.Call {
	return m.mock.ctrl.RecordCall(m.mock, "GetTeamByInvite", inviteCode)
}

// UpdateTeam mocks the UpdateTeam method.
func (m *MockTeamRepositoryMockRecorder) UpdateTeam(team models.Team) *gomock.Call {
	return m.mock.ctrl.RecordCall(m.mock, "UpdateTeam", team)
}

// DeleteTeamByID mocks the DeleteTeamByID method.
func (m *MockTeamRepositoryMockRecorder) DeleteTeamByID(id uint) *gomock.Call {
	return m.mock.ctrl.RecordCall(m.mock, "DeleteTeamByID", id)
}

// GetUnprotectedTeams mocks the GetUnprotectedTeams method.
func (m *MockTeamRepositoryMockRecorder) GetUnprotectedTeams() *gomock.Call {
	return m.mock.ctrl.RecordCall(m.mock, "GetUnprotectedTeams")
}

// UpdateTeamSuperAdmin mocks the UpdateTeamSuperAdmin method.
func (m *MockTeamRepositoryMockRecorder) UpdateTeamSuperAdmin(teamID, userID uint) *gomock.Call {
	return m.mock.ctrl.RecordCall(m.mock, "UpdateTeamSuperAdmin", teamID, userID)
}
