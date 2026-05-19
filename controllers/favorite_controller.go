package controllers

import (
	"net/http"
	"real-estate-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FavoriteController struct{}

// Toggle - POST /api/v1/favorites/:property_id
func (f FavoriteController) Toggle(c *gin.Context) {
	userID, _ := c.Get("userID")

	propertyIDStr := c.Param("property_id")
	propertyID, _ := strconv.Atoi(propertyIDStr)

	message, err := services.ToggleFavorite(userID.(uint), uint(propertyID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}

// GetMyFavorites - GET /api/v1/favorites
func (f FavoriteController) GetMyFavorites(c *gin.Context) {
	userID, _ := c.Get("userID")

	favorites, err := services.GetUserFavorites(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch favorites"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": favorites})
}
