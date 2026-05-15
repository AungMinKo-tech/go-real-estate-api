package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Property struct {
	gorm.Model

	UserID uint `json:"user_id"`

	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description"`

	TransactionType string `json:"transaction_type" gorm:"type:varchar(20);not null"` // sale or rent
	PropertyType    string `json:"property_type" gorm:"type:varchar(20);not null"`    // house or condo

	Price    decimal.Decimal `json:"price" gorm:"type:numeric(15,2); not null"`
	Currency string          `json:"currency" gorm:"default:MMK"`

	Bedrooms  int `json:"bedrooms"`
	Bathrooms int `json:"bathrooms"`
	Area      int `json:"area"` // square feet

	Address string `json:"address"`
	City    string `json:"city"`

	Status     string `json:"status" gorm:"default:active"` // active, sold, rented, inactive
	IsFeatured bool   `json:"is_featured" gorm:"default:false"`

	CreatedAt time.Time

	User User `gorm:"foreignKey:UserID"`
}
