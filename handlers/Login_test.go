package handlers

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goccy/go-json"

	"github.com/stretchr/testify/assert"

	"github.com/gofiber/fiber/v2"
)

func TestLogin(t *testing.T) {
	app := fiber.New()
	app.Post("/api/login", Login)

	payload := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		Username: "satoru93",
		Password: "satoru93#",
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		t.Error("Failed to convert JSON data")
		return
	}

	req := httptest.NewRequest(http.MethodPost, "/api/login", bytes.NewBuffer(jsonData))
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
		t.Errorf("Error: %v\n", string(body))
	}
}
