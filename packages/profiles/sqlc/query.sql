-- name: GetProfileFromUserId :one
SELECT * FROM profiles WHERE user_id = ?;

-- name: CreateProfile :one
INSERT INTO profiles (name, bio, link, avatar, user_id, created_at, updated_at) 
VALUES (?, ?, ?, ?, ?, ?, ?)
RETURNING *;