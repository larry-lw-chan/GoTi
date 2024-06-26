-- name: GetAllProfiles :many
SELECT * FROM profiles;

-- name: GetProfileFromUserId :one
SELECT * FROM profiles WHERE user_id = ?;

-- name: GetProfileFromUsername :one
SELECT * FROM profiles WHERE username = ?;

-- name: GetProfileAvatarFromUserId :one
SELECT avatar FROM profiles WHERE user_id = ?;

-- name: CreateProfile :one
INSERT INTO profiles (username, name, bio, link, avatar, private, user_id, created_at, updated_at) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: UpdateProfile :one
UPDATE profiles 
SET username = ?, name = ?, bio = ?, link = ?, private = ?, updated_at = ? 
WHERE user_id = ? 
RETURNING *;


-- name: UpdateProfileAvatar :one
UPDATE profiles 
SET avatar = ?, updated_at = ? 
WHERE user_id = ? 
RETURNING *;