package booking

import (
	"context"

	db "cinema_booking/internals/postgres/generated"

	glide "github.com/valkey-io/valkey-glide/go/v2"
)

type IBookingStore interface {
	CreateBooking(bookingData *db.CreateBookingParams) (*db.CreateBookingRow, error)
	GetBookingByID(bookingID string) (*db.GetBookingByIDRow, error)
	ListBookingsByUser(userID string) ([]*db.ListBookingsByUserRow, error)
	VerifyBooking(bookingID string) (*db.VerifyBookingRow, error)
	DeleteBooking(bookingID string) error
	UpdateBookingPayment(params *db.UpdateBookingPaymentParams) (*db.UpdateBookingPaymentRow, error)
}

type BookingStore struct {
	db     *db.Queries
	valkey *glide.Client
	ctx    context.Context
}

func NewBookingStore(ctx context.Context, db *db.Queries, valkey *glide.Client) IBookingStore {
	return &BookingStore{
		db:     db,
		valkey: valkey,
		ctx:    ctx,
	}
}

// CreateBooking inserts a new booking into the database
func (s *BookingStore) CreateBooking(bookingData *db.CreateBookingParams) (*db.CreateBookingRow, error) {
	return s.db.CreateBooking(s.ctx, bookingData)
}

// GetBookingByID fetches a booking by its ID
func (s *BookingStore) GetBookingByID(bookingID string) (*db.GetBookingByIDRow, error) {
	return s.db.GetBookingByID(s.ctx, bookingID)
}

// ListBookingsByUser fetches all bookings for a given user
func (s *BookingStore) ListBookingsByUser(userID string) ([]*db.ListBookingsByUserRow, error) {
	return s.db.ListBookingsByUser(s.ctx, userID)
}

// VerifyBooking marks a booking as verified
func (s *BookingStore) VerifyBooking(bookingID string) (*db.VerifyBookingRow, error) {
	return s.db.VerifyBooking(s.ctx, bookingID)
}

// DeleteBooking removes a booking from the database
func (s *BookingStore) DeleteBooking(bookingID string) error {
	return s.db.DeleteBooking(s.ctx, bookingID)
}

// UpdateBookingPayment updates the payment ID for a booking
func (s *BookingStore) UpdateBookingPayment(params *db.UpdateBookingPaymentParams) (*db.UpdateBookingPaymentRow, error) {
	return s.db.UpdateBookingPayment(s.ctx, params)
}
