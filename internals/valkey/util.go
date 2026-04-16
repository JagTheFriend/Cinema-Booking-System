package valkey

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/valkey-io/valkey-glide/go/v2/options"
)

type Booking struct {
	Status    string `json:"status"`
	SeatID    string `json:"seatId"`
	MovieID   string `json:"movieId"`
	UserID    string `json:"userId"`
	PaymentID string `json:"paymentId"`
}

func FormatClientKey(status, seatID, movieID string) string {
	return fmt.Sprintf("client:%s:%s:%s", status, seatID, movieID)
}

func AddBooking(ctx context.Context, client *Booking) error {
	valkey := GetValKeyClient()

	key := FormatClientKey(client.Status, client.SeatID, client.MovieID)

	// Enforce uniqueness
	exists, err := valkey.Exists(ctx, []string{key})
	if err != nil {
		return err
	}
	if exists > 0 {
		return errors.New("booking already exists for this seat and movie")
	}

	data, err := json.Marshal(client)
	if err != nil {
		return err
	}

	// Set with 10-minute expiry
	_, err = valkey.SetWithOptions(ctx, key, string(data), options.SetOptions{
		Expiry: options.NewExpiryIn(10 * time.Minute),
	})
	return err
}

func GetBooking(ctx context.Context, status, seatID, movieID string) (*Booking, error) {
	valkey := GetValKeyClient()

	key := FormatClientKey(status, seatID, movieID)

	result, err := valkey.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	var booking Booking
	if err := json.Unmarshal([]byte(result.Value()), &booking); err != nil {
		return nil, err
	}

	return &booking, nil
}

func DeleteBooking(ctx context.Context, status, seatID, movieID string) error {
	valkey := GetValKeyClient()

	key := FormatClientKey(status, seatID, movieID)

	_, err := valkey.Del(ctx, []string{key})
	return err
}

func UpdateBookingPayment(ctx context.Context, status, seatID, movieID, newPaymentID string) error {
	valkey := GetValKeyClient()

	key := FormatClientKey(status, seatID, movieID)

	// Get existing booking
	result, err := valkey.Get(ctx, key)
	if err != nil {
		return err
	}

	var booking Booking
	if err = json.Unmarshal([]byte(result.Value()), &booking); err != nil {
		return err
	}

	// Update payment info
	booking.PaymentID = newPaymentID

	updatedData, err := json.Marshal(booking)
	if err != nil {
		return err
	}

	_, err = valkey.Set(ctx, key, string(updatedData))
	return err
}
