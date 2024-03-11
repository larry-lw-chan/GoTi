-- name: GetUserFromEmail :one
SELECT * FROM users WHERE email = ?;

-- name: CreateUser :one
INSERT INTO users (username, email, password, created_at, updated_at) 
VALUES (?, ?, ?, ?, ?)
RETURNING *;