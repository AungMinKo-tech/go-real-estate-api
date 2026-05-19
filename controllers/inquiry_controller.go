package controllers

import (
	"net/http"
	"real-estate-api/services"

	"github.com/gin-gonic/gin"
)

type InquiryController struct{}

type InquiryInput struct {
	PropertyID uint   `json:"property_id" binding:"required"`
	Message    string `json:"message" binding:"required"`
}

// SendInquiry - POST /api/v1/inquiries
func (i InquiryController) SendInquiry(c *gin.Context) {
	userID, _ := c.Get("userID")

	var input InquiryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	inquiry, err := services.CreateInquiry(userID.(uint), input.PropertyID, input.Message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Inquiry sent successfully", "data": inquiry})
}

// GetMyLeads - GET /api/v1/inquiries/agent
func (i InquiryController) GetMyLeads(c *gin.Context) {
	agentID, _ := c.Get("userID")

	inquiries, err := services.GetAgentInquiries(agentID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch inquiries"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": inquiries})
}
