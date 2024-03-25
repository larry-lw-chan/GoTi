-- name: GetUserFromEmail :one
SELECT * FROM users WHERE email = ?;

-- name: GetUserFromUsername :one
SELECT * FROM users WHERE username = ?;

-- name: CreateUser :one
INSERT INTO users (uuid, username, email, password, created_at, updated_at) 
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;