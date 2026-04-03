-- =========================
-- USERS
-- =========================

-- name: CreateUser :one
INSERT INTO users (id) 
VALUES ($1)
RETURNING id, created_at, updated_at;

-- name: GetUserByID :one
SELECT id, created_at, updated_at
FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET updated_at = NOW()
WHERE id = $1
RETURNING id, created_at, updated_at;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT id, created_at, updated_at
FROM users
ORDER BY created_at DESC;
