package api

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/jihadable/stockwise-be/model/request"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	requestBody := RequestBodyParser(request.UserRequest{
		Email:    "jihadumar1021@gmail.com",
		Password: os.Getenv("PRIVATE_PASSWORD"),
	})
	request := httptest.NewRequest(fiber.MethodPost, "/api/users/login", requestBody)
	request.Header.Set("Content-Type", "application/json")

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)
	assert.Equal(t, "success", responseBody["status"])

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	jwt, ok := data["jwt"].(string)
	assert.True(t, ok)
	assert.NotEmpty(t, jwt)
	JWT = jwt

	user, ok := data["user"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, user["id"])
	assert.Equal(t, "jihadumar", user["username"])
	assert.Equal(t, "jihadumar1021@gmail.com", user["email"])

	t.Log("✅")
}

func TestSendEmailVerification(t *testing.T) {
	request := httptest.NewRequest(fiber.MethodPost, "/api/email-verifications/send-email-verification", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+JWT)

	response, err := App.Test(request, -1)

	assert.Nil(t, err)
	assert.Equal(t, response.StatusCode, fiber.StatusOK)

	responseBody := ResponseBodyParser(response.Body)
	assert.Equal(t, "success", responseBody["status"])

	t.Log("✅")
}

func TestSendEmailVerificationWithoutJWT(t *testing.T) {
	request := httptest.NewRequest(fiber.MethodPost, "/api/email-verifications/send-email-verification", nil)
	request.Header.Set("Content-Type", "application/json")

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusUnauthorized, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)
	assert.Equal(t, "fail", responseBody["status"])
	assert.NotEmpty(t, responseBody["message"])

	t.Log("✅")
}

func TestVerifyEmail(t *testing.T) {
	requestBody := RequestBodyParser(request.VerifyEmailRequest{
		Token: os.Getenv("EMAIL_VERIFICATION_TOKEN"),
	})
	request := httptest.NewRequest(fiber.MethodPost, "/api/email-verifications/verify-email", requestBody)
	request.Header.Set("Content-Type", "application/json")

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	assert.Equal(t, "success", responseBody["status"])

	t.Log("✅")
}

func TestVerifyEmailWithoutToken(t *testing.T) {
	request := httptest.NewRequest(fiber.MethodPost, "/api/email-verifications/verify-email", nil)
	request.Header.Set("Content-Type", "application/json")

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	assert.Equal(t, "fail", responseBody["status"])
	assert.NotEmpty(t, responseBody["message"])

	t.Log("✅")
}
