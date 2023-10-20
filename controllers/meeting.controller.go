package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/GDGVIT/attendance-app-backend/models"
	"github.com/GDGVIT/attendance-app-backend/services"
	"github.com/gin-gonic/gin"
)

// MeetingController handles meeting-related routes.
type MeetingController struct {
	meetingService services.MeetingServiceInterface
}

// NewMeetingController creates a new MeetingController.
func NewMeetingController(meetingService services.MeetingServiceInterface) *MeetingController {
	return &MeetingController{meetingService}
}

// CreateMeeting creates a new meeting.
func (mc *MeetingController) CreateMeeting(c *gin.Context) {
	// Get team ID from route parameters
	teamID, err := strconv.ParseUint(c.Param("teamID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}

	var newMeeting struct {
		Title       string          `json:"title" binding:"required"`
		Description string          `json:"description" binding:"required"`
		Venue       string          `json:"venue" binding:"required"`
		Location    models.Location `json:"location" binding:"required"`
		StartTime   time.Time       `json:"startTime" binding:"required"`
	}

	// Bind request body to meeting structure
	if err := c.ShouldBindJSON(&newMeeting); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body", "error": err.Error()})
		return
	}

	meeting := models.Meeting{
		TeamID:      uint(teamID),
		Title:       newMeeting.Title,
		Description: newMeeting.Description,
		Venue:       newMeeting.Venue,
		Location:    newMeeting.Location,
		StartTime:   newMeeting.StartTime,
	}

	// Call the meeting service to create the meeting
	createdMeeting, err := mc.meetingService.CreateMeeting(meeting.TeamID, meeting.Title, meeting.Description, meeting.Venue, meeting.Location, meeting.StartTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create the meeting", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdMeeting)
}

// GetMeetingsByTeamIDAndMeetingOver retrieves all meetings for a team.
func (mc *MeetingController) GetMeetingsByTeamIDAndMeetingOver(c *gin.Context) {
	// Get team ID from route parameters
	teamID, err := strconv.ParseUint(c.Param("teamID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}

	// Get meetingOver from query parameters
	meetingOver, err := strconv.ParseBool(c.Query("meetingOver"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid meetingOver"})
		return
	}

	// Call the meeting service to get all meetings for a team
	meetings, err := mc.meetingService.GetMeetingsByTeamIDAndMeetingOver(uint(teamID), meetingOver)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get meetings", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, meetings)
}
