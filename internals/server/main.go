// Package server contains the code to start the server and register the routes.
package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	db "cinema_booking/internals/postgres/generated"
	"cinema_booking/internals/server/routes/booking"
	"cinema_booking/internals/server/routes/payment"
	"cinema_booking/internals/server/routes/user"
	"cinema_booking/internals/valkey"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}

func StartServer() {
	port, ok := os.LookupEnv("SERVER_PORT")
	if !ok {
		panic("PORT not set")
	}

	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
	e.Use(middleware.ContextTimeout(time.Second * 5))

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
				)
			} else {
				logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	}))

	e.Validator = &CustomValidator{validator: validator.New()}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	conn, err := pgx.Connect(ctx, "user=pqgotest dbname=pqgotest sslmode=verify-full")
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	defer func() {
		err := conn.Close(ctx)
		if err != nil {
			slog.Error("Failed to close database connection", "error", err.Error())
		}
	}()

	queries := db.New(conn)
	valkey := valkey.GetValKeyClient()

	groupedRoute := e.Group("/api/v1")

	groupedRoute.GET("/health", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	userRoute := user.NewUserRoute(groupedRoute, queries, valkey)
	userRoute.RegisterRoutes()

	bookingRoute := booking.NewBookingRoute(groupedRoute, queries, valkey)
	bookingRoute.RegisterRoutes()

	paymentRoute := payment.NewPaymentRoute(groupedRoute, queries, valkey)
	paymentRoute.RegisterRoutes()

	e.AcquireContext().Logger().Info("Server started on port: " + port)

	if err := e.Start(":" + port); err != nil {
		e.Logger.Error("Failed to start server", "error", err)
	}
}
