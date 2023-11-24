package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gofiber/fiber/v2"
)

func TestLogin(t *testing.T) {
	app := fiber.New()
	app.Post("/api/login", Login)

	payload := `{"username":"ahmadhabibi14","password":"pass1234"}`
	req := httptest.NewRequest(http.MethodPost, "/api/login", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	// Http response
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "Login failed")
	// Do something with results:
	getCookie := resp.Header.Get("Set-Cookie") // Get cookie from header response
	body, _ := io.ReadAll(resp.Body)           // Get all response body
	if resp.StatusCode == fiber.StatusOK {
		t.Logf("Response Header (Cookie) : %s\n", string(getCookie))
		t.Logf("Response Body : %s\n", string(body))
	} else {
		t.Errorf("Error: Invalid request: %v\n", resp.Body)
	}
}
