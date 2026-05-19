package models

import "time"

type Inquiry struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	PropertyID uint      `gorm:"not null" json:"property_id"`
	Property   Property  `gorm:"foreignKey:PropertyID" json:"property,omitempty"`  // ဘယ်အိမ်လဲ
	UserID     uint      `gorm:"not null" json:"user_id"`                          // ပို့တဲ့သူ
	User       User      `gorm:"foreignKey:UserID" json:"user,omitempty"`          // ပို့တဲ့သူ့အချက်အလက်
	Message    string    `gorm:"type:text;not null" json:"message"`                // မေးချင်တဲ့ message
	Status     string    `gorm:"type:varchar(20);default:'pending'" json:"status"` // pending, contacted, closed
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
