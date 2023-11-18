package handlers

import (
	"context"
	"database/sql"
	"log"
	"time"

	json "github.com/goccy/go-json"

	"vftalk/conf"
	"vftalk/models/database/sqlc"
	"vftalk/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type (
	registerIn struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8,containsany=!@#?*%&>_<}-{+"`
		Username string `json:"username" validate:"required,omitempty,min=4"`
		Fullname string `json:"fullname" validate:"required"`
	}
	registerOut struct {
		Ok       bool   `json:"ok"`
		Token    string `json:"token"`
		Username string `json:"username"`
		Message  string `json:"message"`
	}
	registerError struct {
		Ok       bool   `json:"ok"`
		ErrorMsg string `json:"error"`
	}
)

const (
	OutRegisterMsg = "Register successful !"

	ErrRegisterInvalidInput  = "The payload or input provided is invalid. Please check your request and try again."
	ErrRegisterUsernameExist = "Username already exists, try another one"
	ErrRegisterEmailExist    = "Email already exist, try another one"
	ErrRegisterGenerateToken = "Error generate session token"
)

func Register(c *fiber.Ctx) error {
	var db *sql.DB = conf.ConnectMariaDB()
	queries := sqlc.New(db)
	ctx := context.Background()
	var (
		REQ_IN   registerIn
		RESP_OUT registerOut
		RESP_ERR registerError
	)

	if err := c.BodyParser(&REQ_IN); err != nil {
		RESP_ERR.Ok = false
		RESP_ERR.ErrorMsg = ErrRegisterInvalidInput
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

	_, isUsernameExist := queries.GetUserByUsername(ctx, REQ_IN.Username)
	if isUsernameExist == nil {
		RESP_ERR.Ok = false
		RESP_ERR.ErrorMsg = ErrRegisterUsernameExist
		errResp, _ := json.Marshal(RESP_ERR)
		return c.Status(fiber.StatusBadRequest).JSON(string(errResp))
	}

	_, isEmailExist := queries.GetUserByEmail(ctx, REQ_IN.Email)
	if isEmailExist == nil {
		RESP_ERR.Ok = false
		RESP_ERR.ErrorMsg = ErrRegisterEmailExist
		errResp, _ := json.Marshal(RESP_ERR)
		return c.Status(fiber.StatusBadRequest).JSON(string(errResp))
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(REQ_IN.Password), bcrypt.DefaultCost)
	userData := sqlc.CreateNewUserParams{
		UserID:   utils.GenerateRandomID(20),
		Username: REQ_IN.Username,
		FullName: REQ_IN.Fullname,
		Email:    REQ_IN.Email,
		Password: string(hashedPassword),
	}
	err := queries.CreateNewUser(ctx, userData)
	if err != nil {
		log.Println(`Error queries.CreateNewUser : `, err)
		RESP_ERR.Ok = false
		RESP_ERR.ErrorMsg = err.Error()
		errResp, _ := json.Marshal(RESP_ERR)
		return c.Status(fiber.StatusInternalServerError).JSON(string(errResp))
	}

	token, err := conf.GenerateJWT(REQ_IN.Username, time.Now().AddDate(0, 2, 0))
	if err != nil {
		RESP_ERR.Ok = false
		RESP_ERR.ErrorMsg = ErrRegisterGenerateToken
		errResp, _ := json.Marshal(RESP_ERR)
		return c.Status(fiber.StatusInternalServerError).JSON(string(errResp))
	}

	RESP_OUT = registerOut{
		Ok:       true,
		Token:    token,
		Username: REQ_IN.Username,
		Message:  OutRegisterMsg,
	}
	outResp, _ := json.Marshal(RESP_OUT)
	conf.SetJWTasCookie(c, token, time.Now().AddDate(0, 2, 0))

	// mail.SendUserRegisterMail()
	defer db.Close()
	return c.Status(fiber.StatusCreated).JSON(string(outResp))
}
