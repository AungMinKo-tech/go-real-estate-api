package models

import "time"

type PropertyImage struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	PropertyID uint      `json:"property_id"`
	ImagePath  string    `json:"image_path"`
	CreatedAt  time.Time `json:"created_at"`
}
