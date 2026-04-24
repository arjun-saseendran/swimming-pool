package booking

type InputBooking struct {
	Title  string `json:"title" binding:"required"`
	UserID uint   `json:"user_id" binding:"required"`
}

type UpdateInputBooking struct {
	Title  string       `json:"title"`
	Status BookedStatus `json:"status" binding:"omitempty,oneof=BOOKED CANCELED COMPLETED"`
}

func NewInputBooking() *InputBooking {
	return &InputBooking{}
}

func NewUpdateInputBooking() *UpdateInputBooking {
	return &UpdateInputBooking{}
}
