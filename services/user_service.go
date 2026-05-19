package services

import (
	"errors"
	"real-estate-api/config"
	"real-estate-api/models"
)

// GetProfileByID
func GetProfileByID(userID uint) (models.User, error) {
	var user models.User

	if err := config.DB.Omit("password").First(&user, userID).Error; err != nil {
		return user, errors.New("user not found")
	}

	return user, nil
}

// UpdateUserProfile
func UpdateUserProfile(userID uint, name string, phone string) (models.User, error) {
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return user, errors.New("user not found")
	}

	if name != "" {
		user.Name = name
	}
	if phone != "" {
		user.Phone = phone
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// GetAllUsersList
func GetAllUsersList() ([]models.User, error) {
	var users []models.User
	if err := config.DB.Omit("password").Order("id DESC").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// ToggleUserBanStatus
func ToggleUserBanStatus(targetUserID uint, banStatus bool) error {
	var user models.User
	if err := config.DB.First(&user, targetUserID).Error; err != nil {
		return errors.New("user not found")
	}

	if user.Role == "admin" {
		return errors.New("cannot ban an admin account")
	}

	user.IsBanned = banStatus
	return config.DB.Save(&user).Error
}
