// Package payment contains the payment routes.
package payment

import (
	"context"
	"net/http"

	db "cinema_booking/internals/postgres/generated"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
	glide "github.com/valkey-io/valkey-glide/go/v2"
)

type PaymentRoute struct {
	e      *echo.Group
	db     *db.Queries
	valkey *glide.Client
	store  PaymentStore
}

func NewPaymentRoute(e *echo.Group, db *db.Queries, valkey *glide.Client) *PaymentRoute {
	grouped := e.Group("/payment")
	store := NewPaymentStore(context.Background(), db, valkey)
	return &PaymentRoute{
		e:      grouped,
		db:     db,
		valkey: valkey,
		store:  store,
	}
}

func (r *PaymentRoute) RegisterRoutes() {
	r.e.POST("/create", r.CreatePayment)
	r.e.GET("/:id", r.GetPayment)
	r.e.GET("/user/:user_id", r.ListPaymentsByUser)
	r.e.PUT("/:id", r.UpdatePayment)
	r.e.DELETE("/:id", r.DeletePayment)
}

func (r *PaymentRoute) CreatePayment(c *echo.Context) error {
	userID := c.Get("user_id").(string)

	var req CreatePaymentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	payment, err := r.store.CreatePayment(&db.CreatePaymentParams{
		ID:        uuid.NewString(),
		UserID:    userID,
		BookingID: req.BookingID,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not create payment"})
	}

	return c.JSON(http.StatusCreated, payment)
}

func (r *PaymentRoute) GetPayment(c *echo.Context) error {
	paramID := c.Param("id")
	tokenUserID := c.Get("user_id").(string)

	payment, err := r.store.GetPaymentByID(paramID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "payment not found"})
	}

	if payment.UserID != tokenUserID {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "forbidden"})
	}

	return c.JSON(http.StatusOK, payment)
}

func (r *PaymentRoute) ListPaymentsByUser(c *echo.Context) error {
	paramUserID := c.Param("user_id")
	tokenUserID := c.Get("user_id").(string)

	if paramUserID != tokenUserID {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "forbidden"})
	}

	payments, err := r.store.ListPaymentsByUser(paramUserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not fetch payments"})
	}

	return c.JSON(http.StatusOK, payments)
}

func (r *PaymentRoute) UpdatePayment(c *echo.Context) error {
	paramID := c.Param("id")
	tokenUserID := c.Get("user_id").(string)

	payment, err := r.store.GetPaymentByID(paramID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "payment not found"})
	}

	if payment.UserID != tokenUserID {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "forbidden"})
	}

	updated, err := r.store.UpdatePayment(paramID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not update payment"})
	}

	return c.JSON(http.StatusOK, updated)
}

func (r *PaymentRoute) DeletePayment(c *echo.Context) error {
	paramID := c.Param("id")
	tokenUserID := c.Get("user_id").(string)

	payment, err := r.store.GetPaymentByID(paramID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "payment not found"})
	}

	if payment.UserID != tokenUserID {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "forbidden"})
	}

	if err := r.store.DeletePayment(paramID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not delete payment"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "payment deleted"})
}
