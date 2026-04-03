-- =========================
-- PAYMENTS
-- =========================

-- name: CreatePayment :one
INSERT INTO payments (id, user_id, booking_id)
VALUES ($1, $2, $3)
RETURNING id, user_id, booking_id, created_at, updated_at;

-- name: GetPaymentByID :one
SELECT id, user_id, booking_id, created_at, updated_at
FROM payments
WHERE id = $1;

-- name: UpdatePayment :one
UPDATE payments
SET updated_at = NOW()
WHERE id = $1
RETURNING id, user_id, booking_id, created_at, updated_at;

-- name: DeletePayment :exec
DELETE FROM payments
WHERE id = $1;

-- name: ListPaymentsByUser :many
SELECT id, user_id, booking_id, created_at, updated_at
FROM payments
WHERE user_id = $1
ORDER BY created_at DESC;
