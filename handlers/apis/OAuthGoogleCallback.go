package apis

import (
	"encoding/json"
	"net/http"
	"strings"
	"vftalk/configs"
	"vftalk/services"

	"github.com/gofiber/fiber/v2"
)

func (a *ApisHandler) OAuthGoogle(c *fiber.Ctx) error {
	ctx := c.Context()
	response := HTTPResponse{}

	state := c.FormValue("state")
	if state == "" {
		response = HTTPResponse{
			Code:   fiber.StatusBadRequest,
			Status: STATUS_BADREQUEST,
			Errors: "Invalid csrf state",
			Data:   "",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	code := c.FormValue("code")
	oAuthToken, err := a.OAuth.Google.Exchange(ctx, code)
	if err != nil {
		response = HTTPResponse{
			Code:   fiber.StatusBadRequest,
			Status: STATUS_BADREQUEST,
			Errors: "Code exchange failed",
			Data:   "",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	r, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + oAuthToken.AccessToken)
	if err != nil {
		response = HTTPResponse{
			Code:   fiber.StatusBadRequest,
			Status: STATUS_BADREQUEST,
			Errors: "Failed to get user info",
			Data:   "",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	defer r.Body.Close()
	var userInfo map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&userInfo); err != nil {
		response = HTTPResponse{
			Code:   fiber.StatusBadRequest,
			Status: STATUS_BADREQUEST,
			Errors: "Failed to decode user info",
			Data:   "",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	GOOGLE_id := userInfo["id"].(string)
	GOOGLE_email := userInfo["email"].(string)
	GOOGLE_username := strings.Split(GOOGLE_email, "@")[0]
	GOOGLE_fullname := userInfo["name"].(string)
	GOOGLE_avatar := userInfo["picture"].(string)

	in := services.InUser_OAuthCreate{
		UserID:   GOOGLE_id,
		Username: GOOGLE_username,
		FullName: GOOGLE_fullname,
		Email:    GOOGLE_email,
		Avatar:   GOOGLE_avatar,
	}

	user := services.NewUser(a.Db, a.Log)
	token, err := user.OAuthCreateUser(ctx, in)
	if err != nil {
		response = HTTPResponse{
			Code:   fiber.StatusBadRequest,
			Status: STATUS_BADREQUEST,
			Errors: err.Error(),
			Data:   "",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	configs.SetJWTasCookie(c, token)
	response = HTTPResponse{
		Code:   fiber.StatusOK,
		Status: STATUS_OK,
		Errors: "",
		Data:   "Login successful !",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
