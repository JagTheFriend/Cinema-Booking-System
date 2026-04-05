package payment

type CreatePaymentDTO struct {
	BookingID string `json:"bookingId" validate:"required"`
}

type UpdatePaymentDTO struct {
	// extend later if needed
}
