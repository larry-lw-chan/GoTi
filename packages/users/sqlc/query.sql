-- name: GetUserFromID :one
SELECT * FROM users WHERE id = ?;
