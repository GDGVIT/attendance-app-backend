package controllers

import (
	"net/http"
	"time"

	"github.com/GDGVIT/attendance-app-backend/infra/logger"
	"github.com/GDGVIT/attendance-app-backend/models"
	"github.com/GDGVIT/attendance-app-backend/repository"
	"github.com/GDGVIT/attendance-app-backend/utils/auth"
	"github.com/GDGVIT/attendance-app-backend/utils/email"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepo   *repository.UserRepository
	forgotRepo *repository.ForgotPasswordRepository
	verifRepo  *repository.VerificationEntryRepository
}

func NewUserController() *UserController {
	userRepo := repository.NewUserRepository()
	forgotRepo := repository.NewForgotPasswordRepository()
	verifRepo := repository.NewVerificationEntryRepository()
	return &UserController{userRepo, forgotRepo, verifRepo}
}

// RegisterUser handles user registration
func (uc *UserController) RegisterUser(c *gin.Context) {
	var registerData struct {
		Email        string `json:"email"`
		Name         string `json:"name"`
		Password     string `json:"password"`
		ProfileImage string `json:"profile_image"`
	}

	if err := c.ShouldBindJSON(&registerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Errorf("Failed to bind JSON: %v", err)
		return
	}

	// Check if the user already exists
	existingUser, _ := uc.userRepo.GetUserByEmail(registerData.Email)

	var emptyUser models.User
	if existingUser != emptyUser {
		email.SendRegistrationMail("Account Alert", "Someone attempted to create an account using your email. If this was you, try applying for password reset in case you have lost access to your account.", existingUser.Email, existingUser.ID, existingUser.Name, false)
		c.JSON(http.StatusCreated, gin.H{"message": "User created. Verification email sent!"})
		// c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		// we lie!
		return
	}

	if !auth.CheckPasswordStrength(registerData.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password not strong enough."})
		return
	}

	user := models.User{Name: registerData.Name, Email: registerData.Email, Password: registerData.Password, ProfileImage: registerData.ProfileImage}

	// Hash the user's password
	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		logger.Errorf("Failed to hash password: %v", err)
		return
	}

	// Create the user in the database
	if err := uc.userRepo.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		logger.Errorf("Failed to create user: %v", err)
		return
	}

	email.SendRegistrationMail("Account Verification.", "Please visit the following link to verify your account: ", user.Email, user.ID, user.Name, true)
	c.JSON(http.StatusCreated, gin.H{"message": "User created. Verification email sent!"})
	logger.Infof("New User Created.")
}

// Login handles user login
func (uc *UserController) Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Email: loginData.Email, Password: loginData.Password}

	token, user, err := auth.LoginCheck(user.Email, user.Password)

	if !user.Verified {
		c.JSON(http.StatusForbidden, gin.H{"error": "Please verify your email first."})
		return
	}

	if err != nil {
		println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "The email or password is not correct"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}

func (uc *UserController) RequestVerificationAgain(c *gin.Context) {
	useremail := c.Query("email")

	user, err := uc.userRepo.GetUserByEmail(useremail)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Verification email sent."})
		return
	}

	if user.Verified {
		c.JSON(http.StatusOK, gin.H{"message": "Verification email sent."})
		return
	}

	_, err = uc.verifRepo.GetVerificationEntryByEmail(user.Email)
	if err == nil {
		err = uc.verifRepo.DeleteVerificationEntry(user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting verification entry."})
			return
		}
	}

	// Send verification email
	err = email.SendRegistrationMail("Account Verification.", "Please visit the following link to verify your account: ", user.Email, user.ID, user.Name, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in sending email."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Verification email sent to you again."})
	logger.Infof("Verification requested again")
}

// VerifyEmail takes your email and otp sent of registration to verify a user account.
func (uc *UserController) VerifyEmail(c *gin.Context) {
	email := c.Query("email")
	otp := c.Query("otp")

	// Fetch the verification entry by email
	verificationEntry, err := uc.verifRepo.GetVerificationEntryByEmail(email)
	if err != nil {
		logger.Errorf("Error while verifying: " + err.Error())
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid verification."})
		return
	}

	if verificationEntry.OTP == otp {
		// Verify the email by updating the user's verification status
		err = uc.userRepo.VerifyUserEmail(email)
		if err != nil {
			logger.Errorf("Error while verifying: " + err.Error())
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid verification."})
			return
		}

		// Delete the verification entry
		err = uc.verifRepo.DeleteVerificationEntry(email)
		if err != nil {
			logger.Errorf("Error while deleting verification entry: " + err.Error())
		}

		c.JSON(http.StatusOK, gin.H{"message": "Verified! You can now log in."})
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid verification."})
	}
}

// ForgotPasswordRequest handles forgot password requests by sending a mail with an OTP
func (uc *UserController) ForgotPasswordRequest(c *gin.Context) {
	useremail := c.Query("email")

	// Fetch the user by email
	user, err := uc.userRepo.GetUserByEmail(useremail)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Forgot Password mail sent."})
		return
	}

	// Check if a forgot password entry already exists for the user's email
	err = uc.forgotRepo.DeleteForgotPasswordByEmail(user.Email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Forgot Password mail sent."})
		return
	}

	// Send the forgot password email
	err = email.SendForgotPasswordMail(user.Email, user.ID, user.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in sending email."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Forgot Password mail sent."})
	logger.Infof("Forgot password request")
}

// SetNewPassword sets a new password for the user after forgot password request
func (uc *UserController) SetNewPassword(c *gin.Context) {
	var forgotPasswordInput struct {
		NewPassword string `json:"new_password"`
	}
	if err := c.ShouldBindJSON(&forgotPasswordInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	useremail := c.Query("email")
	otp := c.Query("otp")

	// Fetch the forgot password entry by email
	forgotPasswordEntry, err := uc.forgotRepo.GetForgotPasswordByEmail(useremail)
	if err != nil {
		logger.Errorf("Error while verifying: %v", err.Error())
		c.JSON(http.StatusForbidden, gin.H{"message": "Invalid verification."})
		return
	}

	if forgotPasswordEntry.ValidTill.Before(time.Now()) {
		c.JSON(http.StatusForbidden, gin.H{"error": "token expired, please request forgot password again."})
		return
	}

	if forgotPasswordEntry.OTP == otp {
		// Fetch the user by email
		user, err := uc.userRepo.GetUserByEmail(useremail)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
			return
		}

		if !auth.CheckPasswordStrength(forgotPasswordInput.NewPassword) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Password not strong enough."})
			return
		}
		user.Password = forgotPasswordInput.NewPassword
		user.HashPassword()

		err = uc.userRepo.SaveUser(user)
		if err != nil {
			logger.Errorf("Save user after forgot and new: " + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
			return
		}

		email.GenericSendMail("Password Reset", "Password for your account was reset recently.", user.Email, user.Name)

		// Delete the forgot password entry
		err = uc.forgotRepo.DeleteForgotPasswordByEmail(useremail)
		if err != nil {
			logger.Errorf("Error while deleting forgot password entry: " + err.Error())
		}

		c.JSON(http.StatusOK, gin.H{"message": "Password set successfully. Please proceed to login."})
	} else {
		c.JSON(http.StatusForbidden, gin.H{"message": "Invalid verification. Password not updated."})
	}
}
