-- name: GetUserFromID :one
SELECT * FROM users WHERE id = ?;

-- name: CreateUser :one
INSERT INTO users (username, email, password, created_at, updated_at) 
VALUES (?, ?, ?, ?, ?)
RETURNING *;