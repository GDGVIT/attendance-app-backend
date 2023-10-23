package services

import (
	"testing"

	"github.com/GDGVIT/attendance-app-backend/mocks" // Import your mock package
	"github.com/GDGVIT/attendance-app-backend/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestEmailService_SendMeetingNotification(t *testing.T) {
	// Create a new instance of the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock repository
	mockTeamRepo := mocks.NewMockTeamRepository(ctrl)
	mockTeamMemberRepo := mocks.NewMockTeamMemberRepository(ctrl)
	mockUserRepo := mocks.NewMockUserRepository(ctrl)

	// Create an instance of EmailService
	emailService := NewEmailService(mockTeamRepo, mockTeamMemberRepo, mockUserRepo)

	// Define your test data
	teamID := uint(1)
	meeting := models.Meeting{
		// Initialize with meeting data
	}

	// Set expectations on the mock repositories
	mockTeamRepo.EXPECT().GetTeamByID(teamID).Return(models.Team{}, nil)
	mockTeamMemberRepo.EXPECT().GetTeamMembersByTeamID(teamID).Return([]models.TeamMember{}, nil)
	mockUserRepo.EXPECT().GetUserByID(1).Return(models.User{}, nil).AnyTimes()

	// Call the SendMeetingNotification method
	err := emailService.SendMeetingNotification(teamID, meeting)

	// Assert that no error occurred during the execution
	assert.NoError(t, err)
}

func TestEmailService_GenericSendMail(t *testing.T) {
	// Create a new instance of the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTeamRepo := mocks.NewMockTeamRepository(ctrl)
	mockTeamMemberRepo := mocks.NewMockTeamMemberRepository(ctrl)
	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	emailService := NewEmailService(mockTeamRepo, mockTeamMemberRepo, mockUserRepo)

	// Define your test data
	subject := "Test Subject"
	content := "Test Content"
	toEmail := "anirudh04mishra@gmail.com"
	userName := "Test User"

	// Call the GenericSendMail method
	err := emailService.GenericSendMail(subject, content, toEmail, userName)

	// Assert that no error occurred during the execution
	assert.NoError(t, err)
}
