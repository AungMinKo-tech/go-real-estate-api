package controllers

import (
	"net/http"
	"real-estate-api/services"

	"github.com/gin-gonic/gin"
)

type NotificationController struct{}

// GetMyNotifications - GET /api/v1/notifications
func (n NotificationController) GetMyNotifications(c *gin.Context) {
	userID, _ := c.Get("userID")

	notifs, err := services.GetUserNotifications(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notifications"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": notifs})
}

// ReadAll - PUT /api/v1/notifications/read
func (n NotificationController) ReadAll(c *gin.Context) {
	userID, _ := c.Get("userID")

	if err := services.MarkAsRead(userID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update notifications"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All notifications marked as read"})
}
