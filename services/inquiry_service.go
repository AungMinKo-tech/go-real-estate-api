package services

import (
	"errors"
	"real-estate-api/config"
	"real-estate-api/models"
)

// CreateInquiry
func CreateInquiry(userID uint, propertyID uint, message string) (models.Inquiry, error) {
	var property models.Property

	if err := config.DB.First(&property, propertyID).Error; err != nil {
		return models.Inquiry{}, errors.New("property not found")
	}

	inquiry := models.Inquiry{
		PropertyID: propertyID,
		UserID:     userID,
		Message:    message,
	}

	if err := config.DB.Create(&inquiry).Error; err != nil {
		return models.Inquiry{}, err
	}

	go SendNotification(property.UserID, "New Property Inquiry", "Someone wants to check out your property listing.", "inquiry")

	return inquiry, nil
}

// GetAgentInquiries
func GetAgentInquiries(agentID uint) ([]models.Inquiry, error) {
	var inquiries []models.Inquiry

	err := config.DB.Preload("Property").Preload("User").
		Joins("JOIN properties ON properties.id = inquiries.property_id").
		Where("properties.agent_id = ?", agentID).
		Order("inquiries.created_at DESC").
		Find(&inquiries).Error

	return inquiries, err
}
