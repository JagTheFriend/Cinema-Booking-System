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
	grouped := e.Group("/user")
	return &UserRoute{
		e:      grouped,
		db:     db,
		valkey: valkey,
	}
}

func (r *UserRoute) RegisterRoutes() {
	r.e.POST("/profile", r.SignUp)
	r.e.POST("/signup", r.SignUp)
	r.e.POST("/login", r.Login)
	r.e.PUT("/update", r.UpdateUser)
	r.e.DELETE("/delete", r.DeleteUser)
}

func (r *UserRoute) GetUser(ctx *echo.Context) error {
	return nil
}

func (r *UserRoute) SignUp(ctx *echo.Context) error {
	return nil
}

func (r *UserRoute) Login(ctx *echo.Context) error {
	return nil
}

func (r *UserRoute) UpdateUser(ctx *echo.Context) error {
	return nil
}

func (r *UserRoute) DeleteUser(ctx *echo.Context) error {
	return nil
}
