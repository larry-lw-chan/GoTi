-- name: CreateFollower :one
INSERT INTO relationships (follower_id, followee_id, created_at, updated_at) 
VALUES (?, ?, ?, ?)
RETURNING *;