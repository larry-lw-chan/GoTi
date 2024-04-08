-- name: CreateThread :one
INSERT INTO threads (content, thread_id, user_id, created_at, updated_at)
VALUES (?, ?, ?, ?, ?)
RETURNING *;


-- name: GetAllThreads :many
SELECT threads.id, content, username, avatar, (
    SELECT COUNT(*) 
    FROM likes 
    WHERE likes.thread_id=threads.id
) AS likes
FROM threads
JOIN profiles ON profiles.user_id = threads.user_id
ORDER BY threads.created_at desc;



-- name: GetUserThreads :many
SELECT threads.id, content, username, avatar, (
    SELECT COUNT(*) 
    FROM likes 
    WHERE likes.thread_id=threads.id
) AS likes
FROM threads
JOIN profiles ON profiles.user_id = threads.user_id
WHERE threads.user_id = ?
ORDER BY threads.created_at desc;



-- name: GetThreadByID :one
SELECT threads.id, content, username, avatar
FROM threads
JOIN profiles ON profiles.user_id = threads.user_id
WHERE threads.id = ?;


-- name: GetLikes :many
SELECT *
FROM likes
WHERE thread_id = ? AND user_id = ?;


-- name: InsertLike :one
INSERT INTO likes (thread_id, user_id, created_at, updated_at)
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: DeleteLike :one
DELETE FROM likes
WHERE thread_id = ? AND user_id = ?
RETURNING *;