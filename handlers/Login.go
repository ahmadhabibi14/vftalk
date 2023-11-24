package handlers

import (
	"context"
	"database/sql"
	"time"
	"vftalk/conf"
	"vftalk/models/database/sqlc"
	"vftalk/utils"

	json "github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type (
	loginIn struct {
		Username string `json:"username" validate:"required,omitempty,min=4"`
		Password string `json:"password" validate:"required,min=8,containsany=!@#?*%&>_<}-{+"`
	}
	loginOut struct {
		Ok       bool   `json:"ok"`
		Token    string `json:"token"`
		Username string `json:"username"`
		UserId   string `json:"user_id"`
		Message  string `json:"message"`
	}
	loginError struct {
		Ok       bool   `json:"ok"`
		ErrorMsg string `json:"error"`
	}
)

const (
	OutLoginMsg = "Login successful !"

	ErrLoginInvalidInput    = "The payload or input provided is invalid. Please check your request and try again."
	ErrLoginUserNotFound    = "User not found"
	ErrLoginInvalidPassword = "Password does not match the user's password."
)

func Login(c *fiber.Ctx) error {
	var db *sql.DB = conf.ConnectMariaDB()
	queries := sqlc.New(db)
	ctx := context.Background()
	var (
		REQ_IN   loginIn
		RESP_OUT loginOut
		RESP_ERR loginError
	)

	if err := c.BodyParser(&REQ_IN); err != nil {
		RESP_ERR.Ok = false
		RESP_ERR.ErrorMsg = ErrLoginInvalidInput
		errResp, _ := json.Marshal(RESP_ERR)
		return c.Status(fiber.StatusBadRequest).JSON(string(errResp))
	}

	isValid := utils.ValidateStruct(REQ_IN)
	if isValid != nil {
		RESP_ERR.Ok = false
		RESP_ERR.ErrorMsg = isValid.Error()
		errResp, _ := json.Marshal(RESP_ERR)
		return c.Status(fiber.StatusBadRequest).JSON(string(errResp))
	}

	userLoginRow, isUserExist := queries.UserLogin(ctx, REQ_IN.Username)
	if isUserExist != nil {
		RESP_ERR.Ok = false
		RESP_ERR.ErrorMsg = ErrLoginUserNotFound
		errorResp, _ := json.Marshal(RESP_ERR)
		return c.Status(fiber.StatusBadRequest).JSON(string(errorResp))
	}

	err := verifyPassword(REQ_IN.Password, userLoginRow.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		RESP_ERR.Ok = false
		RESP_ERR.ErrorMsg = ErrLoginInvalidPassword
		errorResp, _ := json.Marshal(RESP_ERR)
		return c.Status(fiber.StatusBadRequest).JSON(string(errorResp))
	}

	token, err := conf.GenerateJWT(userLoginRow.Username, userLoginRow.UserID, time.Now().AddDate(0, 2, 0))
	if err != nil {
		RESP_ERR.Ok = false
		RESP_ERR.ErrorMsg = "Error generate token"
		errorResp, _ := json.Marshal(RESP_ERR)
		return c.Status(fiber.StatusInternalServerError).JSON(string(errorResp))
	}

	RESP_OUT = loginOut{
		Ok:       true,
		Token:    token,
		Username: userLoginRow.Username,
		UserId:   userLoginRow.UserID,
		Message:  OutLoginMsg,
	}
	successResp, _ := json.Marshal(RESP_OUT)
	conf.SetJWTasCookie(c, token, time.Now().AddDate(0, 2, 0))

	defer db.Close()
	return c.Status(fiber.StatusOK).JSON(string(successResp))
}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
