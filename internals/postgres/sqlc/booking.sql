-- =========================
-- BOOKINGS
-- =========================

-- name: CreateBooking :one
INSERT INTO bookings (id, seat_id, movie_id, user_id, payment_id, is_verified)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, seat_id, movie_id, user_id, payment_id, is_verified, created_at, updated_at;

-- name: GetBookingByID :one
SELECT id, seat_id, movie_id, user_id, payment_id, is_verified, created_at, updated_at
FROM bookings
WHERE id = $1;

-- name: VerifyBooking :one
UPDATE bookings
SET is_verified = TRUE, updated_at = NOW()
WHERE id = $1
RETURNING id, seat_id, movie_id, user_id, payment_id, is_verified, created_at, updated_at;

-- name: UpdateBookingPayment :one
UPDATE bookings
SET payment_id = $2, updated_at = NOW()
WHERE id = $1
RETURNING id, seat_id, movie_id, user_id, payment_id, is_verified, created_at, updated_at;

-- name: DeleteBooking :exec
DELETE FROM bookings
WHERE id = $1;

-- name: ListBookingsByUser :many
SELECT id, seat_id, movie_id, user_id, payment_id, is_verified, created_at, updated_at
FROM bookings
WHERE user_id = $1
ORDER BY created_at DESC;

