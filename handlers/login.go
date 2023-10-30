package handlers

import (
	"log"
	"time"
	"vftalk/conf"

	json "github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
)

type (
	loginIn struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	loginOut struct {
		Token    string `json:"token"`
		Username string `json:"username"`
		Message  string `json:"message"`
	}
	loginError struct {
		ErrorMsg string `json:"error"`
	}
)

func Login(c *fiber.Ctx) error {
	var (
		in     loginIn
		out    loginOut
		errmsg loginError
	)

	if err := c.BodyParser(&in); err != nil {
		errmsg.ErrorMsg = "Invalid Input"
		errorResp, _ := json.Marshal(errmsg)
		return c.Status(fiber.StatusBadRequest).JSON(string(errorResp))
	}

	log.Println("login username = ", in.Username)

	token, err := conf.GenerateJWT(in.Username, time.Now().AddDate(0, 2, 0))
	if err != nil {
		errmsg.ErrorMsg = "Error generate token"
		errorResp, _ := json.Marshal(errmsg)
		return c.Status(fiber.StatusInternalServerError).JSON(string(errorResp))
	}

	out = loginOut{
		Token:    token,
		Username: in.Username,
		Message:  "Login success",
	}
	successResp, _ := json.Marshal(out)
	conf.SetJWTasCookie(c, token, time.Now().AddDate(0, 2, 0))
	return c.Status(fiber.StatusOK).JSON(string(successResp))
}
