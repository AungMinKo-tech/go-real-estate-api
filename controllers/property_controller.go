package controllers

import (
	"errors"
	"net/http"

	"real-estate-api/services"

	"github.com/gin-gonic/gin"
)

type PropertyController struct{}

func (p PropertyController) Create(c *gin.Context) {
	var input services.CreatePropertyInput

	// Input Validation
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDValue, userExists := c.Get("user_id")
	roleValue, roleExists := c.Get("role")

	if !userExists || !roleExists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userRole := roleValue.(string)
	if userRole != "agent" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Access denied. Only agents can create properties.",
		})
		return
	}

	userID := userIDValue.(uint)

	property, err := services.CreateProperty(userID, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create property"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Property created successfully",
		"property": property,
	})
}

func (p PropertyController) GetAll(c *gin.Context) {
	properties, err := services.GetAllProperties()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch properties"})
		return
	}
	c.JSON(http.StatusOK, properties)
}

func (p PropertyController) GetByID(c *gin.Context) {
	id := c.Param("id")
	property, err := services.GetPropertyByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Property not found"})
		return
	}
	c.JSON(http.StatusOK, property)
}

func (p PropertyController) Update(c *gin.Context) {
	id := c.Param("id")
	currentUserID := c.MustGet("user_id").(uint)
	currentRole := c.MustGet("role").(string)

	var input services.UpdatePropertyInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProperty, err := services.UpdateProperty(id, currentUserID, currentRole, input)
	if err != nil {
		var status int
		switch {
		case errors.Is(err, services.ErrUnauthorized):
			status = http.StatusForbidden
		case errors.Is(err, services.ErrNotFound):
			status = http.StatusNotFound
		default:
			status = http.StatusInternalServerError
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedProperty)
}

func (p PropertyController) Delete(c *gin.Context) {
	id := c.Param("id")
	currentUserID := c.MustGet("user_id").(uint)
	currentRole := c.MustGet("role").(string)

	err := services.DeleteProperty(id, currentUserID, currentRole)
	if err != nil {
		if errors.Is(err, services.ErrUnauthorized) {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Property not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
