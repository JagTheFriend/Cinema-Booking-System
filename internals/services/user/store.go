package user

import (
	"context"

	db "cinema_booking/internals/postgres/generated"

	glide "github.com/valkey-io/valkey-glide/go/v2"
)

type UserStore interface {
	CreateUser(id string) (*db.User, error)
	GetUserByID(id string) (*db.User, error)
	ListUsers() ([]*db.User, error)
	UpdateUser(id string) (*db.User, error)
	DeleteUser(id string) error
}

type UserStoreImpl struct {
	db     *db.Queries
	valkey *glide.Client
	ctx    context.Context
}

func NewUserStore(ctx context.Context, db *db.Queries, valkey *glide.Client) UserStore {
	return &UserStoreImpl{
		db:     db,
		valkey: valkey,
		ctx:    ctx,
	}
}

// CreateUser inserts a new user into the database
func (s *UserStoreImpl) CreateUser(id string) (*db.User, error) {
	return s.db.CreateUser(s.ctx, id)
}

// GetUserByID fetches a user by ID
func (s *UserStoreImpl) GetUserByID(id string) (*db.User, error) {
	return s.db.GetUserByID(s.ctx, id)
}

// ListUsers returns all users
func (s *UserStoreImpl) ListUsers() ([]*db.User, error) {
	return s.db.ListUsers(s.ctx)
}

// UpdateUser updates the updated_at timestamp of a user
func (s *UserStoreImpl) UpdateUser(id string) (*db.User, error) {
	return s.db.UpdateUser(s.ctx, id)
}

// DeleteUser removes a user by ID
func (s *UserStoreImpl) DeleteUser(id string) error {
	return s.db.DeleteUser(s.ctx, id)
}
