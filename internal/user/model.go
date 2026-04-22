package user

import "time"

type UserType string

const (
	Admin    UserType = "ADMIN"
	Customer UserType = "CUSTOMER"
)

type User struct {
	ID        int       `json:"_id,omitempty"`
	Name      string    `json:"name"`
	Mobile    string    `json:"mobile"`
	Role      UserType  `json:"user_type"`
	CreatedAt time.Time `json:"created_at"`
}
