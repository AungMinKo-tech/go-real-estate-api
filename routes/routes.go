package routes

import (
	"real-estate-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(route *gin.Engine) {

	auth := controllers.AuthController{}

	route.POST("api/v1/register", auth.Register)
	route.POST("api/v1/login", auth.Login)

	route.GET("api/v1/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Real Estate Api is running",
		})
	})
}
