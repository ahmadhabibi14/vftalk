package apis

import (
	"fmt"
	"net/http"
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
		response = HTTPResponse{
			Code:   fiber.StatusUnauthorized,
			Status: http.StatusText(fiber.StatusUnauthorized),
			Errors: ERROR_UNAUTHORIZED,
			Data:   "",
		}
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	imgFile, err := c.FormFile("avatar")
	if err != nil {
		a.Log.Error().Str("Error", err.Error()).Msg("Cannot get image file when update user avatar")
		response = HTTPResponse{
			Code:   fiber.StatusBadRequest,
			Status: http.StatusText(fiber.StatusBadRequest),
			Errors: ERROR_INVALIDPAYLOAD,
			Data:   "",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	imgValid := utils.ImageValidation(imgFile)
	if imgValid != nil {
		response = HTTPResponse{
			Code:   fiber.StatusBadRequest,
			Status: http.StatusText(fiber.StatusBadRequest),
			Errors: imgValid.Error(),
			Data:   "",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	imgPath := fmt.Sprintf("contents/img/avatars/%v.png", userId)
	imgSave := c.SaveFile(imgFile, imgPath)
	if imgSave != nil {
		response = HTTPResponse{
			Code:   fiber.StatusInternalServerError,
			Status: http.StatusText(fiber.StatusInternalServerError),
			Errors: "Something went wrong",
			Data:   "",
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	imgPathStored := fmt.Sprintf("/img/avatars/%v.png", userId)

	in.UserID = userId.(string)
	in.Avatar = imgPathStored

	user := services.NewUser(a.Db, a.Log)
	updateAvatar := user.UpdateAvatar(ctx, in)
	if updateAvatar != nil {
		response = HTTPResponse{
			Code:   fiber.StatusInternalServerError,
			Status: http.StatusText(fiber.StatusInternalServerError),
			Errors: updateAvatar.Error(),
			Data:   "",
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response = HTTPResponse{
		Code:   fiber.StatusOK,
		Status: http.StatusText(fiber.StatusOK),
		Errors: "",
		Data: struct {
			Msg    string `json:"message"`
			Avatar string `json:"avatar"`
		}{
			Msg:    "Profile picture updated !",
			Avatar: imgPathStored,
		},
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
