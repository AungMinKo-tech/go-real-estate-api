package services

import (
	"errors"
	"real-estate-api/config"
	"real-estate-api/models"
	"real-estate-api/utils"

	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	User  models.User
	Token string
}

func Register(name, email, password, phone string) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	user := models.User{
		Name:         name,
		Email:        email,
		PasswordHash: string(hash),
		Phone:        phone,
	}

	return config.DB.Create(&user).Error
}

func Login(email, password string) (LoginResponse, error) {
	var user models.User

	err := config.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return LoginResponse{}, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return LoginResponse{}, errors.New("Invalid password!")
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return LoginResponse{}, err
	}

	return LoginResponse{
		User:  user,
		Token: token,
	}, nil
}
