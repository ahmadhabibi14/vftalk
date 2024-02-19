package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
	"vftalk/configs"
	"vftalk/models/repository"
	"vftalk/utils"

	"github.com/eefret/gravatar"
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
		UserID     string    `json:"user_id"`
		Username   string    `json:"username"`
		FullName   string    `json:"full_name"`
		Email      string    `json:"email"`
		Password   string    `json:"password"`
		Avatar     string    `json:"avatar"`
		JoinAt     time.Time `json:"join_at"`
		LastActive time.Time `json:"last_active"`
		Website    string    `json:"website"`
		Location   string    `json:"location"`
	}
)

func (u *userImpl) FindById(ctx context.Context, in InUser_FindById) (OutUser_FindById, error) {
	outUser := OutUser_FindById{}

	userrepo := repository.NewUser(u.DB, u.Log)
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
		UserID   string
		Username string `json:"username" form:"username" validate:"required,omitempty,min=5"`
		FullName string `json:"full_name" form:"full_name" validate:"required,omitempty,min=5"`
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required,min=8"`
	}
)

func (u *userImpl) CreateUser(ctx context.Context, in InUser_Create) (token string, err error) {
	uid := fmt.Sprintf("%v", uuid.New())
	in.UserID = uid

	userrepo := repository.NewUser(u.DB, u.Log)
	if userrepo.FindUsername(ctx, in.Username) != "" {
		return "", fmt.Errorf("username already exist")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	g, _ := gravatar.New()
	avatar := g.URLParse(in.Email)
	user := repository.CreateUserIn{
		UserID:   in.UserID,
		Username: in.Username,
		FullName: in.FullName,
		Email:    in.Email,
		Password: string(hashedPassword),
		Avatar:   avatar,
	}
	err = userrepo.CreateUser(ctx, user)
	if err != nil {
		return "", fmt.Errorf("something went wrong")
	}

	t, err := configs.GenerateJWT(in.Username, uid, time.Now().AddDate(0, 2, 0))
	if err != nil {
		return "", fmt.Errorf("error generate session token")
	}
	return t, nil
}

type (
	InUser_OAuthGoogle struct {
		UserID   string
		Username string
		FullName string
		Email    string
		Avatar   string
		GoogleID string
	}
)

func (u *userImpl) OAuthGoogle(ctx context.Context, in InUser_OAuthGoogle) (token string, err error) {
	userrepo := repository.NewUser(u.DB, u.Log)
	user, err := userrepo.FindByGoogleID(ctx, in.GoogleID)
	if err == nil {
		t, _ := configs.GenerateJWT(user.Username, user.UserID, time.Now().AddDate(0, 2, 0))
		return t, nil
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(in.GoogleID), bcrypt.DefaultCost)
	userIn := repository.OAuthGoogleIn{
		UserID:   in.UserID,
		Username: in.Username,
		FullName: in.FullName,
		Email:    in.Email,
		Password: string(hashedPassword),
		Avatar:   in.Avatar,
		GoogleID: in.GoogleID,
	}
	err = userrepo.OAuthGoogle(ctx, userIn)
	if err != nil {
		return "", fmt.Errorf("something went wrong")
	}

	t, _ := configs.GenerateJWT(in.Username, in.UserID, time.Now().AddDate(0, 2, 0))
	return t, nil
}

type (
	InUser_AuthLogin struct {
		Username string `json:"username" form:"username" validate:"required,omitempty,min=5"`
		Password string `json:"password" form:"password" validate:"required,min=8"`
	}
)

func (u *userImpl) AuthLogin(ctx context.Context, in InUser_AuthLogin) (token, username string, err error) {
	userrepo := repository.NewUser(u.DB, u.Log)
	user, err := userrepo.FindByUsername(ctx, in.Username)
	if err != nil {
		return "", "", fmt.Errorf("username not found")
	}

	passwordMatch := utils.VerifyPassword(in.Password, user.Password)
	if passwordMatch != nil {
		return "", "", fmt.Errorf("password does not match the user's password")
	}

	t, err := configs.GenerateJWT(user.Username, user.UserID, time.Now().AddDate(0, 2, 0))
	if err != nil {
		return "", "", fmt.Errorf("error generate session token")
	}

	return t, user.Username, nil
}

type (
	InUser_UpdateProfile struct {
		UserID   string
		Username string `json:"username" validate:"required,min=5"`
		FullName string `json:"full_name" validate:"required,min=5"`
		Location string `json:"location" validate:"required"`
		Website  string `json:"website" validate:"required,http_url"`
	}
)

func (u *userImpl) UpdateProfile(ctx context.Context, in InUser_UpdateProfile) error {
	userrepo := repository.NewUser(u.DB, u.Log)
	if userrepo.FindId(ctx, in.UserID) == `` {
		return errors.New("user not found")
	}

	user := repository.UpdateUserProfileIn{
		UserID:   in.UserID,
		Username: in.Username,
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
		UserID string `form:"user_id"`
		Avatar string `form:"avatar" validate:"required"`
	}
)

func (u *userImpl) UpdateAvatar(ctx context.Context, in InUser_UpdateAvatar) error {
	userrepo := repository.NewUser(u.DB, u.Log)
	if userrepo.FindId(ctx, in.UserID) == `` {
		return errors.New("user not found")
	}

	user := repository.UpdateUserAvatarIn{
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
		Username string `json:"username"`
		FullName string `json:"full_name"`
		Avatar   string `json:"avatar"`
	}
)

func (u *userImpl) UserLists(ctx context.Context) ([]OutUserLists, error) {
	userrepo := repository.NewUser(u.DB, u.Log)
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
	userrepo := repository.NewUser(u.DB, u.Log)
	return userrepo.FindId(ctx, id) != ``
}
