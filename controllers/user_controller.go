package controllers

import (
	"net/http"
	"real-estate-api/services"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

type UpdateProfileInput struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type BanInput struct {
	UserID   uint `json:"user_id" binding:"required"`
	IsBanned bool `json:"is_banned"`
}

// GetProfile - GET /api/v1/users/profile
func (u UserController) GetProfile(c *gin.Context) {
	userID, _ := c.Get("userID")

	user, err := services.GetProfileByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// UpdateProfile - PUT /api/v1/users/profile
func (u UserController) UpdateProfile(c *gin.Context) {
	userID, _ := c.Get("userID")

	var input UpdateProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := services.UpdateUserProfile(userID.(uint), input.Name, input.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully", "data": user})
}

// AdminGetAllUsers - GET /api/v1/admin/users
func (u UserController) AdminGetAllUsers(c *gin.Context) {
	users, err := services.GetAllUsersList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (u UserController) AdminUpdateBanStatus(c *gin.Context) {
	var input BanInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.ToggleUserBanStatus(input.UserID, input.IsBanned)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	action := "unbanned"
	if input.IsBanned {
		action = "banned"
	}

	c.JSON(http.StatusOK, gin.H{"message": "User has been successfully " + action})
}
