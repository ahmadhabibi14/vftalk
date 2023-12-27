package configs

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OAuthConf struct {
	Google *oauth2.Config
}

func EnvOAuth() (res OAuthConf) {
	GOOGLE_REDIRECT_URL := os.Getenv("GOOGLE_OAUTH_REDIRECT_URL_LOCAL")
	if os.Getenv("WEB_ENV") == "prod" {
		GOOGLE_REDIRECT_URL = os.Getenv("GOOGLE_OAUTH_REDIRECT_URL")
	}
	GOOGLE_OAUTH := &oauth2.Config{
		RedirectURL:  GOOGLE_REDIRECT_URL,
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		Scopes: []string{
			"openid",
			"profile",
			"email",
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	res.Google = GOOGLE_OAUTH
	return
}
