package routes

import (
	"real-estate-api/controllers"
	"real-estate-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(route *gin.Engine) {
	auth := controllers.AuthController{}
	propertyController := controllers.PropertyController{}
	favoriteController := controllers.FavoriteController{}

	// Public Routes
	v1 := route.Group("/api/v1")
	{
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Real Estate Api is running"})
		})

		v1.POST("/register", auth.Register)
		v1.POST("/login", auth.Login)

		v1.GET("/properties", propertyController.GetAll)
		route.Static("/uploads", "./uploads")
	}

	// Protected Routes
	protected := v1.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		v1.GET("/properties/:id", propertyController.GetByID)

		v1.POST("/favorites/:property_id", favoriteController.Toggle)
		v1.GET("/favorites", favoriteController.GetMyFavorites)

		// Admin Only Routes
		adminOnly := protected.Group("/")
		adminOnly.Use(middleware.RoleMiddleware("admin"))
		{
			//
		}

		// Staff Only Routes
		staffOnly := protected.Group("/")
		staffOnly.Use(middleware.RoleMiddleware("agent", "admin"))
		{
			staffOnly.POST("/properties", propertyController.Create)
			staffOnly.PUT("/properties/:id", propertyController.Update)
			staffOnly.DELETE("/properties/:id", propertyController.Delete)
			staffOnly.POST("/properties/:id/images", propertyController.UploadImages)
		}
	}
}
