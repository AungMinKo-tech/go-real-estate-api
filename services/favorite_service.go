package services

import (
	"errors"
	"real-estate-api/config"
	"real-estate-api/models"

	"gorm.io/gorm"
)

func ToggleFavorite(userID uint, propertyID uint) (string, error) {
	var property models.Property
	if err := config.DB.First(&property, propertyID).Error; err != nil {
		return "", errors.New("property not found")
	}

	var favorite models.Favorite

	err := config.DB.Where("user_id = ? AND property_id = ?", userID, propertyID).First(&favorite).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		newFav := models.Favorite{UserID: userID, PropertyID: propertyID}
		if err := config.DB.Create(&newFav).Error; err != nil {
			return "", err
		}
		return "Added to favorites", nil
	} else if err == nil {
		if err := config.DB.Delete(&favorite).Error; err != nil {
			return "", err
		}
		return "Removed from favorites", nil
	}

	return "", err
}

// GetUserFavorites
func GetUserFavorites(userID uint) ([]models.Favorite, error) {
	var favorites []models.Favorite

	// Preload("Property.Images")
	err := config.DB.Preload("Property.Images").Where("user_id = ?", userID).Find(&favorites).Error
	return favorites, err
}
