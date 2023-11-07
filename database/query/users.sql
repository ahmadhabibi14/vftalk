-- name: CreateNewUser :exec
INSERT INTO Users (
  user_id, username, full_name, email, password
) VALUES (
  ?, ?, ?, ?, ?
);

-- name: GetUserByUsername :one
SELECT username FROM Users
WHERE username = ?;

-- name: GetUserByEmail :one
SELECT email FROM Users
WHERE email = ?;

-- name: ListUsers :many
SELECT username, full_name, email, avatar, join_at FROM Users
ORDER BY join_at DESC;

-- name: UserLogin :one
SELECT user_id, username, password FROM Users
WHERE username = ?;

-- name: GetUserDataByUserId :one
SELECT username, full_name, email, avatar, join_at FROM Users
WHERE user_id = ?;