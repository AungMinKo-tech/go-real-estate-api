package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"real-estate-api/config"
	"real-estate-api/models"
	"real-estate-api/routes"
)

func main() {
	_ = godotenv.Load()

	config.ConnectDB()
	config.DB.AutoMigrate(
		&models.User{},
		&models.Property{},
	)

	route := gin.Default()
	routes.SetupRoutes(route)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	route.Run(":" + port)
}
