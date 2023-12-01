package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gofiber/fiber/v2"
)

func TestRegister(t *testing.T) {
	app := fiber.New()
	app.Post("/api/register", Register)

	payload := struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Fullname string `json:"fullname"`
		Password string `json:"password"`
	}{
		Email:    "gojo8139@proton.me",
		Username: "satoru1993",
		Fullname: "Gojo Satoru",
		Password: "satoru93#",
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		t.Error("Failed to convert JSON data")
		return
	}

	req := httptest.NewRequest(http.MethodPost, "/api/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Http response
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusCreated, resp.StatusCode, "Register failed")
	// Do something with results:
	getCookie := resp.Header.Get("Set-Cookie") // Get cookie from header response
	body, _ := io.ReadAll(resp.Body)           // Get all response body
	if resp.StatusCode == fiber.StatusCreated {
		t.Logf("Response Header (Cookie) : %s\n", string(getCookie))
		t.Logf("Response Body : %s\n", string(body))
	} else {
		t.Errorf("Error: %v\n", string(body))
	}
}
