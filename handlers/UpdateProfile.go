package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"vftalk/conf"
	"vftalk/models/database/sqlc"
	"vftalk/utils"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
)

type (
	updateProfileIn struct {
		FullName string `json:"full_name" validate:"required"`
		Location string `json:"location" validate:"required"`
		Website  string `json:"website" validate:"required"`
	}
	updateProfileOut struct {
		Ok      bool   `json:"ok"`
		Message string `json:"message"`
	}
	updateProfileError struct {
		Ok       bool   `json:"ok"`
		ErrorMsg string `json:"error"`
	}
)

const (
	OutUpdateProfile_Msg = "Profile updated"

	ErrUpdateProfile_InvalidInput = "The payload or input provided is invalid. Please check your request and try again."
	ErrUpdateProfile_InvalidUser  = "You are unauthorized to do this operation!"
	ErrUpdateProfile_ServerError  = "Something went wrong!"
)

func UpdateProfile(c *fiber.Ctx) error {
	var db *sql.DB = conf.ConnectMariaDB()
	defer db.Close()
	queries := sqlc.New(db)
	ctx := context.Background()
	var (
		REQ_IN   updateProfileIn
		RESP_OUT updateProfileOut
		RESP_ERR updateProfileError
	)

	if err := c.BodyParser(&REQ_IN); err != nil {
		RESP_ERR.Ok = false
		RESP_ERR.ErrorMsg = ErrUpdateProfile_InvalidInput
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

	userId, err := conf.GetUserIdFromJWTfunc(c)
	if err != nil {
		RESP_ERR.Ok = false
		RESP_ERR.ErrorMsg = ErrUpdateProfile_InvalidUser
		errResp, _ := json.Marshal(RESP_ERR)
		return c.Status(fiber.StatusBadRequest).JSON(string(errResp))
	}

	userData := sqlc.UpdateUserProfileParams{
		FullName: REQ_IN.FullName,
		Location: REQ_IN.Location,
		Website:  REQ_IN.Website,
		UserID:   fmt.Sprintf("%v", userId),
	}
	upProfErr := queries.UpdateUserProfile(ctx, userData)
	if upProfErr != nil {
		RESP_ERR.Ok = false
		RESP_ERR.ErrorMsg = ErrUpdateProfile_ServerError
		errResp, _ := json.Marshal(RESP_ERR)
		return c.Status(fiber.StatusBadRequest).JSON(string(errResp))
	}

	RESP_OUT = updateProfileOut{
		Ok:      true,
		Message: OutUpdateProfile_Msg,
	}
	outResp, _ := json.Marshal(RESP_OUT)
	return c.Status(fiber.StatusOK).JSON(string(outResp))
}
