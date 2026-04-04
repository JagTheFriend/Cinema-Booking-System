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
	return &BookingRoute{
		e:      e,
		db:     db,
		valkey: valkey,
	}
}

func (r *BookingRoute) RegisterRoutes() {}
