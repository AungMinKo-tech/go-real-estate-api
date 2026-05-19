package services

import (
	"real-estate-api/config"
	"real-estate-api/models"
)

// SendNotification
func SendNotification(userID uint, title string, message string, notifType string) error {
	notif := models.Notification{
		UserID:  userID,
		Title:   title,
		Message: message,
		Type:    notifType,
	}
	if err := config.DB.Create(&notif).Error; err != nil {
		return err
	}

	GlobalHub.Mu.Lock()
	conn, online := GlobalHub.Clients[userID]
	GlobalHub.Mu.Unlock()

	if online {
		payload := map[string]interface{}{
			"event": "notification",
			"data":  notif,
		}
		conn.WriteJSON(payload)
	}

	return nil
}

// GetUserNotifications
func GetUserNotifications(userID uint) ([]models.Notification, error) {
	var notifs []models.Notification
	err := config.DB.Where("user_id = ?", userID).Order("created_at DESC").Find(&notifs).Error
	return notifs, err
}

// MarkAsRead
func MarkAsRead(userID uint) error {
	return config.DB.Model(&models.Notification{}).Where("user_id = ? AND is_read = ?", userID, false).Update("is_read", true).Error
}
