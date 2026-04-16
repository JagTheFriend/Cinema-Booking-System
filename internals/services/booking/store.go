package booking

import (
	"context"
	"errors"

	db "cinema_booking/internals/postgres/generated"
	"cinema_booking/internals/valkey"
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
	db  *db.Queries
	ctx context.Context
}

func NewBookingStore(ctx context.Context, db *db.Queries) IBookingStore {
	return &BookingStore{
		db:  db,
		ctx: ctx,
	}
}

// CreateBooking
func (s *BookingStore) CreateBooking(bookingData *db.CreateBookingParams) (*db.CreateBookingRow, error) {
	// 1. Lock seat in Valkey
	err := valkey.AddBooking(s.ctx, &valkey.Booking{
		Status:    "pending",
		SeatID:    bookingData.SeatID,
		MovieID:   bookingData.MovieID,
		UserID:    bookingData.UserID,
		PaymentID: "",
	})
	if err != nil {
		return nil, err
	}

	// 2. Persist in DB
	result, err := s.db.CreateBooking(s.ctx, bookingData)
	if err != nil {
		// rollback cache
		_ = valkey.DeleteBooking(s.ctx, "pending", bookingData.SeatID, bookingData.MovieID)
		return nil, err
	}

	return result, nil
}

// GetBookingByID (DB is source of truth)
func (s *BookingStore) GetBookingByID(bookingID string) (*db.GetBookingByIDRow, error) {
	return s.db.GetBookingByID(s.ctx, bookingID)
}

// ListBookingsByUser
func (s *BookingStore) ListBookingsByUser(userID string) ([]*db.ListBookingsByUserRow, error) {
	return s.db.ListBookingsByUser(s.ctx, userID)
}

// VerifyBooking
func (s *BookingStore) VerifyBooking(bookingID string) (*db.VerifyBookingRow, error) {
	// fetch booking for cache key
	booking, err := s.db.GetBookingByID(s.ctx, bookingID)
	if err != nil {
		return nil, err
	}

	// update DB
	result, err := s.db.VerifyBooking(s.ctx, bookingID)
	if err != nil {
		return nil, err
	}

	// remove from Valkey
	_ = valkey.DeleteBooking(s.ctx, "pending", booking.SeatID, booking.MovieID)

	return result, nil
}

// DeleteBooking
func (s *BookingStore) DeleteBooking(bookingID string) error {
	booking, err := s.db.GetBookingByID(s.ctx, bookingID)
	if err != nil {
		return err
	}

	err = s.db.DeleteBooking(s.ctx, bookingID)
	if err != nil {
		return err
	}

	// cleanup Valkey
	_ = valkey.DeleteBooking(s.ctx, "pending", booking.SeatID, booking.MovieID)

	return nil
}

// UpdateBookingPayment
func (s *BookingStore) UpdateBookingPayment(params *db.UpdateBookingPaymentParams) (*db.UpdateBookingPaymentRow, error) {
	if params.PaymentID == nil {
		return nil, errors.New("payment ID cannot be nil")
	}

	// fetch booking for cache key
	booking, err := s.db.GetBookingByID(s.ctx, params.ID)
	if err != nil {
		return nil, err
	}

	// update DB
	result, err := s.db.UpdateBookingPayment(s.ctx, params)
	if err != nil {
		return nil, err
	}

	// update Valkey using util
	_ = valkey.UpdateBookingPayment(
		s.ctx,
		"pending",
		booking.SeatID,
		booking.MovieID,
		*params.PaymentID,
	)

	return result, nil
}
