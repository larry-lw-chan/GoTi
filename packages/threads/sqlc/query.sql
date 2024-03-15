-- name: CreateThread :one
INSERT INTO threads (content, thread_id, user_id, created_at, updated_at)
VALUES (?, ?, ?, ?, ?)
RETURNING *;