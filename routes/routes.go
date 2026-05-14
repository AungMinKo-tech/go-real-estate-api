package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(route *gin.Engine) {
	route.GET("api/v1/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Real Estate Api is running",
		})
	})
}
