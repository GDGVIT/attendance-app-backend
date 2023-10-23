package mocks

import (
	"github.com/GDGVIT/attendance-app-backend/models"
	"github.com/golang/mock/gomock"
)

// MockEmailService is a mock for EmailServiceInterface.
type MockEmailService struct {
	ctrl     *gomock.Controller
	recorder *MockEmailServiceMockRecorder
}

// NewMockEmailService creates a new mock for EmailServiceInterface.
func NewMockEmailService(ctrl *gomock.Controller) *MockEmailService {
	mock := &MockEmailService{ctrl: ctrl}
	mock.recorder = &MockEmailServiceMockRecorder{mock}
	return mock
}

// EXPECT methods for expected calls with return values
func (m *MockEmailService) EXPECT() *MockEmailServiceMockRecorder {
	return m.recorder
}

// SendMeetingNotification mocks the SendMeetingNotification method.
func (m *MockEmailService) SendMeetingNotification(teamID uint, meeting models.Meeting) error {
	ret := m.ctrl.Call(m, "SendMeetingNotification", teamID, meeting)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericSendMail mocks the GenericSendMail method.
func (m *MockEmailService) GenericSendMail(subject string, content string, toEmail string, userName string) error {
	ret := m.ctrl.Call(m, "GenericSendMail", subject, content, toEmail, userName)
	ret0, _ := ret[0].(error)
	return ret0
}

// MockEmailServiceMockRecorder is a recorder for the MockEmailService.
type MockEmailServiceMockRecorder struct {
	mock *MockEmailService
}

// SendMeetingNotification mocks the SendMeetingNotification method.
func (m *MockEmailServiceMockRecorder) SendMeetingNotification(teamID uint, meeting models.Meeting) *gomock.Call {
	return m.mock.ctrl.RecordCall(m.mock, "SendMeetingNotification", teamID, meeting)
}

// GenericSendMail mocks the GenericSendMail method.
func (m *MockEmailServiceMockRecorder) GenericSendMail(subject string, content string, toEmail string, userName string) *gomock.Call {
	return m.mock.ctrl.RecordCall(m.mock, "GenericSendMail", subject, content, toEmail, userName)
}
