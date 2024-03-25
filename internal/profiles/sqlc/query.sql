-- name: GetProfileFromUserId :one
SELECT * FROM profiles WHERE user_id = ?;

-- name: GetProfileAvatarFromUserId :one
SELECT avatar FROM profiles WHERE user_id = ?;

-- name: CreateProfile :one
INSERT INTO profiles (username, name, bio, link, avatar, private, user_id, created_at, updated_at) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: UpdateProfile :one
UPDATE profiles 
SET username = ?, name = ?, bio = ?, link = ?, avatar = ?, private = ?, updated_at = ? 
WHERE user_id = ? 
RETURNING *;