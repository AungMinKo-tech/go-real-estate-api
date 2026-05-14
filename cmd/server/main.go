package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"real-estate-api/config"
	"real-estate-api/routes"
)

func main() {
	config.ConnectDB()

	route := gin.Default()
	routes.SetupRoutes(route)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	route.Run(":" + port)
}
