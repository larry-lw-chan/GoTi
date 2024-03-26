-- name: CreateThread :one
INSERT INTO threads (content, thread_id, user_id, created_at, updated_at)
VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: GetAllThreads :many
SELECT content, username, avatar
FROM threads
JOIN profiles ON profiles.user_id = threads.user_id;


-- name: GetUserThreads :many
SELECT content, username, avatar
FROM threads
JOIN profiles ON profiles.user_id = threads.user_id
WHERE threads.user_id = ?;