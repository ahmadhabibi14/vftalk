package services

import (
	"context"
	"database/sql"
	"errors"
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
	if userrepo.FindUsername(ctx, in.Username) != "" {
		return "", fmt.Errorf("Username already exist")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	user := databases.CreateUserIn{
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

type (
	InUser_OAuthCreate struct {
		UserID   string `json:"id" form:"id" validate:"required,min=21,max=36"`
		Username string `json:"username" form:"username" validate:"required,omitempty,min=4"`
		FullName string `json:"full_name" form:"full_name" validate:"required,omitempty,min=4"`
		Email    string `json:"email" form:"email" validate:"required,email"`
		Avatar   string `json:"avatar" form:"avatar" validate:"required"`
	}
)

func (u *userImpl) OAuthCreateUser(ctx context.Context, in InUser_OAuthCreate) (token string, err error) {
	msg, err := utils.ValidateStruct(in)
	if err != nil {
		return "", fmt.Errorf(msg)
	}

	t, err := configs.GenerateJWT(in.Username, in.UserID, time.Now().AddDate(0, 2, 0))
	if err != nil {
		return "", fmt.Errorf("Error generate session token")
	}

	userrepo := databases.NewUser(u.DB, u.Log)
	if userrepo.FindUsername(ctx, in.Username) != `` {
		return t, nil
	}

	if userrepo.FindId(ctx, in.UserID) != `` {
		return t, nil
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(in.Username), bcrypt.DefaultCost)
	user := databases.OAuthCreateUserIn{
		UserID:   in.UserID,
		Username: in.Username,
		FullName: in.FullName,
		Email:    in.Email,
		Password: string(hashedPassword),
		Avatar:   in.Avatar,
	}
	err = userrepo.OAuthCreateUser(ctx, user)
	if err != nil {
		return "", fmt.Errorf("Something went wrong")
	}

	return t, nil
}

type (
	InUser_AuthLogin struct {
		Username string `json:"username" form:"username" validate:"required,omitempty,min=4"`
		Password string `json:"password" form:"password" validate:"required,min=8"`
	}
)

func (u *userImpl) AuthLogin(ctx context.Context, in InUser_AuthLogin) (token, username string, err error) {
	msg, err := utils.ValidateStruct(in)
	if err != nil {
		return "", "", fmt.Errorf(msg)
	}

	userrepo := databases.NewUser(u.DB, u.Log)
	user, err := userrepo.FindByUsername(ctx, in.Username)
	if err != nil {
		return "", "", fmt.Errorf("Username not found")
	}

	passwordMatch := utils.VerifyPassword(in.Password, user.Password)
	if passwordMatch != nil {
		return "", "", fmt.Errorf("Password does not match the user's password")
	}

	t, err := configs.GenerateJWT(user.Username, user.UserID, time.Now().AddDate(0, 2, 0))
	if err != nil {
		return "", "", fmt.Errorf("Error generate session token")
	}

	return t, user.Username, nil
}

type (
	InUser_UpdateProfile struct {
		UserID   string `json:"user_id" validate:"required"`
		FullName string `json:"full_name" validate:"required"`
		Location string `json:"location" validate:"required"`
		Website  string `json:"website" validate:"required"`
	}
)

func (u *userImpl) UpdateProfile(ctx context.Context, in InUser_UpdateProfile) error {
	msg, err := utils.ValidateStruct(in)
	if err != nil {
		return fmt.Errorf(msg)
	}
	userrepo := databases.NewUser(u.DB, u.Log)
	if userrepo.FindId(ctx, in.UserID) == `` {
		return errors.New("User not found")
	}

	user := databases.UpdateUserProfileIn{
		UserID:   in.UserID,
		FullName: in.FullName,
		Location: in.Location,
		Website:  in.Website,
	}
	updateProfile := userrepo.UpdateUserProfile(ctx, user)
	if updateProfile != nil {
		u.Log.Error().Msg(updateProfile.Error())
		return updateProfile
	}

	return nil
}

type (
	InUser_UpdateAvatar struct {
		UserID string `form:"user_id" validate:"required"`
		Avatar string `form:"avatar" validate:"required"`
	}
)

func (u *userImpl) UpdateAvatar(ctx context.Context, in InUser_UpdateAvatar) error {
	msg, err := utils.ValidateStruct(in)
	if err != nil {
		return fmt.Errorf(msg)
	}
	userrepo := databases.NewUser(u.DB, u.Log)
	if userrepo.FindId(ctx, in.UserID) == `` {
		return errors.New("User not found")
	}

	user := databases.UpdateUserAvatarIn{
		UserID: in.UserID,
		Avatar: in.Avatar,
	}
	updateAvatar := userrepo.UpdateUserAvatar(ctx, user)
	if updateAvatar != nil {
		u.Log.Error().Msg(updateAvatar.Error())
		return updateAvatar
	}

	return nil
}

type (
	OutUserLists struct {
		Username string `db:"username" json:"username"`
		FullName string `db:"full_name" json:"full_name"`
		Avatar   string `db:"avatar" json:"avatar"`
	}
)

func (u *userImpl) UserLists(ctx context.Context) ([]OutUserLists, error) {
	userrepo := databases.NewUser(u.DB, u.Log)
	users, err := userrepo.FindAll(ctx)

	outUsers := []OutUserLists{}
	if err != nil {
		return outUsers, err
	}

	for _, v := range users {
		outUser := OutUserLists{}
		outUser.Username = v.Username
		outUser.FullName = v.FullName
		outUser.Avatar = v.Avatar

		outUsers = append(outUsers, outUser)
	}

	return outUsers, nil
}

func (u *userImpl) Debug(ctx context.Context, id string) bool {
	userrepo := databases.NewUser(u.DB, u.Log)
	if userrepo.FindId(ctx, id) == `` {
		return false
	}

	return true
}
