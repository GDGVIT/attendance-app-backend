// have to refactor other email functions as well

package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GDGVIT/attendance-app-backend/infra/logger"
	"github.com/GDGVIT/attendance-app-backend/models"
	"github.com/GDGVIT/attendance-app-backend/repository"
	"github.com/spf13/viper"
)

// EmailService handles sending email notifications.
type EmailService struct {
	teamRepo       repository.TeamRepositoryInterface
	teamMemberRepo repository.TeamMemberRepositoryInterface
	userRepo       repository.UserRepositoryInterface
}

// NewEmailService creates a new EmailService.
func NewEmailService(teamRepo repository.TeamRepositoryInterface, teamMemberRepo repository.TeamMemberRepositoryInterface, userRepo repository.UserRepositoryInterface) *EmailService {
	return &EmailService{
		teamRepo,
		teamMemberRepo,
		userRepo,
	}
}

type EmailAddress struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type GenericEmail struct {
	Subject  string         `json:"subject"`
	From     EmailAddress   `json:"from"`
	To       []EmailAddress `json:"to"`
	Category string         `json:"category"`
	Text     string         `json:"text"`
}

type EmailServiceInterface interface {
	SendMeetingNotification(teamID uint, meeting models.Meeting) error
	GenericSendMail(subject string, content string, toEmail string, userName string) error
}

// SendMeetingNotification sends a meeting notification to team members.
func (es *EmailService) SendMeetingNotification(teamID uint, meeting models.Meeting) error {
	// get name of the team
	team, err := es.teamRepo.GetTeamByID(teamID)
	if err != nil {
		return err
	}

	// get members of the team and send an email to each member
	teamMembers, err := es.teamMemberRepo.GetTeamMembersByTeamID(teamID)
	if err != nil {
		return err
	}

	teamMemberEmails := []string{}
	for _, teamMember := range teamMembers {
		user, err := es.userRepo.GetUserByID(teamMember.UserID)
		if err != nil {
			return err
		}
		// add user email to teamMemberEmails
		teamMemberEmails = append(teamMemberEmails, user.Email)
	}

	content := "A new meeting " + meeting.Title + " has been scheduled for the team " + team.Name + " at " + meeting.StartTime.String() + "."
	subject := "New Meeting."

	for _, email := range teamMemberEmails {
		err := es.GenericSendMail(subject, content, email, team.Name+" Team")
		if err != nil {
			return err
		}
	}

	return nil
}

// GenericSendMail sends a generic email.
func (es *EmailService) GenericSendMail(subject string, content string, toEmail string, userName string) error {
	url := "https://send.api.mailtrap.io/api/send"
	method := "POST"

	data := GenericEmail{
		Subject: subject,
		From: EmailAddress{
			Email: "nock.noreply@dscvit.com",
			Name:  "Attendance App",
		},
		To: []EmailAddress{
			{
				Email: toEmail,
				Name:  userName,
			},
		},
		Category: "AttendanceApp",
		Text:     content,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		logger.Errorf("Email Error: %v", err)
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))

	if err != nil {
		logger.Errorf("Email Error: %v", err)
		return err
	}

	bearer := fmt.Sprintf("Bearer %s", viper.GetString("MAILTRAP_API_TOKEN"))
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		logger.Errorf("Email Error: %v", err)
		return err
	}

	defer res.Body.Close()
	return nil
}
