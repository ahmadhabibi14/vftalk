package services

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"vftalk/models/databases"
	"vftalk/utils"

	"github.com/rs/zerolog"
)

type userImpl struct {
	DB  *sql.DB
	Log *zerolog.Logger
}

func NewUser(db *sql.DB, l *zerolog.Logger) *userImpl {
	return &userImpl{
		DB:  db,
		Log: l,
	}
}

type (
	InUser_FindById struct {
		Id string `validate:"required,min=5,max=36" json:"id"`
	}
	OutUser_FindById struct {
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
)

func (u *userImpl) FindById(id string) (OutUser_FindById, error) {
	ctx := context.Background()
	outUser := OutUser_FindById{}

	msg, err := utils.ValidateStruct(InUser_FindById{Id: id})
	if err != nil {
		return outUser, fmt.Errorf(msg)
	}

	userrepo := databases.NewUser(u.DB, u.Log)
	user, err := userrepo.FindById(ctx, id)
	if err != nil {
		return outUser, err
	}
	outUser = OutUser_FindById{
		UserID:     user.UserID,
		Username:   user.Username,
		FullName:   user.FullName,
		Email:      user.Email,
		Password:   user.Password,
		Avatar:     user.Avatar,
		JoinAt:     user.JoinAt,
		LastActive: user.LastActive,
		Website:    user.Website,
		Location:   user.Location,
	}
	return outUser, nil
}
