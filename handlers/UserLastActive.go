package handlers

import (
	"context"
	"database/sql"
	"encoding/json"

	"vftalk/conf"
	"vftalk/models/database/sqlc"

	"github.com/gofiber/fiber/v2"
)

type (
	userLastActiveIn struct {
		UserId string `json:"user_id" form:"user_id"`
	}
	userLastActiveOut struct {
		Ok      bool   `json:"ok"`
		Message string `json:"message"`
	}
	userLastActiveError struct {
		Ok       bool   `json:"ok"`
		ErrorMsg string `json:"error"`
	}
)

const (
	OutUserLastActive_Msg = "userLastActive successful !"

	ErrUserLastActive_InvalidInput = "The payload or input provided is invalid. Please check your request and try again."
	ErrUserLastActive_UserNotFound = "User not found"
)

func UserLastActives(c *fiber.Ctx) error {
	var db *sql.DB = conf.ConnectMariaDB()
	queries := sqlc.New(db)
	ctx := context.Background()
	var (
		REQ_IN   userLastActiveIn
		RESP_OUT userLastActiveOut
		RESP_ERR userLastActiveError
	)

	if err := c.BodyParser(&REQ_IN); err != nil {
		RESP_ERR.Ok = false
		RESP_ERR.ErrorMsg = ErrUserLastActive_InvalidInput
		errResp, _ := json.Marshal(RESP_ERR)
		return c.Status(fiber.StatusBadRequest).JSON(string(errResp))
	}

	_, isUserIdExist := queries.GetUserDataByUserId(ctx, REQ_IN.UserId)
	if isUserIdExist == nil {
		RESP_ERR.Ok = false
		RESP_ERR.ErrorMsg = ErrUserLastActive_UserNotFound
		errResp, _ := json.Marshal(RESP_ERR)
		return c.Status(fiber.StatusBadRequest).JSON(string(errResp))
	}

	RESP_OUT = userLastActiveOut{
		Ok:      true,
		Message: OutUserLastActive_Msg,
	}
	outResp, _ := json.Marshal(RESP_OUT)
	defer db.Close()
	return c.Status(fiber.StatusCreated).JSON(string(outResp))
}
