// if you find `Ws` in the beginning of function name,
// that means it uses websocket

package configs

import (
	"os"
	"time"

	"fmt"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(username, user_id string, expirationTime time.Time) (tokenString string, err error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"username":   username,
		"user_id":    user_id,
		"exp":        expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "Generate JWT Error", err
	}
	return tokenString, nil
}

func TokenValid(c *fiber.Ctx) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			c.ClearCookie()
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		c.ClearCookie()
		return err
	}
	return nil
}

func ExtractToken(c *fiber.Ctx) string {
	jwtToken := c.Cookies(AUTH_COOKIE, "")
	if jwtToken == "" {
		jwtToken = c.Get("X-API-KEY", "")

		if jwtToken == "" {
			return ""
		}
	}
	return jwtToken
}

func WsExtractToken(c *websocket.Conn) string {
	jwtToken := c.Cookies(AUTH_COOKIE, "")
	if jwtToken == "" {
		jwtToken = c.Headers("X-API-KEY", "")
		if jwtToken == "" {
			return ""
		}
	}
	return jwtToken
}

func GetUsernameFromJWT(c *fiber.Ctx) (interface{}, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uname := claims["username"]
		return uname, nil
	}
	return "", nil
}

func GetUserIdFromJWTfunc(c *fiber.Ctx) (interface{}, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid := claims["user_id"]
		return uid, nil
	}
	return "", nil
}

func WsGetUsernameFromJWT(c *websocket.Conn) (interface{}, error) {
	tokenString := WsExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uname := claims["username"]
		return uname, nil
	}
	return "", nil
}
