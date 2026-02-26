package api

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/jihadable/stockwise-be/model/request"
	"github.com/stretchr/testify/assert"
)

func TestSendPasswordResetEmail(t *testing.T) {
	requestBody := RequestBodyParser(request.SendPasswordResetEmailRequest{
		Email: "jihadumar1021@gmail.com",
	})
	request := httptest.NewRequest(fiber.MethodPost, "/api/password-resets/send-password-reset-email", requestBody)
	request.Header.Set("Content-Type", "application/json")

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(request.Body)

	assert.Equal(t, "success", responseBody["status"])
}

func TestSendPasswordResetEmailWithInvalidEmail(t *testing.T) {
	requestBody := RequestBodyParser(request.SendPasswordResetEmailRequest{
		Email: "invalidemail",
	})
	request := httptest.NewRequest(fiber.MethodPost, "/api/password-resets/send-password-reset-email", requestBody)
	request.Header.Set("Content-Type", "application/json")

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	assert.Equal(t, "fail", responseBody["status"])
	assert.NotEmpty(t, responseBody["message"])
}

func TestResetPassword(t *testing.T) {
	// requestBody := RequestBodyParser(request.ResetPasswordRequest{
	// 	Token:       os.Getenv("PASSWORD_RESET_TOKEN"),
	// 	NewPassword: "topsecret123",
	// })
}

func TestResetPasswordWithInvalidToken(t *testing.T) {
	requestBody := RequestBodyParser(request.ResetPasswordRequest{
		Token:       "invalidtoken",
		NewPassword: "topsecret123",
	})
	request := httptest.NewRequest(fiber.MethodPost, "/api/password-resets/reset-password", requestBody)
	request.Header.Set("Content-Type", "application/json")

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	assert.Equal(t, "fail", responseBody["status"])
	assert.NotEmpty(t, responseBody["message"])
}
