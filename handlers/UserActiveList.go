package handlers

import (
	"context"
	"database/sql"
	"vftalk/conf"
	"vftalk/models/database/sqlc"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
)

type (
	userActiveListOut struct {
		Ok             bool                     `json:"ok"`
		Message        string                   `json:"message"`
		UserActiveList []sqlc.ListUserActiveRow `json:"users"`
	}
	userActiveListError struct {
		Ok       bool   `json:"ok"`
		ErrorMsg string `json:"error"`
	}
)

const (
	OutUserActiveList_Msg   = "User active list"
	OutUserActiveList_Empty = "No user atives"

	ErrUserActiveList_InvalidInput = ""
)

func GetUserActiveLists(c *fiber.Ctx) error {
	var db *sql.DB = conf.ConnectMariaDB()
	queries := sqlc.New(db)
	ctx := context.Background()
	var (
		RESP_OUT userActiveListOut
		RESP_ERR userActiveListError
	)

	usersList, err := queries.ListUserActive(ctx)
	if err != nil {
		RESP_ERR.Ok = false
		RESP_ERR.ErrorMsg = OutUserActiveList_Empty
		errorResp, _ := json.Marshal(RESP_ERR)
		return c.Status(fiber.StatusBadRequest).JSON(string(errorResp))
	}

	RESP_OUT = userActiveListOut{
		Ok:             true,
		Message:        OutUserActiveList_Msg,
		UserActiveList: usersList,
	}
	successResp, _ := json.Marshal(RESP_OUT)

	defer db.Close()
	return c.Status(fiber.StatusOK).JSON(string(successResp))
}
