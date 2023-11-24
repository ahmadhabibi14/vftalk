// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: users.sql

package sqlc

import (
	"context"
	"time"
)

const createNewUser = `-- name: CreateNewUser :exec
INSERT INTO Users (
  user_id, username, full_name, email, password
) VALUES (
  ?, ?, ?, ?, ?
)
`

type CreateNewUserParams struct {
	UserID   string `db:"user_id" json:"user_id"`
	Username string `db:"username" json:"username"`
	FullName string `db:"full_name" json:"full_name"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

func (q *Queries) CreateNewUser(ctx context.Context, arg CreateNewUserParams) error {
	_, err := q.db.ExecContext(ctx, createNewUser,
		arg.UserID,
		arg.Username,
		arg.FullName,
		arg.Email,
		arg.Password,
	)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT email FROM Users
WHERE email = ?
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (string, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	err := row.Scan(&email)
	return email, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT username FROM Users
WHERE username = ?
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (string, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	err := row.Scan(&username)
	return username, err
}

const getUserDataByUserId = `-- name: GetUserDataByUserId :one
SELECT user_id, username, full_name, email, avatar, join_at FROM Users
WHERE user_id = ?
`

type GetUserDataByUserIdRow struct {
	UserID   string    `db:"user_id" json:"user_id"`
	Username string    `db:"username" json:"username"`
	FullName string    `db:"full_name" json:"full_name"`
	Email    string    `db:"email" json:"email"`
	Avatar   string    `db:"avatar" json:"avatar"`
	JoinAt   time.Time `db:"join_at" json:"join_at"`
}

func (q *Queries) GetUserDataByUserId(ctx context.Context, userID string) (GetUserDataByUserIdRow, error) {
	row := q.db.QueryRowContext(ctx, getUserDataByUserId, userID)
	var i GetUserDataByUserIdRow
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.FullName,
		&i.Email,
		&i.Avatar,
		&i.JoinAt,
	)
	return i, err
}

const getUserDataByUsername = `-- name: GetUserDataByUsername :one
SELECT username, full_name, email, avatar, join_at FROM Users
WHERE username = ?
`

type GetUserDataByUsernameRow struct {
	Username string    `db:"username" json:"username"`
	FullName string    `db:"full_name" json:"full_name"`
	Email    string    `db:"email" json:"email"`
	Avatar   string    `db:"avatar" json:"avatar"`
	JoinAt   time.Time `db:"join_at" json:"join_at"`
}

func (q *Queries) GetUserDataByUsername(ctx context.Context, username string) (GetUserDataByUsernameRow, error) {
	row := q.db.QueryRowContext(ctx, getUserDataByUsername, username)
	var i GetUserDataByUsernameRow
	err := row.Scan(
		&i.Username,
		&i.FullName,
		&i.Email,
		&i.Avatar,
		&i.JoinAt,
	)
	return i, err
}

const listUserActive = `-- name: ListUserActive :many
SELECT username, full_name, email, avatar, join_at, last_active FROM Users
WHERE last_active > CURRENT_TIMESTAMP - INTERVAL 10 MINUTE
ORDER BY last_active DESC
`

type ListUserActiveRow struct {
	Username   string    `db:"username" json:"username"`
	FullName   string    `db:"full_name" json:"full_name"`
	Email      string    `db:"email" json:"email"`
	Avatar     string    `db:"avatar" json:"avatar"`
	JoinAt     time.Time `db:"join_at" json:"join_at"`
	LastActive time.Time `db:"last_active" json:"last_active"`
}

func (q *Queries) ListUserActive(ctx context.Context) ([]ListUserActiveRow, error) {
	rows, err := q.db.QueryContext(ctx, listUserActive)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListUserActiveRow{}
	for rows.Next() {
		var i ListUserActiveRow
		if err := rows.Scan(
			&i.Username,
			&i.FullName,
			&i.Email,
			&i.Avatar,
			&i.JoinAt,
			&i.LastActive,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsers = `-- name: ListUsers :many
SELECT username, full_name, email, avatar, join_at, last_active FROM Users
ORDER BY join_at DESC
`

type ListUsersRow struct {
	Username   string    `db:"username" json:"username"`
	FullName   string    `db:"full_name" json:"full_name"`
	Email      string    `db:"email" json:"email"`
	Avatar     string    `db:"avatar" json:"avatar"`
	JoinAt     time.Time `db:"join_at" json:"join_at"`
	LastActive time.Time `db:"last_active" json:"last_active"`
}

func (q *Queries) ListUsers(ctx context.Context) ([]ListUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListUsersRow{}
	for rows.Next() {
		var i ListUsersRow
		if err := rows.Scan(
			&i.Username,
			&i.FullName,
			&i.Email,
			&i.Avatar,
			&i.JoinAt,
			&i.LastActive,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserLastActive = `-- name: UpdateUserLastActive :exec
UPDATE Users
SET last_active = CURRENT_TIMESTAMP + INTERVAL 10 MINUTE
WHERE user_id = ?
`

func (q *Queries) UpdateUserLastActive(ctx context.Context, userID string) error {
	_, err := q.db.ExecContext(ctx, updateUserLastActive, userID)
	return err
}

const userLogin = `-- name: UserLogin :one
SELECT user_id, username, password FROM Users
WHERE username = ?
`

type UserLoginRow struct {
	UserID   string `db:"user_id" json:"user_id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}

func (q *Queries) UserLogin(ctx context.Context, username string) (UserLoginRow, error) {
	row := q.db.QueryRowContext(ctx, userLogin, username)
	var i UserLoginRow
	err := row.Scan(&i.UserID, &i.Username, &i.Password)
	return i, err
}
