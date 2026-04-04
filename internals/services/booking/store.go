package booking

import (
	"context"

	db "cinema_booking/internals/postgres/generated"
)

type BookingStore interface {
	CreateBooking(ctx context.Context, bookingData *db.CreateBookingParams) (*db.CreateBookingRow, error)
	GetBookingByID(ctx context.Context, bookingID string) (*db.GetBookingByIDRow, error)
	ListBookingsByUser(ctx context.Context, userID string) ([]*db.ListBookingsByUserRow, error)
	VerifyBooking(ctx context.Context, bookingID string) (*db.VerifyBookingRow, error)
	DeleteBooking(ctx context.Context, bookingID string) error
	UpdateBookingPayment(ctx context.Context, params *db.UpdateBookingPaymentParams) (*db.UpdateBookingPaymentRow, error)
}
