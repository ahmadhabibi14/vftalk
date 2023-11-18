package handlers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"vftalk/conf"
	"vftalk/models/database/sqlc"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func GetUserDataByUsername(c *fiber.Ctx) (sqlc.GetUserDataByUsernameRow, error) {
	u, _ := conf.GetUsernameFromJWT(c)
	username := fmt.Sprintf("%v", u)
	var db *sql.DB = conf.ConnectMariaDB()
	queries := sqlc.New(db)
	ctx := context.Background()

	userData, err := queries.GetUserDataByUsername(ctx, username)
	if err != nil {
		return sqlc.GetUserDataByUsernameRow{}, errors.New("User not found")
	}

	return userData, nil
}

type userDataout struct {
	Ok       bool                           `json:"ok"`
	UserData *sqlc.GetUserDataByUsernameRow `json:"userdata"`
	Message  string                         `json:"message"`
}

func GetUserData(c *fiber.Ctx) error {
	var U_OUT userDataout
	u, _ := conf.GetUsernameFromJWT(c)
	username := fmt.Sprintf("%v", u)
	var db *sql.DB = conf.ConnectMariaDB()
	queries := sqlc.New(db)
	ctx := context.Background()

	userData, err := queries.GetUserDataByUsername(ctx, username)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	U_OUT = userDataout{
		Ok:       true,
		UserData: &userData,
		Message:  `Successfully retrieving user data`,
	}
	ouResp, _ := json.Marshal(U_OUT)

	return c.Status(fiber.StatusOK).JSON(string(ouResp))
}
