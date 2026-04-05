package payment

type CreatePaymentRequest struct {
	BookingID string `json:"bookingId" validate:"required"`
}

type UpdatePaymentRequest struct {
	ID string `json:"id" validate:"required"`
}

type DeletePaymentRequest struct {
	ID string `json:"id" validate:"required"`
}
