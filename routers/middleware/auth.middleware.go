package middleware

import (
	"net/http"
	"strconv"

	"github.com/GDGVIT/attendance-app-backend/infra/logger"
	"github.com/GDGVIT/attendance-app-backend/models"
	"github.com/GDGVIT/attendance-app-backend/repository"
	"github.com/GDGVIT/attendance-app-backend/utils/token"
	"github.com/gin-gonic/gin"
)

// BaseAuthMiddleware checks if the user is authenticated
func BaseAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		userID, err := token.ValidateToken(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "auth", "message": "Please login to continue."})
			// logger.Errorf("Auth Middleware Error: %v", err)
			c.Abort()
			return
		}

		var user models.User

		userRepo := repository.NewUserRepository()
		user, err = userRepo.GetUserByID(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user-fetch-error", "message": "Internal error while fetching user."})
			logger.Errorf("Fetching Authenticated User Error: %v", err)
			c.Abort()
			return
		}

		c.Set("user", &user)

		c.Next()
	}
}

func AuthorizeSuperAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the user is a super_admin
		user, _ := c.Get("user") // Assuming you have user information in the context
		teamID, _ := strconv.Atoi(c.Param("teamID"))
		// Check the user's role and permissions for the team
		teamMemberRepo := repository.NewTeamMemberRepository()
		teamMember, err := teamMemberRepo.GetTeamMemberByID(uint(teamID), user.(*models.User).ID)
		if err == nil && teamMember.Role == models.SuperAdminRole {
			// User is authorized as superadmin for given team, proceed to the next handler
			c.Next()
		} else {
			// User is not authorized, return an error response
			c.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to do that action on this team."})
			c.Abort()
		}
	}
}

// AuthorizeAdmin checks if the user is a super_admin or admin for the given team
func AuthorizeAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("user") // Assuming you have user information in the context
		teamID, _ := strconv.Atoi(c.Param("teamID"))
		// Check the user's role and permissions for the team
		teamMemberRepo := repository.NewTeamMemberRepository()
		teamMember, err := teamMemberRepo.GetTeamMemberByID(uint(teamID), user.(*models.User).ID)
		if err == nil && (teamMember.Role == models.SuperAdminRole || teamMember.Role == models.AdminRole) {
			c.Next()
		} else {
			// User is not authorized, return an error response
			c.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to do that action on this team."})
			c.Abort()
		}
	}
}

// AuthorizeMember checks if the user is a super_admin, admin or member for the given team
func AuthorizeMember() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("user") // Assuming you have user information in the context
		teamID, _ := strconv.Atoi(c.Param("teamID"))
		// Check the user's role and permissions for the team
		teamMemberRepo := repository.NewTeamMemberRepository()
		teamMember, err := teamMemberRepo.GetTeamMemberByID(uint(teamID), user.(*models.User).ID)
		if err == nil && (teamMember.Role == models.SuperAdminRole || teamMember.Role == models.AdminRole || teamMember.Role == models.MemberRole) {
			c.Next()
		} else {
			// User is not authorized, return an error response
			c.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to do that action on this team."})
			c.Abort()
		}
	}
}
