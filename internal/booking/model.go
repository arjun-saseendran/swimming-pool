package booking

import "time"

type BookedStatus string

const (
	Booked    BookedStatus = "BOOKED"
	Canceled  BookedStatus = "CANCELED"
	Completed BookedStatus = "COMPLETED"
)

type Booking struct {
	ID        uint         `gorm:"primaryKey" json:"id"`
	Title     string       `json:"title"`
	UserID    uint         `json:"user_id"`
	Status    BookedStatus `gorm:"default:'BOOKED'" json:"status"`
	CreatedAt time.Time    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
}
