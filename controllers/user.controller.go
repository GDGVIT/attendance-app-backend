package controllers

import (
	"net/http"

	"github.com/GDGVIT/attendance-app-backend/infra/logger"
	"github.com/GDGVIT/attendance-app-backend/models"
	"github.com/GDGVIT/attendance-app-backend/repository"
	"github.com/GDGVIT/attendance-app-backend/utils/auth"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepo *repository.UserRepository
}

func NewUserController() *UserController {
	userRepo := repository.NewUserRepository()
	return &UserController{userRepo}
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
	existingUser, err := uc.userRepo.GetUserByEmail(registerData.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		logger.Errorf("Failed to get user by email: %v", err)
		return
	}
	var emptyUser models.User
	if existingUser != emptyUser {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
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

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
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
