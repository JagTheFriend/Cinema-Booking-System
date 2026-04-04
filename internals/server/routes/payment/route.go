// Package payment contains the payment routes.
package payment

import (
	db "cinema_booking/internals/postgres/generated"

	"github.com/labstack/echo/v5"
)

type PaymentRoute struct {
	e  *echo.Group
	db *db.Queries
}

func NewPaymentRoute(e *echo.Group, db *db.Queries) *PaymentRoute {
	return &PaymentRoute{
		e:  e,
		db: db,
	}
}

func (r *PaymentRoute) RegisterRoutes() {}
