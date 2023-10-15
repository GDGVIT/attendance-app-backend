package controllers

import (
	"net/http"

	"github.com/GDGVIT/attendance-app-backend/infra/logger"
	"github.com/GDGVIT/attendance-app-backend/models"
	"github.com/GDGVIT/attendance-app-backend/repository"
	"github.com/GDGVIT/attendance-app-backend/utils/auth"
	"github.com/GDGVIT/attendance-app-backend/utils/email"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepo     *repository.UserRepository
	forgotRepo   *repository.ForgotPasswordRepository
	verifRepo    *repository.VerificationEntryRepository
	deletionRepo *repository.DeletionConfirmationRepository
}

func NewUserController() *UserController {
	userRepo := repository.NewUserRepository()
	forgotRepo := repository.NewForgotPasswordRepository()
	verifRepo := repository.NewVerificationEntryRepository()
	deletionRepo := repository.NewDeletionConfirmationRepository()
	return &UserController{userRepo, forgotRepo, verifRepo, deletionRepo}
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Improper JSON."})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "password-strength", "message": "Password not strong enough."})
		return
	}

	user := models.User{Name: registerData.Name, Email: registerData.Email, Password: registerData.Password, ProfileImage: registerData.ProfileImage}

	// Hash the user's password
	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "hashing", "message": "Failed to hash password"})
		logger.Errorf("Failed to hash password: %v", err)
		return
	}

	// Create the user in the database
	if err := uc.userRepo.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "creation-error", "message": "Failed to create user."})
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

	if err != nil {
		println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "credentials-error", "message": "The email or password is not correct"})
		return
	}

	if !user.Verified {
		c.JSON(http.StatusForbidden, gin.H{"error": "unverified", "message": "Please verify your email before logging in."})
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "deletion", "message": "Error deleting verification entry."})
			return
		}
	}

	// Send verification email
	err = email.SendRegistrationMail("Account Verification.", "Please visit the following link to verify your account: ", user.Email, user.ID, user.Name, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "mail", "message": "Error in sending email."})
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
