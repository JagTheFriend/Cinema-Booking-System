// Package booking contains the booking routes.
package booking

import (
	db "cinema_booking/internals/postgres/generated"

	"github.com/labstack/echo/v5"
)

type BookingRoute struct {
	e  *echo.Group
	db *db.Queries
}

func NewBookingRoute(e *echo.Group, db *db.Queries) *BookingRoute {
	return &BookingRoute{
		e:  e,
		db: db,
	}
}

func (r *BookingRoute) RegisterRoutes() {}
