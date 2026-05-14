package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"real-estate-api/config"
	"real-estate-api/models"
	"real-estate-api/routes"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})

	route := gin.Default()
	routes.SetupRoutes(route)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	route.Run(":" + port)
}
