package apis

import (
	"fmt"
	"vftalk/configs"
	"vftalk/services"
	"vftalk/utils"

	"github.com/gofiber/fiber/v2"
)

func (a *ApisHandler) UpdateAvatar(c *fiber.Ctx) error {
	ctx := c.Context()
	response := HTTPResponse{}
	in := services.InUser_UpdateAvatar{}

	userId, err := configs.GetUserIdFromJWTfunc(c)
	if err != nil {
		response = JSONResponse(fiber.StatusUnauthorized, ERROR_UNAUTHORIZED, "")
		c.ClearCookie(configs.AUTH_COOKIE)
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	imgFile, err := c.FormFile("avatar")
	if err != nil {
		a.Log.Error().Str("Error", err.Error()).Msg("Cannot get image file when update user avatar")
		response = JSONResponse(fiber.StatusBadRequest, ERROR_INVALIDPAYLOAD, "")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	imgValid := utils.ImageValidation(imgFile)
	if imgValid != nil {
		response = JSONResponse(fiber.StatusBadRequest, imgValid.Error(), "")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	imgPath := fmt.Sprintf("contents/img/avatars/%v.png", userId)
	imgSave := c.SaveFile(imgFile, imgPath)
	if imgSave != nil {
		response = JSONResponse(fiber.StatusInternalServerError, "Something went wrong", "")
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	imgPathStored := fmt.Sprintf("/img/avatars/%v.png", userId)

	in.UserID = userId.(string)
	in.Avatar = imgPathStored

	user := services.NewUser(a.Db, a.Log)
	updateAvatar := user.UpdateAvatar(ctx, in)
	if updateAvatar != nil {
		response = JSONResponse(fiber.StatusInternalServerError, updateAvatar.Error(), "")
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	var data = struct {
		Msg    string `json:"message"`
		Avatar string `json:"avatar"`
	}{
		Msg:    "Profile picture updated !",
		Avatar: imgPathStored,
	}
	response = JSONResponse(fiber.StatusOK, "Something went wrong", data)
	return c.Status(fiber.StatusOK).JSON(response)
}
