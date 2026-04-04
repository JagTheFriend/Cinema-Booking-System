package payment

import (
	"context"

	db "cinema_booking/internals/postgres/generated"
)

type PaymentStore interface {
	CreatePayment(ctx context.Context, arg *db.CreatePaymentParams) (*db.Payment, error)
	GetPaymentByID(ctx context.Context, id string) (*db.Payment, error)
	ListPaymentsByUser(ctx context.Context, userID string) ([]*db.Payment, error)
	UpdatePayment(ctx context.Context, id string) (*db.Payment, error)
	DeletePayment(ctx context.Context, id string) error
}
