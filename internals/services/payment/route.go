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
	grouped := e.Group("/payment")
	return &PaymentRoute{
		e:      grouped,
		db:     db,
		valkey: valkey,
	}
}

func (r *PaymentRoute) RegisterRoutes() {
	r.e.POST("/create", r.CreatePayment)
	r.e.GET("/:id", r.GetPayment)
	r.e.GET("/user/:user_id", r.ListPaymentsByUser)
	r.e.PUT("/:id", r.UpdatePayment)
	r.e.DELETE("/:id", r.DeletePayment)
}

func (r *PaymentRoute) CreatePayment(ctx *echo.Context) error {
	return nil
}

func (r *PaymentRoute) GetPayment(ctx *echo.Context) error {
	return nil
}

func (r *PaymentRoute) ListPaymentsByUser(ctx *echo.Context) error {
	return nil
}

func (r *PaymentRoute) UpdatePayment(ctx *echo.Context) error {
	return nil
}

func (r *PaymentRoute) DeletePayment(ctx *echo.Context) error {
	return nil
}
