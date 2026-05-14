package controllers

import (
	"net/http"
	"real-estate-api/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (a AuthController) Register(c *gin.Context) {
	var body struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Phone    string `json:"phone"`
	}

	c.BindJSON(&body)

	err := services.Register(body.Name, body.Email, body.Password, body.Phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered"})
}

func (a AuthController) Login(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	c.BindJSON(&body)

	result, err := services.Login(body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login success",
		"token":   result.Token,
		"user": gin.H{
			"id":    result.User.ID,
			"name":  result.User.Name,
			"email": result.User.Email,
			"phone": result.User.Phone,
			"role":  result.User.Role,
		},
	})
}
