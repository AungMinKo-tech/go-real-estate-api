package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string `json:"name"`
	Email        string `json:"email" gorm:"uniqueIndex; not null"`
	PasswordHash string `json:"-"`
	Phone        string `json:"phone"`
	Role         string `json:"role" gorm:"default:user"` // user/agent/admin
	IsBanned     bool   `json:"is_banned" gorm:"default:false"`
}
