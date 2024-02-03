package routers

import (
	"net/http"

	"github.com/GDGVIT/attendance-app-backend/controllers"
	"github.com/GDGVIT/attendance-app-backend/repository"
	"github.com/GDGVIT/attendance-app-backend/routers/middleware"
	"github.com/GDGVIT/attendance-app-backend/services"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes add all routing list here automatically get main router
func RegisterRoutes(route *gin.Engine) {
	route.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})
	route.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusTeapot, gin.H{"live": "ok"}) })

	v1 := route.Group("/v1")

	meetingRepo := repository.NewMeetingRepository()
	teamRepo := repository.NewTeamRepository()
	userRepo := repository.NewUserRepository()
	teamMemberRepo := repository.NewTeamMemberRepository()
	emailService := services.NewEmailService(teamRepo, teamMemberRepo, userRepo)
	meetingService := services.NewMeetingService(meetingRepo, emailService, userRepo, teamRepo, teamMemberRepo)
	meetingController := controllers.NewMeetingController(meetingService)

	userController := controllers.NewUserController()

	teamController := controllers.NewTeamController()

	auth := v1.Group("/auth") // Create an /auth/ group
	{
		// Define the user registration route
		auth.POST("/register", userController.RegisterUser)

		// Define the user login route
		auth.POST("/login", userController.Login)

		// Verify user account by providing otp
		auth.POST("/verify", userController.VerifyEmail)

		// Request another verification email
		auth.GET("/request-verification", userController.RequestVerificationAgain)

		// Send forgot password request
		auth.GET("/forgot-password", userController.ForgotPasswordRequest)

		// Set forgotten password
		auth.POST("/set-forgotten-password", userController.SetNewPassword)

		// Test baseauth middleware
		auth.GET("/test-auth", middleware.BaseAuthMiddleware(), userController.TestAuth)

		// Reset password by logged in user
		auth.POST("/reset-password", middleware.BaseAuthMiddleware(), userController.ResetPassword)

		// Send account deletion request
		auth.GET("/request-delete-account", middleware.BaseAuthMiddleware(), userController.RequestDeletion)

		// Delete account
		auth.DELETE("/delete-account", middleware.BaseAuthMiddleware(), userController.DeleteAccount)

		// Google login
		auth.GET("/google/login", userController.GoogleLogin)

		// Google Callback
		auth.GET("/google/verify", userController.GoogleCallback)
	}

	user := v1.Group("/user")
	{
		// Get my details
		user.GET("/me", middleware.BaseAuthMiddleware(), userController.GetMyDetails)

		// Update my details
		user.PATCH("/me", middleware.BaseAuthMiddleware(), userController.UpdateMyDetails)

		// Get my teams, query ?role=member/admin/super_admin
		user.GET("/me/teams", middleware.BaseAuthMiddleware(), userController.GetMyTeams)

		// Get my upcoming meetings
		user.GET("/me/meetings", middleware.BaseAuthMiddleware(), meetingController.UpcomingUserMeetings)

		// Get my team requests, query ?status=accepted/rejected/pending
		user.GET("/me/requests", middleware.BaseAuthMiddleware(), userController.GetMyRequests)

		// Get past user attendance, TODO: filterable by team
		user.GET("/me/attendance", middleware.BaseAuthMiddleware(), meetingController.GetUserAttendanceRecords)
	}

	team := v1.Group("/team")
	{
		// Get all unprotected teams
		team.GET("/", middleware.BaseAuthMiddleware(), teamController.GetUnprotectedTeams)

		// Create a team
		team.POST("/", middleware.BaseAuthMiddleware(), teamController.CreateTeam)

		// Get a team by invite code
		team.GET("/invite/:inviteCode", middleware.BaseAuthMiddleware(), teamController.GetTeamByInviteCode)

		// Update a team
		team.PATCH("/:teamID", middleware.BaseAuthMiddleware(), middleware.AuthorizeSuperAdmin(), teamController.UpdateTeam)

		// Regenerate invite code
		team.GET("/:teamID/regenerate-invite", middleware.BaseAuthMiddleware(), middleware.AuthorizeSuperAdmin(), teamController.RegenerateInviteCode)

		// Join a team
		team.POST("/invite/:inviteCode/join", middleware.BaseAuthMiddleware(), teamController.JoinTeamByInviteCode)

		// Get team requests, query ?status=accepted/rejected/pending
		team.GET("/:teamID/requests", middleware.BaseAuthMiddleware(), middleware.AuthorizeAdmin(), teamController.GetTeamRequests)

		// Accept or reject a request Patch /team/:teamID/requests/:requestID by admin/superadmin
		team.PATCH("/:teamID/requests/:requestID", middleware.BaseAuthMiddleware(), middleware.AuthorizeAdmin(), teamController.UpdateTeamRequestStatus)

		// same as /team/invite/:inviteCode, for uniformity with next routes
		team.GET("/:teamID", middleware.BaseAuthMiddleware(), middleware.AuthorizeMember(), teamController.GetTeamByID)

		// get current user's role in a team
		team.GET("/:teamID/myrole", middleware.BaseAuthMiddleware(), middleware.AuthorizeMember(), teamController.GetCurrentUserRoleInTeam)

		// get team members, query ?role=member/admin/super_admin
		team.GET("/:teamID/members", middleware.BaseAuthMiddleware(), middleware.AuthorizeMember(), teamController.GetTeamMembers)

		// promote/demote a member to admin/member, visible to superadmin only
		team.PATCH("/:teamID/members/:memberID", middleware.BaseAuthMiddleware(), middleware.AuthorizeSuperAdmin(), teamController.PromoteOrDemoteTeamMember)

		// leave a team
		team.DELETE("/:teamID/leave", middleware.BaseAuthMiddleware(), middleware.AuthorizeMember(), teamController.LeaveTeam)

		// kick a member
		team.DELETE("/:teamID/members/:memberID", middleware.BaseAuthMiddleware(), middleware.AuthorizeSuperAdmin(), teamController.KickTeamMember)

		// handover superadmin to another member
		team.PATCH("/:teamID/handover", middleware.BaseAuthMiddleware(), middleware.AuthorizeSuperAdmin(), teamController.HandoverTeamSuperAdmin)

		// /:teamID/meetings to create one, by super admin
		team.POST("/:teamID/meetings", middleware.BaseAuthMiddleware(), middleware.AuthorizeSuperAdmin(), meetingController.CreateMeeting)

		// get all meetings of a team, query ?meetingOver=true/false
		team.GET("/:teamID/meetings", middleware.BaseAuthMiddleware(), middleware.AuthorizeMember(), meetingController.GetMeetingsByTeamID)

		// get a meeting by id
		team.GET("/:teamID/meetings/:meetingID", middleware.BaseAuthMiddleware(), middleware.AuthorizeMember(), meetingController.GetMeetingDetails)

		// start a meeting
		team.PATCH("/:teamID/meetings/:meetingID/start", middleware.BaseAuthMiddleware(), middleware.AuthorizeAdmin(), meetingController.StartMeeting)

		// end a meeting
		team.PATCH("/:teamID/meetings/:meetingID/end", middleware.BaseAuthMiddleware(), middleware.AuthorizeAdmin(), meetingController.EndMeeting)

		// start attendance
		team.PATCH("/:teamID/meetings/:meetingID/attendance/start", middleware.BaseAuthMiddleware(), middleware.AuthorizeAdmin(), meetingController.StartAttendance)

		// end attendance
		team.PATCH("/:teamID/meetings/:meetingID/attendance/end", middleware.BaseAuthMiddleware(), middleware.AuthorizeAdmin(), meetingController.EndAttendance)

		// delete a meeting
		team.DELETE("/:teamID/meetings/:meetingID", middleware.BaseAuthMiddleware(), middleware.AuthorizeSuperAdmin(), meetingController.DeleteMeetingByID)

		// mark attendance for a user in a meeting
		team.PATCH("/:teamID/meetings/:meetingID/attendance", middleware.BaseAuthMiddleware(), middleware.AuthorizeMember(), meetingController.MarkAttendance)

		// admin get attendance for a meeting
		team.GET("/:teamID/meetings/:meetingID/attendance", middleware.BaseAuthMiddleware(), middleware.AuthorizeAdmin(), meetingController.GetAttendanceForMeeting)
	}
}

// TODO delete team
// TODO controller-service-repo pattern
// TODO unit of work pattern
// For google oauth, make slight change. Instead of redirecting to the callback on backend directly, redirect to frontend url (or app uri), and have a route which accepts the auth code that frontend sends and does wht my callback is doing rn.
// TODO maybe some event broker like kafka for notifs
