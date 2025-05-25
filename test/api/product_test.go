package api

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"path/filepath"
	"stockwise-be/model/request"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var productIdWithoutImage string
var productIdWithImage string

func TestPostProductWithoutImage(t *testing.T) {
	requestBody := RequestBodyParser(request.ProductRequest{
		Name:        "product 1",
		Category:    "category 1",
		Price:       1,
		Quantity:    1,
		Description: "description 1",
	})
	request := httptest.NewRequest(fiber.MethodPost, "/api/products", requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+JWT)

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusCreated, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)
	assert.Equal(t, "success", responseBody["status"])

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	product, ok := data["product"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, product["id"])
	productIdWithoutImage = product["id"].(string)
	assert.Equal(t, "product 1", product["name"])
	assert.Equal(t, "category 1", product["category"])
	assert.Equal(t, float64(1), product["price"])
	assert.Equal(t, float64(1), product["quantity"])
	assert.Nil(t, product["image"])
	assert.Equal(t, "description 1", product["description"])

	t.Log("✅")
}

func TestPostProductWithImage(t *testing.T) {
	filePath := filepath.Join("..", "..", "static", "ss.png")
	file, err := os.Open(filePath)
	assert.Nil(t, err)
	defer file.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	_ = writer.WriteField("name", "product 2")
	_ = writer.WriteField("category", "category 2")
	_ = writer.WriteField("price", "2")
	_ = writer.WriteField("quantity", "2")
	_ = writer.WriteField("description", "description 2")

	part, err := writer.CreateFormFile("image", filepath.Base(filePath))
	assert.Nil(t, err)
	_, err = io.Copy(part, file)
	assert.Nil(t, err)

	writer.Close()

	request := httptest.NewRequest(fiber.MethodPost, "/api/products", &requestBody)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("Authorization", "Bearer "+JWT)

	response, err := App.Test(request, 2*1000)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusCreated, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)
	assert.Equal(t, "success", responseBody["status"])

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	product, ok := data["product"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, product["id"])
	productIdWithImage = product["id"].(string)
	assert.Equal(t, "product 2", product["name"])
	assert.Equal(t, "category 2", product["category"])
	assert.Equal(t, float64(2), product["price"])
	assert.Equal(t, float64(2), product["quantity"])
	assert.NotNil(t, product["image"])
	assert.Equal(t, "description 2", product["description"])

	t.Log("✅")
}

func TestPostProductWithInvalidPayload(t *testing.T) {
	requestBody := RequestBodyParser(request.ProductRequest{})
	request := httptest.NewRequest(fiber.MethodPost, "/api/products", requestBody)
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

func TestGetProductsByUserWithToken(t *testing.T) {
	request := httptest.NewRequest(fiber.MethodGet, "/api/products", nil)
	request.Header.Set("Authorization", "Bearer "+JWT)

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)
	assert.Equal(t, "success", responseBody["status"])

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	products, ok := data["products"].([]any)
	assert.True(t, ok)
	assert.NotEmpty(t, products)

	t.Log("✅")
}

func TestGetProductsByUserWithoutToken(t *testing.T) {
	request := httptest.NewRequest(fiber.MethodGet, "/api/products", nil)

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusUnauthorized, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)
	assert.Equal(t, "fail", responseBody["status"])
	assert.Equal(t, "Token tidak ditemukan", responseBody["message"])

	t.Log("✅")
}

func TestGetProductByIdWithValidId(t *testing.T) {
	request := httptest.NewRequest(fiber.MethodGet, "/api/products/"+productIdWithoutImage, nil)
	request.Header.Set("Authorization", "Bearer "+JWT)

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)
	assert.Equal(t, "success", responseBody["status"])

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	product, ok := data["product"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, product["id"])
	assert.Equal(t, "product 1", product["name"])
	assert.Equal(t, float64(1), product["price"])
	assert.Equal(t, float64(1), product["quantity"])
	assert.Nil(t, product["image"])
	assert.Equal(t, "description 1", product["description"])

	t.Log("✅")
}

func TestGetProductByIdWithInvalidId(t *testing.T) {
	request := httptest.NewRequest(fiber.MethodGet, "/api/products/xxx", nil)
	request.Header.Set("Authorization", "Bearer "+JWT)

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusNotFound, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)
	assert.Equal(t, "fail", responseBody["status"])
	assert.Equal(t, "Gagal mendapatkan produk. Produk tidak ditemukan", responseBody["message"])

	t.Log("✅")
}

func TestUpdateProductByIdWithoutImage(t *testing.T) {
	requestBody := RequestBodyParser(request.ProductRequest{
		Name:        "update product 1",
		Category:    "update category 1",
		Price:       1,
		Quantity:    1,
		Description: "update description 1",
	})
	request := httptest.NewRequest(fiber.MethodPut, "/api/products/"+productIdWithImage, requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+JWT)

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)
	assert.Equal(t, "success", responseBody["status"])

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	product, ok := data["product"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, product["id"])
	assert.Equal(t, "update product 1", product["name"])
	assert.Equal(t, "update category 1", product["category"])
	assert.Equal(t, float64(1), product["price"])
	assert.Equal(t, float64(1), product["quantity"])
	assert.NotNil(t, product["image"])
	assert.Equal(t, "update description 1", product["description"])

	t.Log("✅")
}

func TestUpdateProductByIdWithImage(t *testing.T) {
	filePath := filepath.Join("..", "..", "static", "ss.png")
	file, err := os.Open(filePath)
	assert.Nil(t, err)
	defer file.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	_ = writer.WriteField("name", "update product 2")
	_ = writer.WriteField("category", "update category 2")
	_ = writer.WriteField("price", "2")
	_ = writer.WriteField("quantity", "2")
	_ = writer.WriteField("description", "update description 2")

	part, err := writer.CreateFormFile("image", filepath.Base(filePath))
	assert.Nil(t, err)
	_, err = io.Copy(part, file)
	assert.Nil(t, err)

	writer.Close()

	request := httptest.NewRequest(fiber.MethodPut, "/api/products/"+productIdWithoutImage, &requestBody)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("Authorization", "Bearer "+JWT)

	response, err := App.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)
	assert.Equal(t, "success", responseBody["status"])

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	product, ok := data["product"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, product["id"])
	assert.Equal(t, "update product 2", product["name"])
	assert.Equal(t, "update category 2", product["category"])
	assert.Equal(t, float64(2), product["price"])
	assert.Equal(t, float64(2), product["quantity"])
	assert.NotNil(t, product["image"])
	assert.Equal(t, "update description 2", product["description"])

	t.Log("✅")
}

func TestUpdateProductByIdWithInvalidPayload(t *testing.T) {
	requestBody := RequestBodyParser(request.ProductRequest{})
	request := httptest.NewRequest(fiber.MethodPut, "/api/products/"+productIdWithoutImage, requestBody)
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

func TestDeleteProductById(t *testing.T) {
	request := httptest.NewRequest(fiber.MethodDelete, "/api/products/"+productIdWithoutImage, nil)
	request.Header.Set("Authorization", "Bearer "+JWT)

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)
	assert.Equal(t, "success", responseBody["status"])

	t.Log("✅")
}
