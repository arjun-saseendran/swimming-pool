package models

import "time"

type BookedStatus string

const (
	Pending   BookedStatus = "PENDING"
	Booked    BookedStatus = "BOOKED"
	Canceled  BookedStatus = "CANCELED"
	Completed BookedStatus = "COMPLETED"
)

type Booking struct {
	ID        int          `json:"_id"`
	Title     string       `json:"title"`
	UserID    int          `json:"user_id"`
	Status    BookedStatus `json:"status"`
	CreatedAt time.Time    `json:"created_at"`
}
