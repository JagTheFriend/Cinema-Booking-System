package booking

type CreateBookingRequest struct {
	SeatID    string  `json:"seatId" validate:"required"`
	MovieID   string  `json:"movieId" validate:"required"`
	PaymentID *string `json:"paymentId"`
}

type UpdateBookingPaymentRequest struct {
	ID        string  `json:"id" validate:"required"`
	PaymentID *string `json:"paymentId" validate:"required"`
}
