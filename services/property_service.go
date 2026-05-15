package services

import (
	"errors"
	"real-estate-api/config"
	"real-estate-api/models"
	"real-estate-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type CreatePropertyInput struct {
	Title           string  `json:"title" binding:"required,min=5,max=255"`
	Description     string  `json:"description"`
	TransactionType string  `json:"transaction_type" binding:"required,oneof=sale rent"`
	PropertyType    string  `json:"property_type" binding:"required"`
	Price           float64 `json:"price" binding:"required,gt=0"`
	Currency        string  `json:"currency"`
	Bedrooms        int     `json:"bedrooms" binding:"gte=0"`
	Bathrooms       int     `json:"bathrooms" binding:"gte=0"`
	Area            int     `json:"area" binding:"gte=0"`
	Address         string  `json:"address"`
	City            string  `json:"city" binding:"required"`
}

type UpdatePropertyInput struct {
	Title           string  `json:"title" binding:"required,min=5,max=255"`
	Description     string  `json:"description"`
	TransactionType string  `json:"transaction_type" binding:"required,oneof=sale rent"`
	PropertyType    string  `json:"property_type" binding:"required"`
	Price           float64 `json:"price" binding:"required,gt=0"`
	Currency        string  `json:"currency"`
	Bedrooms        int     `json:"bedrooms" binding:"gte=0"`
	Bathrooms       int     `json:"bathrooms" binding:"gte=0"`
	Area            int     `json:"area" binding:"gte=0"`
	Address         string  `json:"address"`
	City            string  `json:"city" binding:"required"`
}

var (
	ErrUnauthorized = errors.New("unauthorized: you can only update your own property")
	ErrNotFound     = errors.New("record not found")
)

func CreateProperty(userID uint, input CreatePropertyInput) (*models.Property, error) {
	property := models.Property{
		UserID:          userID,
		Title:           input.Title,
		Description:     input.Description,
		TransactionType: input.TransactionType,
		PropertyType:    input.PropertyType,
		Price:           decimal.NewFromFloat(input.Price),
		Currency:        input.Currency,
		Bedrooms:        input.Bedrooms,
		Bathrooms:       input.Bathrooms,
		Area:            input.Area,
		Address:         input.Address,
		City:            input.City,
	}

	if property.Currency == "" {
		property.Currency = "MMK"
	}

	if err := config.DB.Create(&property).Error; err != nil {
		return nil, err
	}

	return &property, nil
}

// GetAllProperties
func GetAllProperties(c *gin.Context) (*utils.Pagination, error) {
	var properties []models.Property

	query := config.DB.Model(&models.Property{})

	if city := c.Query("city"); city != "" {
		query = query.Where("city LIKE ?", "%"+city+"%")
	}

	if minPrice := c.Query("min_price"); minPrice != "" {
		query = query.Where("price >= ?", minPrice)
	}

	if maxPrice := c.Query("max_price"); maxPrice != "" {
		query = query.Where("price <= ?", maxPrice)
	}

	if bedrooms := c.Query("bedrooms"); bedrooms != "" {
		query = query.Where("bedrooms = ?", bedrooms)
	}

	return utils.Paginate(c, query, &models.Property{}, &properties)
}

// GetPropertyByID
func GetPropertyByID(id string) (models.Property, error) {
	var property models.Property
	err := config.DB.First(&property, id).Error
	return property, err
}

func UpdateProperty(id string, userID uint, userRole string, input UpdatePropertyInput) (models.Property, error) {
	var property models.Property
	if err := config.DB.First(&property, id).Error; err != nil {
		return property, ErrNotFound
	}

	if property.UserID != userID && userRole != "admin" {
		return property, ErrUnauthorized
	}

	config.DB.Model(&property).Updates(input)
	return property, nil
}

// Delete
func DeleteProperty(id string, userID uint, userRole string) error {
	var property models.Property
	if err := config.DB.First(&property, id).Error; err != nil {
		return ErrNotFound
	}

	if property.UserID != userID && userRole != "admin" {
		return ErrUnauthorized
	}

	return config.DB.Delete(&property).Error
}
