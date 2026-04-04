package payment

import (
	"context"

	db "cinema_booking/internals/postgres/generated"

	glide "github.com/valkey-io/valkey-glide/go/v2"
)

type PaymentStore interface {
	CreatePayment(arg *db.CreatePaymentParams) (*db.Payment, error)
	GetPaymentByID(id string) (*db.Payment, error)
	ListPaymentsByUser(userID string) ([]*db.Payment, error)
	UpdatePayment(id string) (*db.Payment, error)
	DeletePayment(id string) error
}

type PaymentStoreImpl struct {
	db     *db.Queries
	valkey *glide.Client
	ctx    context.Context
}

func NewPaymentStore(ctx context.Context, db *db.Queries, valkey *glide.Client) PaymentStore {
	return &PaymentStoreImpl{
		db:     db,
		valkey: valkey,
		ctx:    ctx,
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
