-- name: CreateNewUser :exec
INSERT INTO Users (
  user_id, username, full_name, email, password
) VALUES (
  ?, ?, ?, ?, ?
);

-- name: UpdateUserAvatar :exec
UPDATE Users
SET avatar = ?
WHERE user_id = ?