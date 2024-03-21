-- name: GetProfileFromUserId :one
SELECT * FROM profiles WHERE user_id = ?;

-- name: CreateProfile :one
INSERT INTO profiles (name, bio, link, avatar, private, user_id, created_at, updated_at) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: UpdateProfile :one
UPDATE profiles 
SET name = ?, bio = ?, link = ?, avatar = ?, private = ?, updated_at = ? 
WHERE user_id = ? 
RETURNING *;