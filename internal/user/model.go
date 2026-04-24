package user

import "time"

type UserType string

const (
	Admin    UserType = "ADMIN"
	Customer UserType = "CUSTOMER"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Mobile    string    `gorm:"uniqueIndex" json:"mobile"`
	Role      UserType  `gorm:"default:'CUSTOMER'" json:"user_type"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
