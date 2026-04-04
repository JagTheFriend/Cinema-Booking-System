// Package user contains the user routes.
package user

import (
	db "cinema_booking/internals/postgres/generated"

	"github.com/labstack/echo/v5"
	glide "github.com/valkey-io/valkey-glide/go/v2"
)

type UserRoute struct {
	e      *echo.Group
	db     *db.Queries
	valkey *glide.Client
}

func NewUserRoute(e *echo.Group, db *db.Queries, valkey *glide.Client) *UserRoute {
	return &UserRoute{
		e:      e,
		db:     db,
		valkey: valkey,
	}
}

func (r *UserRoute) RegisterRoutes() {}
