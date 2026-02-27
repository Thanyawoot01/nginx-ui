package middleware

import (
	"net/http"

	"github.com/0xJacky/Nginx-UI/model"
	"github.com/gin-gonic/gin"
)

// RequireRole checks if the current user has one of the allowed roles
func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ensure user is authenticated first
		u, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		userModel, ok := u.(*model.User)
		if !ok || userModel == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid user context"})
			return
		}

		role := userModel.Role
		if role == "" {
			role = model.RoleAdmin
		}

		// Admin always has access to everything
		if role == model.RoleAdmin {
			c.Next()
			return
		}

		// Check if user's role is in the allowed list
		isAllowed := false
		for _, r := range allowedRoles {
			if role == r {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden: You do not have permission to access this resource"})
			return
		}

		c.Next()
	}
}
