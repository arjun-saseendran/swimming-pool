package models

import "time"

type User struct {
	ID        int       `json:"_id"`
	Name      string    `json:"name:"`
	Mobile    string      `json:"mobile"`
	CreatedAt time.Time `json:"created_at"`
}
