// Package user contains the user routes.
package user

import (
	db "cinema_booking/internals/postgres/generated"

	"github.com/labstack/echo/v5"
)

type UserRoute struct {
	e  *echo.Group
	db *db.Queries
}

func NewUserRoute(e *echo.Group, db *db.Queries) *UserRoute {
	return &UserRoute{
		e:  e,
		db: db,
	}
}

func (r *UserRoute) RegisterRoutes() {}
