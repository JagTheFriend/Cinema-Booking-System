package payment

import (
	"context"

	db "cinema_booking/internals/postgres/generated"
)

type PaymentStore interface {
	CreatePayment(arg *db.CreatePaymentParams) (*db.Payment, error)
	GetPaymentByID(id string) (*db.Payment, error)
	ListPaymentsByUser(userID string) ([]*db.Payment, error)
	UpdatePayment(id string) (*db.Payment, error)
	DeletePayment(id string) error
}

type PaymentStoreImpl struct {
	db *db.Queries
}

func NewPaymentStore(db *db.Queries) PaymentStore {
	return &PaymentStoreImpl{
		db: db,
	}
}

// CreatePayment inserts a new payment into the database
func (s *PaymentStoreImpl) CreatePayment(arg *db.CreatePaymentParams) (*db.Payment, error) {
	return s.db.CreatePayment(context.Background(), arg)
}

// GetPaymentByID fetches a payment by its ID
func (s *PaymentStoreImpl) GetPaymentByID(id string) (*db.Payment, error) {
	return s.db.GetPaymentByID(context.Background(), id)
}

// ListPaymentsByUser fetches all payments for a given user
func (s *PaymentStoreImpl) ListPaymentsByUser(userID string) ([]*db.Payment, error) {
	return s.db.ListPaymentsByUser(context.Background(), userID)
}

// UpdatePayment updates the updated_at timestamp of a payment
func (s *PaymentStoreImpl) UpdatePayment(id string) (*db.Payment, error) {
	return s.db.UpdatePayment(context.Background(), id)
}

// DeletePayment removes a payment from the database
func (s *PaymentStoreImpl) DeletePayment(id string) error {
	return s.db.DeletePayment(context.Background(), id)
}
