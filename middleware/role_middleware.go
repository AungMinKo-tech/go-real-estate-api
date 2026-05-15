package middleware

import (
	"github.com/gin-gonic/gin"
)

func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, _ := c.Get("role")
		roleStr := userRole.(string)

		for _, r := range allowedRoles {
			if r == roleStr {
				c.Next()
				return
			}
		}
		c.JSON(403, gin.H{"error": "Forbidden"})
		c.Abort()
	}
}
