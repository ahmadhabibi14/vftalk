package services

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"vftalk/configs"
	"vftalk/models/databases"
	"vftalk/utils"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
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
		UserID string `json:"id" form:"id" validate:"required,min=21,max=36"`
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

func (u *userImpl) FindById(ctx context.Context, in InUser_FindById) (OutUser_FindById, error) {
	outUser := OutUser_FindById{}

	msg, err := utils.ValidateStruct(in)
	if err != nil {
		return outUser, fmt.Errorf(msg)
	}

	userrepo := databases.NewUser(u.DB, u.Log)
	user, err := userrepo.FindById(ctx, in.UserID)
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

type (
	InUser_Create struct {
		UserID   string `json:"id" form:"id" validate:"required,min=21,max=36"`
		Username string `json:"username" form:"username" validate:"required,omitempty,min=4"`
		FullName string `json:"full_name" form:"full_name" validate:"required,omitempty,min=4"`
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required,min=8"`
	}
)

func (u *userImpl) CreateUser(ctx context.Context, in InUser_Create) (token string, err error) {
	uid := fmt.Sprintf("%v", uuid.New())
	in.UserID = uid
	msg, err := utils.ValidateStruct(in)
	if err != nil {
		return "", fmt.Errorf(msg)
	}

	userrepo := databases.NewUser(u.DB, u.Log)
	_, err = userrepo.FindByUsername(ctx, in.Username)
	if err == nil {
		return "", fmt.Errorf("Username already exist")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	user := databases.User{
		UserID:   in.UserID,
		Username: in.Username,
		FullName: in.FullName,
		Email:    in.Email,
		Password: string(hashedPassword),
	}
	err = userrepo.CreateUser(ctx, user)
	if err != nil {
		return "", fmt.Errorf("Something went wrong")
	}

	t, err := configs.GenerateJWT(in.Username, uid, time.Now().AddDate(0, 2, 0))
	if err != nil {
		return "", fmt.Errorf("Error generate session token")
	}
	return t, nil
}
