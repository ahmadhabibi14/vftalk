-- name: UpdateUserData :exec
UPDATE Users
SET username = ?, full_name = ?, email = ?
WHERE user_id = ?