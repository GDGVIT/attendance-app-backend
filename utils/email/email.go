package email

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"net/http"
	"time"

	"github.com/GDGVIT/attendance-app-backend/infra/logger"
	"github.com/GDGVIT/attendance-app-backend/models"
	"github.com/GDGVIT/attendance-app-backend/repository"
	"github.com/spf13/viper"
)

func GenerateOTP(maxDigits uint32) string {
	bi, err := rand.Int(
		rand.Reader,
		big.NewInt(int64(math.Pow(10, float64(maxDigits)))),
	)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%0*d", maxDigits, bi)
}

type EmailAddress struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type RegistrationEmail struct {
	Subject  string         `json:"subject"`
	From     EmailAddress   `json:"from"`
	To       []EmailAddress `json:"to"`
	Category string         `json:"category"`
	Text     string         `json:"text"`
}

func GenericSendMail(subject string, content string, toEmail string, userName string) error {
	url := "https://send.api.mailtrap.io/api/send"
	method := "POST"

	data := RegistrationEmail{
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

func SendRegistrationMail(subject string, content string, toEmail string, userID uint, userName string, newUser bool) error {
	otp := ""
	if newUser {
		otp = GenerateOTP(6)
		content += viper.GetString("FRONTEND_BASE") + "/verify?email=" + toEmail + "&otp=" + otp + " . Or, type in the OTP manually: " + otp
	}

	err := GenericSendMail(subject, content, toEmail, userName)
	if err != nil {
		return err
	}

	if newUser {
		entry := models.VerificationEntry{
			Email: toEmail,
			OTP:   otp,
		}
		verifRepo := repository.NewVerificationEntryRepository()
		verifRepo.CreateVerificationEntry(entry)
	}
	return nil
}

func SendDeletionMail(toEmail string, userID uint, userName string) error {
	otp := ""
	otp = GenerateOTP(6)
	confirmationURL := ""
	confirmationURL += viper.GetString("FRONTEND_BASE") + "/delete-account?email=" + toEmail + "&otp=" + otp
	content := "A request for the deletion of the nock account associated with your user has been made. If this was not you, please change your password. Otherwise, click on this link to confirm account deletion: " + confirmationURL + " . This link will be active for 3 minutes."
	subject := "Request for account deletion."

	err := GenericSendMail(subject, content, toEmail, userName)
	if err != nil {
		return err
	}

	entry := models.DeletionConfirmation{
		Email:     toEmail,
		OTP:       otp,
		ValidTill: time.Now().Add(3 * time.Minute),
	}
	deletionRepo := repository.NewDeletionConfirmationRepository()
	deletionRepo.CreateDeletionConfirmation(entry)
	return nil
}

func SendForgotPasswordMail(toEmail string, userID uint, userName string) error {
	otp := ""
	otp = GenerateOTP(6)
	verificationURL := ""
	verificationURL += viper.GetString("FRONTEND_BASE") + "/set-forgotten-password?email=" + toEmail + "&otp=" + otp
	content := "A forgot password request was made for the email associated with your account. If this was not you, feel free to ignore this email. Otherwise, click on this link to post your new password: " + verificationURL + " . This link will be active for 3 minutes."
	subject := "Forgot Password."

	err := GenericSendMail(subject, content, toEmail, userName)
	if err != nil {
		return err
	}

	entry := models.ForgotPassword{
		Email:     toEmail,
		OTP:       otp,
		ValidTill: time.Now().Add(3 * time.Minute),
	}
	forgotRepo := repository.NewForgotPasswordRepository()
	forgotRepo.CreateForgotPassword(entry)
	return nil
}

// SendRequestNotifToTeamAdmins sends a notification email to all admins of a team
func SendRequestNotifToTeamAdmins(toEmail []string, userName string, userEmail string, teamName string) error {
	content := userName + " has requested to join your team " + teamName + ". Please accept or reject the request on the Attendance App."
	subject := "Team Join Request."

	for _, email := range toEmail {
		err := GenericSendMail(subject, content, email, teamName+" Admin Team")
		if err != nil {
			return err
		}
	}
	return nil
}

// SendRequestStatusNotifToUser
func SendRequestStatusNotifToUser(toEmail string, userName string, teamName string, status string) error {
	content := "Your request to join the team " + teamName + " has been " + status + "."
	subject := "Team Join Request."

	err := GenericSendMail(subject, content, toEmail, userName)
	if err != nil {
		return err
	}
	return nil
}

// SendKickNotifToUser
func SendKickNotifToUser(toEmail string, userName string, teamName string) error {
	content := "You have been removed from the team " + teamName + " by its Admin Team."
	subject := "Team Kicked."

	err := GenericSendMail(subject, content, toEmail, userName)
	if err != nil {
		return err
	}
	return nil
}

// send email to all team members about a new meeting
func SendMeetingNotifToTeamMembers(toEmail []string, teamName string, meetingName string, meetingTime time.Time) error {
	content := "A new meeting " + meetingName + " has been scheduled for the team " + teamName + " at " + meetingTime.String() + "."
	subject := "New Meeting."

	for _, email := range toEmail {
		err := GenericSendMail(subject, content, email, teamName+" Team")
		if err != nil {
			return err
		}
	}
	return nil
}
