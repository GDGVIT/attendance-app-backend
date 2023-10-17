package controllers

import (
	"net/http"

	"github.com/GDGVIT/attendance-app-backend/models"
	"github.com/gin-gonic/gin"
)

// GetMyDetails returns details of the authenticated user.
func (uc *UserController) GetMyDetails(c *gin.Context) {
	// Retrieve the authenticated user from the context
	currentUser, _ := c.Get("user")

	// Type-assert the user to the models.User struct
	user := currentUser.(*models.User)

	// Respond with the user details
	c.JSON(http.StatusOK, user)
}

// UpdateMyDetails updates the name and profile picture of the authenticated user.
func (uc *UserController) UpdateMyDetails(c *gin.Context) {
	// Retrieve the authenticated user from the context
	currentUser, _ := c.Get("user")

	// Type-assert the user to the models.User struct
	user := currentUser.(*models.User)

	// Define a struct to bind the JSON request
	var userUpdateRequest struct {
		Name         string `json:"name"`
		ProfileImage string `json:"profile_image"`
	}

	// Bind the JSON request to the userUpdateRequest struct
	if err := c.ShouldBindJSON(&userUpdateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the user's name and profile picture
	if userUpdateRequest.Name != "" {
		user.Name = userUpdateRequest.Name
	}

	if userUpdateRequest.ProfileImage != "" {
		user.ProfileImage = userUpdateRequest.ProfileImage
	}

	// Update the user in the database
	err := uc.userRepo.SaveUser(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	// Respond with the updated user details
	c.JSON(http.StatusOK, gin.H{"message": "User details updated successfully."})
}
