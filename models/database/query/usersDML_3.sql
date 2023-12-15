-- name: UpdateUserProfile :exec
UPDATE Users
SET full_name = ?, location = ?, website = ?
WHERE user_id = ?