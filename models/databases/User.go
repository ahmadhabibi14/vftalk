package databases

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/rs/zerolog"
)

type userImpl struct {
	DB  *sql.DB
	Log *zerolog.Logger
}

type User struct {
	UserID     string    `db:"user_id" json:"user_id"`
	Username   string    `db:"username" json:"username"`
	FullName   string    `db:"full_name" json:"full_name"`
	Email      string    `db:"email" json:"email"`
	Password   string    `db:"password" json:"password"`
	Avatar     string    `db:"avatar" json:"avatar"`
	JoinAt     time.Time `db:"join_at" json:"join_at"`
	LastActive time.Time `db:"last_active" json:"last_active"`
	Website    string    `db:"website" json:"website"`
	Location   string    `db:"location" json:"location"`
}

func NewUser(db *sql.DB, l *zerolog.Logger) *userImpl {
	return &userImpl{
		DB:  db,
		Log: l,
	}
}

func (u *userImpl) FindById(ctx context.Context, id string) (User, error) {
	query := `SELECT user_id, username, full_name, email, password, avatar, join_at, last_active, website, location FROM Users WHERE user_id = ? LIMIT 1`
	rows, err := u.DB.QueryContext(ctx, query, id)
	defer rows.Close()
	user := User{}
	if err != nil {
		u.Log.Error().Str("Error", err.Error()).Msg("Error: User FindById")
		return user, err
	}
	if rows.Next() {
		rows.Scan(
			&user.UserID,
			&user.Username,
			&user.FullName,
			&user.Email,
			&user.Password,
			&user.Avatar,
			&user.JoinAt,
			&user.LastActive,
			&user.Website,
			&user.Location,
		)
	} else {
		u.Log.Error().Str("Error", "User not found")
		return user, errors.New("User not found")
	}
	return user, nil
}
