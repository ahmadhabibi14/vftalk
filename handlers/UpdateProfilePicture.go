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
	updateProfilePictureOut struct {
		Ok        bool   `json:"ok"`
		AvatarUrl string `json:"avatarUrl"`
		Message   string `json:"message"`
	}
	updateProfilePictureError struct {
		Ok       bool   `json:"ok"`
		ErrorMsg string `json:"error"`
	}
)

const (
	OutUpdateProfilePicture_Msg = "Profile picture updated"

	ErrUpdateProfilePicture_InvalidInput = "The payload or input provided is invalid. Please check your request and try again."
	ErrUpdateProfilePicture_ImgTooLarge  = "Image too large, try to resize your image!"
	ErrUpdateProfilePicture_InvalidUser  = "You are unauthorized to do this operation!"
	ErrUpdateProfilePicture_ServerError  = "Something went wrong!"
)

func UpdateProfilePicture(c *fiber.Ctx) error {
	var db *sql.DB = conf.ConnectMariaDB()
	queries := sqlc.New(db)
	ctx := context.Background()
	var (
		RESP_OUT updateProfilePictureOut
		RESP_ERR updateProfilePictureError
	)
	userId, err := conf.GetUserIdFromJWTfunc(c)
	if err != nil {
		RESP_ERR.Ok = false
		RESP_ERR.ErrorMsg = ErrUpdateProfilePicture_InvalidUser
		errResp, _ := json.Marshal(RESP_ERR)
		return c.Status(fiber.StatusBadRequest).JSON(string(errResp))
	}

	form, err := c.MultipartForm()
	if err != nil {
		RESP_ERR.Ok = false
		RESP_ERR.ErrorMsg = ErrUpdateProfilePicture_InvalidInput
		errResp, _ := json.Marshal(RESP_ERR)
		return c.Status(fiber.StatusBadRequest).JSON(string(errResp))
	}
	imgFiles := form.File["avatar"]
	imgValid := utils.ImageValidation(imgFiles[0])
	if imgValid != nil {
		RESP_ERR.Ok = false
		RESP_ERR.ErrorMsg = imgValid.Error()
		errResp, _ := json.Marshal(RESP_ERR)
		return c.Status(fiber.StatusBadRequest).JSON(string(errResp))
	}

	imgPath := fmt.Sprintf("uploads/img/avatars/%v.png", userId)
	imgSave := c.SaveFile(imgFiles[0], imgPath)
	if imgSave != nil {
		RESP_ERR.Ok = false
		RESP_ERR.ErrorMsg = ErrUpdateProfilePicture_ServerError
		errResp, _ := json.Marshal(RESP_ERR)
		return c.Status(fiber.StatusInternalServerError).JSON(string(errResp))
	}

	imgPathStored := fmt.Sprintf("/img/avatars/%v.png", userId)
	// Set image avatar to database
	updateAvatarParam := sqlc.UpdateUserAvatarParams{
		Avatar: imgPathStored,
		UserID: fmt.Sprintf("%v", userId),
	}
	execSqlUpUserAvatar := queries.UpdateUserAvatar(ctx, updateAvatarParam)
	if execSqlUpUserAvatar != nil {
		RESP_ERR.Ok = false
		RESP_ERR.ErrorMsg = ErrUpdateProfilePicture_ServerError
		errResp, _ := json.Marshal(RESP_ERR)
		return c.Status(fiber.StatusInternalServerError).JSON(string(errResp))
	}

	RESP_OUT = updateProfilePictureOut{
		Ok:        true,
		AvatarUrl: imgPathStored,
		Message:   OutUpdateProfilePicture_Msg,
	}
	outResp, _ := json.Marshal(RESP_OUT)

	defer db.Close()
	return c.Status(fiber.StatusOK).JSON(string(outResp))
}
