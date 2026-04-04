// Package payment contains the payment routes.
package payment

import (
	db "cinema_booking/internals/postgres/generated"

	"github.com/labstack/echo/v5"
	glide "github.com/valkey-io/valkey-glide/go/v2"
)

type PaymentRoute struct {
	e      *echo.Group
	db     *db.Queries
	valkey *glide.Client
}

func NewPaymentRoute(e *echo.Group, db *db.Queries, valkey *glide.Client) *PaymentRoute {
	return &PaymentRoute{
		e:      e,
		db:     db,
		valkey: valkey,
	}
}

func (r *PaymentRoute) RegisterRoutes() {}
