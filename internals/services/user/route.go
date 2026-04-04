// Package user contains the user routes.
package user

import (
	"context"
	"net/http"

	db "cinema_booking/internals/postgres/generated"

	"github.com/labstack/echo/v5"
	glide "github.com/valkey-io/valkey-glide/go/v2"
)

type UserRoute struct {
	e      *echo.Group
	db     *db.Queries
	valkey *glide.Client
	store  UserStore
}

func NewUserRoute(e *echo.Group, db *db.Queries, valkey *glide.Client) *UserRoute {
	grouped := e.Group("/user")
	store := NewUserStore(context.Background(), db, valkey)
	return &UserRoute{
		e:      grouped,
		db:     db,
		valkey: valkey,
		store:  store,
	}
}

func (r *UserRoute) RegisterRoutes() {
	r.e.POST("/profile", r.SignUp)
	r.e.POST("/signup", r.SignUp)
	r.e.POST("/login", r.Login)
	r.e.PUT("/update", r.UpdateUser, JWTMiddleware)
	r.e.DELETE("/delete", r.DeleteUser, JWTMiddleware)
}

func (r *UserRoute) GetUser(c *echo.Context) error {
	paramID := c.Param("id")
	tokenUserID := c.Get("user_id").(string)

	if paramID != tokenUserID {
		return c.JSON(403, map[string]string{"error": "forbidden"})
	}

	user, err := r.store.GetUserByID(paramID)
	if err != nil {
		return c.JSON(404, map[string]string{"error": "user not found"})
	}

	return c.JSON(200, user)
}

func (r *UserRoute) SignUp(c *echo.Context) error {
	var req SignUpRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	user, err := r.store.CreateUser(req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not create user"})
	}
	return c.JSON(http.StatusCreated, user)
}

func (r *UserRoute) Login(c *echo.Context) error {
	var req LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{"error": "invalid request"})
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	user, err := r.store.GetUserByID(req.ID)
	if err != nil {
		return c.JSON(401, map[string]string{"error": "invalid credentials"})
	}

	token, err := GenerateJWT(user.ID)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "could not generate token"})
	}

	return c.JSON(200, map[string]any{
		"user":  user,
		"token": token,
	})
}

func (r *UserRoute) UpdateUser(c *echo.Context) error {
	tokenUserID := c.Get("user_id").(string)

	var req UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Prevent updating other users
	if req.ID != tokenUserID {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "forbidden"})
	}

	user, err := r.store.UpdateUser(tokenUserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not update user"})
	}

	return c.JSON(http.StatusOK, user)
}

func (r *UserRoute) DeleteUser(c *echo.Context) error {
	tokenUserID := c.Get("user_id").(string)

	if err := r.store.DeleteUser(tokenUserID); err != nil {
		return c.JSON(500, map[string]string{"error": "could not delete user"})
	}

	return c.JSON(200, map[string]string{"message": "user deleted"})
}
