
-- =========================
-- JOIN QUERIES
-- =========================

-- name: GetBookingWithUserAndPayment :one
SELECT b.id AS booking_id, b.seat_id, b.movie_id, b.is_verified, b.created_at AS booking_created,
       u.id AS user_id, u.created_at AS user_created,
       p.id AS payment_id, p.created_at AS payment_created
FROM bookings b
JOIN users u ON b.user_id = u.id
LEFT JOIN payments p ON b.payment_id = p.id
WHERE b.id = $1;

-- name: ListBookingsWithPaymentStatus :many
SELECT b.id AS booking_id, b.seat_id, b.movie_id, b.is_verified,
       CASE WHEN p.id IS NOT NULL THEN TRUE ELSE FALSE END AS is_paid
FROM bookings b
LEFT JOIN payments p ON b.payment_id = p.id
ORDER BY b.created_at DESC;
