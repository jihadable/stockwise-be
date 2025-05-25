package api

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/gofiber/fiber/v2"
)

func RequestBodyParser(body any) io.Reader {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	return bytes.NewReader(jsonBody)
}

func ResponseBodyParser(body io.ReadCloser) fiber.Map {
	defer body.Close()

	result := fiber.Map{}
	err := json.NewDecoder(body).Decode(&result)
	if err != nil {
		panic(err)
	}

	return result
}
