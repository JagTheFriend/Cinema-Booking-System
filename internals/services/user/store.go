package user

import (
	"context"

	db "cinema_booking/internals/postgres/generated"
)

type UserStore interface {
	CreateUser(ctx context.Context, id string) (*db.User, error)
	GetUserByID(ctx context.Context, id string) (*db.User, error)
	ListUsers(ctx context.Context) ([]*db.User, error)
	UpdateUser(ctx context.Context, id string) (*db.User, error)
	DeleteUser(ctx context.Context, id string) error
}
