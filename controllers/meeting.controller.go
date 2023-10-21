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
func (mc *MeetingController) GetMeetingsByTeamID(c *gin.Context) {
	// Get team ID from route parameters
	teamID, err := strconv.ParseUint(c.Param("teamID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}

	// Get filterBy query parameter
	filterBy := c.Query("filterBy")

	// Get orderBy query parameter
	orderBy := c.Query("orderBy")

	// Call the meeting service to get the meetings
	meetings, err := mc.meetingService.GetMeetingsByTeamID(uint(teamID), filterBy, orderBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get the meetings", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, meetings)
}

// GetMeetingDetails retrieves a meeting by its ID.
func (mc *MeetingController) GetMeetingDetails(c *gin.Context) {
	// Get meeting ID from route parameters
	meetingID, err := strconv.ParseUint(c.Param("meetingID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid meeting ID"})
		return
	}

	// Call the meeting service to get the meeting
	meeting, err := mc.meetingService.GetMeetingByID(uint(meetingID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get the meeting", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, meeting)
}

// StartMeeting starts a meeting.
func (mc *MeetingController) StartMeeting(c *gin.Context) {
	// Get meeting ID from route parameters
	meetingID, err := strconv.ParseUint(c.Param("meetingID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid meeting ID"})
		return
	}

	// Call the meeting service to start the meeting
	meeting, err := mc.meetingService.StartMeeting(uint(meetingID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to start the meeting", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, meeting)
}

// StartAttendance starts attendance for a meeting.
func (mc *MeetingController) StartAttendance(c *gin.Context) {
	// Get meeting ID from route parameters
	meetingID, err := strconv.ParseUint(c.Param("meetingID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid meeting ID"})
		return
	}

	// Call the meeting service to start attendance for the meeting
	meeting, err := mc.meetingService.StartAttendance(uint(meetingID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to start attendance for the meeting", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, meeting)
}

// EndAttendance ends attendance for a meeting.
func (mc *MeetingController) EndAttendance(c *gin.Context) {
	// Get meeting ID from route parameters
	meetingID, err := strconv.ParseUint(c.Param("meetingID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid meeting ID"})
		return
	}

	// Call the meeting service to end attendance for the meeting
	meeting, err := mc.meetingService.EndAttendance(uint(meetingID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to end attendance for the meeting", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, meeting)
}

// EndMeeting ends a meeting.
func (mc *MeetingController) EndMeeting(c *gin.Context) {
	// Get meeting ID from route parameters
	meetingID, err := strconv.ParseUint(c.Param("meetingID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid meeting ID"})
		return
	}

	// Call the meeting service to end the meeting
	meeting, err := mc.meetingService.EndMeeting(uint(meetingID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to end the meeting", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, meeting)
}

// DeleteMeetingByID deletes a meeting by its ID.
func (mc *MeetingController) DeleteMeetingByID(c *gin.Context) {
	// Get meeting ID from route parameters
	meetingID, err := strconv.ParseUint(c.Param("meetingID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid meeting ID"})
		return
	}

	// Call the meeting service to delete the meeting
	err = mc.meetingService.DeleteMeetingByID(uint(meetingID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete the meeting", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Meeting deleted successfully"})
}
