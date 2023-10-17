package routers

import (
	"net/http"

	"github.com/GDGVIT/attendance-app-backend/controllers"
	"github.com/GDGVIT/attendance-app-backend/routers/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes add all routing list here automatically get main router
func RegisterRoutes(route *gin.Engine) {
	route.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})
	route.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"live": "ok"}) })

	v1 := route.Group("/v1")

	auth := v1.Group("/auth") // Create an /auth/ group
	{
		userController := controllers.NewUserController() // Create an instance of the UserController

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
		auth.GET("/google/callback", userController.GoogleCallback)
	}

	user := v1.Group("/user")
	{
		userController := controllers.NewUserController()

		// Get my details
		user.GET("/me", middleware.BaseAuthMiddleware(), userController.GetMyDetails)

		// Update my details
		user.PATCH("/me", middleware.BaseAuthMiddleware(), userController.UpdateMyDetails)
	}

	team := v1.Group("/team")
	{
		teamController := controllers.NewTeamController()

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
	}

	// TODO Get /me/teams teams user is in
	// TODO Get /me/requests requests user has sent
	// TODO Get /team/:teamID/requests requests for a team, visible to admin/superadmin
	// TODO Patch /team/:teamID/requests/:requestID accept/reject a request, visible to admin/superadmin
	// TODO Get /team/:teamID team info, visible to all team members (similar to /team/invite/:inviteCode)
	// TODO Get /team/:teamID/members team members, visible to all team members
	// TODO Patch /team/:teamID/members/:memberID?promote=true change member role to member/admin, visible to superadmin only

	// non critical
	// TODO Get /team all unprotected teams
}
