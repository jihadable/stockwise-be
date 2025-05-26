package api

import (
	"net/http/httptest"
	"testing"

	"github.com/jihadable/stockwise-be/model/request"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPostUserWithValidPayload(t *testing.T) {
	requestBody := RequestBodyParser(request.UserRequest{
		Username: "username 1",
		Email:    "email1@mail.com",
		Password: "password 1",
	})
	request := httptest.NewRequest(fiber.MethodPost, "/api/users/register", requestBody)
	request.Header.Set("Content-Type", "application/json")

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusCreated, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)
	assert.Equal(t, "success", responseBody["status"])

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	token, ok := data["token"].(string)
	assert.True(t, ok)
	assert.NotEmpty(t, token)
	JWT = token

	user, ok := data["user"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, user["id"])
	assert.Equal(t, "username 1", user["username"])
	assert.Equal(t, "email1@mail.com", user["email"])

	t.Log("✅")
}

func TestPostUserWithInvalidPayload(t *testing.T) {
	requestBody := RequestBodyParser(request.UserRequest{})
	request := httptest.NewRequest(fiber.MethodPost, "/api/users/register", requestBody)
	request.Header.Set("Content-Type", "application/json")

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)
	assert.Equal(t, "fail", responseBody["status"])
	assert.NotEmpty(t, responseBody["message"])

	t.Log("✅")
}

func TestVerifyUserWithValidPayload(t *testing.T) {
	requestBody := RequestBodyParser(request.UserRequest{
		Email:    "email1@mail.com",
		Password: "password 1",
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

	token, ok := data["token"].(string)
	assert.True(t, ok)
	assert.NotEmpty(t, token)
	JWT = token

	user, ok := data["user"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, user["id"])
	assert.Equal(t, "username 1", user["username"])
	assert.Equal(t, "email1@mail.com", user["email"])
	assert.Equal(t, "bio 1", user["bio"])

	t.Log("✅")
}

func TestVerifyUserWithInvalidPayload(t *testing.T) {
	requestBody := RequestBodyParser(request.UserRequest{})
	request := httptest.NewRequest(fiber.MethodPost, "/api/users/login", requestBody)
	request.Header.Set("Content-Type", "application/json")

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)
	assert.Equal(t, "fail", responseBody["status"])
	assert.NotEmpty(t, responseBody["message"])

	t.Log("✅")
}

func TestGetUserByIdWithToken(t *testing.T) {
	request := httptest.NewRequest(fiber.MethodGet, "/api/users", nil)
	request.Header.Set("Authorization", "Bearer "+JWT)

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)
	assert.Equal(t, "success", responseBody["status"])

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	user, ok := data["user"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, user["id"])
	assert.Equal(t, "username 1", user["username"])
	assert.Equal(t, "email1@mail.com", user["email"])
	assert.Equal(t, "bio 1", user["bio"])

	t.Log("✅")
}

func TestGetUserByIdWithoutToken(t *testing.T) {
	request := httptest.NewRequest(fiber.MethodGet, "/api/users", nil)

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusUnauthorized, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)
	assert.Equal(t, "fail", responseBody["status"])
	assert.Equal(t, "Token tidak ditemukan", responseBody["message"])

	t.Log("✅")
}

func TestUpdateUserByIdWithValidPayload(t *testing.T) {
	requestBody := RequestBodyParser(request.UpdateUserRequest{
		Username: "username 2",
	})
	request := httptest.NewRequest(fiber.MethodPut, "/api/users", requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+JWT)

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)
	assert.Equal(t, "success", responseBody["status"])

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	user, ok := data["user"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, user["id"])
	assert.Equal(t, "username 2", user["username"])
	assert.Equal(t, "email1@mail.com", user["email"])

	t.Log("✅")
}

func TestUpdateUserByIdWithInvalidPayload(t *testing.T) {
	requestBody := RequestBodyParser(request.UserRequest{})
	request := httptest.NewRequest(fiber.MethodPut, "/api/users", requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+JWT)

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)
	assert.Equal(t, "fail", responseBody["status"])
	assert.NotEmpty(t, responseBody["message"])

	t.Log("✅")
}
