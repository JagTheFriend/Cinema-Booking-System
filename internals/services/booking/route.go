// Package booking contains the booking routes.
package booking

import (
	"context"
	"net/http"

	db "cinema_booking/internals/postgres/generated"
	"cinema_booking/internals/services/user"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

type BookingRoute struct {
	e     *echo.Group
	db    *db.Queries
	store IBookingStore
}

func NewBookingRoute(e *echo.Group, db *db.Queries, valkey any) *BookingRoute { // valkey kept for consistency
	grouped := e.Group("/booking")
	store := NewBookingStore(context.Background(), db, nil)
	return &BookingRoute{
		e:     grouped,
		db:    db,
		store: store,
	}
}

func (r *BookingRoute) RegisterRoutes() {
	r.e.POST("/create", r.CreateBooking, user.JWTMiddleware)
	r.e.GET("/:id", r.GetBooking, user.JWTMiddleware)
	r.e.GET("/user/:user_id", r.ListBookingsByUser, user.JWTMiddleware)
	r.e.PUT("/verify/:id", r.VerifyBooking, user.JWTMiddleware)
	r.e.PUT("/payment", r.UpdateBookingPayment, user.JWTMiddleware)
	r.e.DELETE("/:id", r.DeleteBooking, user.JWTMiddleware)
}

func (r *BookingRoute) CreateBooking(c *echo.Context) error {
	userID := c.Get("user_id").(string)

	var req CreateBookingRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	bookingID := uuid.NewString()
	booking, err := r.store.CreateBooking(&db.CreateBookingParams{
		ID:         bookingID,
		SeatID:     req.SeatID,
		MovieID:    req.MovieID,
		UserID:     userID,
		PaymentID:  req.PaymentID,
		IsVerified: false,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not create booking"})
	}

	return c.JSON(http.StatusCreated, booking)
}

func (r *BookingRoute) GetBooking(c *echo.Context) error {
	paramID := c.Param("id")
	tokenUserID := c.Get("user_id").(string)

	booking, err := r.store.GetBookingByID(paramID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "booking not found"})
	}

	if booking.UserID != tokenUserID {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "forbidden"})
	}

	return c.JSON(http.StatusOK, booking)
}

func (r *BookingRoute) ListBookingsByUser(c *echo.Context) error {
	paramUserID := c.Param("user_id")
	tokenUserID := c.Get("user_id").(string)

	if paramUserID != tokenUserID {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "forbidden"})
	}

	bookings, err := r.store.ListBookingsByUser(paramUserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not fetch bookings"})
	}

	return c.JSON(http.StatusOK, bookings)
}

func (r *BookingRoute) VerifyBooking(c *echo.Context) error {
	paramID := c.Param("id")
	tokenUserID := c.Get("user_id").(string)

	booking, err := r.store.GetBookingByID(paramID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "booking not found"})
	}

	if booking.UserID != tokenUserID {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "forbidden"})
	}

	verified, err := r.store.VerifyBooking(paramID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not verify booking"})
	}

	return c.JSON(http.StatusOK, verified)
}

func (r *BookingRoute) UpdateBookingPayment(c *echo.Context) error {
	var req UpdateBookingPaymentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	booking, err := r.store.GetBookingByID(req.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "booking not found"})
	}

	tokenUserID := c.Get("user_id").(string)
	if booking.UserID != tokenUserID {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "forbidden"})
	}

	updated, err := r.store.UpdateBookingPayment(&db.UpdateBookingPaymentParams{
		ID:        req.ID,
		PaymentID: req.PaymentID,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not update booking payment"})
	}

	return c.JSON(http.StatusOK, updated)
}

func (r *BookingRoute) DeleteBooking(c *echo.Context) error {
	paramID := c.Param("id")
	tokenUserID := c.Get("user_id").(string)

	booking, err := r.store.GetBookingByID(paramID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "booking not found"})
	}

	if booking.UserID != tokenUserID {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "forbidden"})
	}

	if err := r.store.DeleteBooking(paramID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not delete booking"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "booking deleted"})
}
