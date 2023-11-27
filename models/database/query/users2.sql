-- name: UpdateUserAvatar :exec
UPDATE Users
SET avatar = ?
WHERE user_id = ?