package apis

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
		response = JSONResponse(fiber.StatusBadRequest, "Invalid csrf state", "")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	code := c.FormValue("code")
	oAuthToken, err := a.OAuth.Google.Exchange(ctx, code)
	if err != nil {
		response = JSONResponse(fiber.StatusBadRequest, "Code exchange failed", "")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	r, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + oAuthToken.AccessToken)
	if err != nil {
		response = JSONResponse(fiber.StatusBadRequest, "Failed to get user info", "")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	defer r.Body.Close()
	var userInfo map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&userInfo); err != nil {
		response = JSONResponse(fiber.StatusBadRequest, "Failed to decode user info", "")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	GOOGLE_id := userInfo["id"].(string)
	GOOGLE_email := userInfo["email"].(string)
	GOOGLE_username := strings.Split(GOOGLE_email, "@")[0]
	GOOGLE_fullname := userInfo["name"].(string)
	GOOGLE_avatar := userInfo["picture"].(string)

	resp, err := http.Get(GOOGLE_avatar)
	if err != nil || resp.StatusCode != fiber.StatusOK {
		response = JSONResponse(fiber.StatusInternalServerError, "Failed to save user info", "")
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	defer resp.Body.Close()

	imgPath := fmt.Sprintf("contents/img/avatars/%v.png", GOOGLE_id)
	file, _ := os.Create(imgPath)
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		response = JSONResponse(fiber.StatusInternalServerError, "Failed to save user avatar", "")
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	GOOGLE_avatar = fmt.Sprintf("/img/avatars/%v.png", GOOGLE_id)
	in := services.InUser_OAuthGoogle{
		Username: GOOGLE_username,
		FullName: GOOGLE_fullname,
		Email:    GOOGLE_email,
		Avatar:   GOOGLE_avatar,
		GoogleID: GOOGLE_id,
	}

	user := services.NewUser(a.Db, a.Log)
	token, err := user.OAuthGoogle(ctx, in)
	if err != nil {
		response = JSONResponse(fiber.StatusBadRequest, err.Error(), "")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	configs.SetJWTasCookie(c, token)
	return c.Redirect("/", fiber.StatusPermanentRedirect)
}
