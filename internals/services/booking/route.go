// Package booking contains the booking routes.
package booking

import (
	db "cinema_booking/internals/postgres/generated"

	"github.com/labstack/echo/v5"
	glide "github.com/valkey-io/valkey-glide/go/v2"
)

type BookingRoute struct {
	e      *echo.Group
	db     *db.Queries
	valkey *glide.Client
}

func NewBookingRoute(e *echo.Group, db *db.Queries, valkey *glide.Client) *BookingRoute {
	grouped := e.Group("/booking")
	return &BookingRoute{
		e:      grouped,
		db:     db,
		valkey: valkey,
	}
}

func (r *BookingRoute) RegisterRoutes() {
	r.e.POST("/create", r.CreateBooking)
	r.e.GET("/:id", r.GetBooking)
	r.e.GET("/user/:user_id", r.ListBookingsByUser)
	r.e.PUT("/verify/:id", r.VerifyBooking)
	r.e.PUT("/payment", r.UpdateBookingPayment)
	r.e.DELETE("/:id", r.DeleteBooking)
}

func (r *BookingRoute) CreateBooking(ctx *echo.Context) error {
	return nil
}

func (r *BookingRoute) GetBooking(ctx *echo.Context) error {
	return nil
}

func (r *BookingRoute) ListBookingsByUser(ctx *echo.Context) error {
	return nil
}

func (r *BookingRoute) VerifyBooking(ctx *echo.Context) error {
	return nil
}

func (r *BookingRoute) UpdateBookingPayment(ctx *echo.Context) error {
	return nil
}

func (r *BookingRoute) DeleteBooking(ctx *echo.Context) error {
	return nil
}
