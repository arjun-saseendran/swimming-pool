package booking

import "gorm.io/gorm"

type BookingService interface {
	Create(bookingData *InputBooking) (*Booking, error)
}

type bookingService struct{ db *gorm.DB }

func NewBookingService(db *gorm.DB) BookingService {
	return &bookingService{db: db}
}

func (bs *bookingService) Create(bookingData *InputBooking) (*Booking, error) {
	newBooking := &Booking{Title: bookingData.Title, UserID: bookingData.UserID}
	result := bs.db.Create(newBooking)
	if result.Error != nil {
		return nil, result.Error
	}
	return newBooking, nil
}
