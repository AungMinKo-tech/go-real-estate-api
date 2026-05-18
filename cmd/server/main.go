package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"real-estate-api/config"
	"real-estate-api/models"
	"real-estate-api/routes"
)

func main() {
	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		err := os.Mkdir("uploads", 0755)
		if err != nil {
			log.Fatal("Failed to create uploads directory:", err)
		}
	}

	_ = godotenv.Load()

	config.ConnectDB()
	config.DB.AutoMigrate(
		&models.User{},
		&models.Property{},
		&models.PropertyImage{},
	)

	route := gin.Default()
	routes.SetupRoutes(route)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	route.Run(":" + port)
}
