-- name: CreateThread :one
INSERT INTO threads (content, thread_id, user_id, created_at, updated_at)
VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: GetAllThreads :many
SELECT content, username
FROM threads
JOIN users ON threads.user_id = users.id;