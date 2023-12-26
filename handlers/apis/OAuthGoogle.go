package apis

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"time"
// 	"vftalk/configs"
// 	"vftalk/services"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/joho/godotenv"
// 	"golang.org/x/oauth2"
// 	"golang.org/x/oauth2/google"
// )

// var (
// 	zlog              = configs.InitLogger()
// 	GoogleOauthConfig *oauth2.Config
// )

// func init() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		zlog.Error().
// 			Str("ERROR", err.Error()).
// 			Msg("cannot load .env files")
// 	}
// 	redirectURL := "http://localhost:8000/api/oauth/google"
// 	if os.Getenv("WEB_ENV") == "prod" {
// 		redirectURL = "https://vftalk.my.id/api/oauth/google"
// 	}
// 	GoogleOauthConfig = &oauth2.Config{
// 		RedirectURL:  redirectURL,
// 		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
// 		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
// 		Scopes: []string{
// 			"openid",
// 			"profile",
// 			"email",
// 			"https://www.googleapis.com/auth/userinfo.email",
// 			"https://www.googleapis.com/auth/userinfo.profile",
// 		},
// 		Endpoint: google.Endpoint,
// 	}
// }

// func (a *ApisHandler) OAuthGoogle(c *fiber.Ctx) error {
// 	ctx := c.Context()
// 	response := HTTPResponse{}

// 	state := c.FormValue("state")
// 	_ = state
// 	code := c.FormValue("code")
// 	oAuthToken, err := GoogleOauthConfig.Exchange(ctx, code)
// 	if err != nil {
// 		response = HTTPResponse{
// 			Code:   fiber.StatusBadRequest,
// 			Status: STATUS_BADREQUEST,
// 			Errors: "Code exchange failed",
// 			Data:   "",
// 		}
// 		return c.Status(fiber.StatusBadRequest).JSON(response)
// 	}
// 	r, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + oAuthToken.AccessToken)
// 	if err != nil {
// 		response = HTTPResponse{
// 			Code:   fiber.StatusBadRequest,
// 			Status: STATUS_BADREQUEST,
// 			Errors: "Failed to get user info",
// 			Data:   "",
// 		}
// 		return c.Status(fiber.StatusBadRequest).JSON(response)
// 	}
// 	defer r.Body.Close()
// 	var userInfo map[string]interface{}
// 	if err := json.NewDecoder(r.Body).Decode(&userInfo); err != nil {
// 		response = HTTPResponse{
// 			Code:   fiber.StatusBadRequest,
// 			Status: STATUS_BADREQUEST,
// 			Errors: "Failed to decode user info",
// 			Data:   "",
// 		}
// 		return c.Status(fiber.StatusBadRequest).JSON(response)
// 	}

// 	GOOGLE_id := userInfo["id"].(string)
// 	GOOGLE_email := userInfo["email"].(string)
// 	GOOGLE_username := strings.Split(GOOGLE_email, "@")[0]
// 	GOOGLE_fullname := userInfo["name"].(string)
// 	GOOGLE_avatar := userInfo["picture"].(string)

// 	in := services.InUser_Create{
// 		UserID: GOOGLE_id,
// 		Username: GOOGLE_username,
// 		FullName: GOOGLE_fullname,
// 		Email: GOOGLE_email,
// 		Password: "--------",
// 	}

// 	user := services.NewUser(a.Db, a.Log)
// 	_, err = user.FindById(ctx, services.InUser_FindById{UserID: in.UserID})
// 	if err == nil {
// 		token, _ := configs.GenerateJWT(in.Username, GOOGLE_id, time.Now().AddDate(0, 2, 0))
// 		configs.SetJWTasCookie(c, token)
// 		response = HTTPResponse{
// 			Code: fiber.StatusOK,
// 			Status: STATUS_OK,
// 			Errors: "",
// 			Data: "Login successful !",
// 		}
// 		return c.Status(fiber.StatusOK).JSON(response)
// 	}
// 	token, err := user.CreateUser(ctx, in)
// 	if err != nil {
// 		response = HTTPResponse{
// 			Code:   fiber.StatusBadRequest,
// 			Status: STATUS_BADREQUEST,
// 			Errors: err.Error(),
// 			Data:   "",
// 		}
// 		return c.Status(fiber.StatusBadRequest).JSON(response)
// 	}

// 	configs.SetJWTasCookie(c, token)
// 	response = HTTPResponse{
// 		Code:   fiber.StatusBadRequest,
// 		Status: STATUS_BADREQUEST,
// 		Errors: err.Error(),
// 		Data:   "",
// 	}
// 	return c.Status(fiber.StatusBadRequest).JSON(response)
// }
