-- name: UpdateUserPassword :exec
UPDATE Users
SET password = ?
WHERE user_id = ?