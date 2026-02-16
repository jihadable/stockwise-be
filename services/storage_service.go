package services

import (
	"os"

	"github.com/jihadable/stockwise-be/model/request"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type StorageService interface {
	AddImage(image request.ImageRequest) (string, error)
	DeleteImage(imageName string) error
}

type StorageServiceImpl struct {
	Client  *fiber.Client
	BaseURL string
	APIKey  string
}

func (service *StorageServiceImpl) AddImage(image request.ImageRequest) (string, error) {
	imageName := uuid.NewString() + image.Ext
	url := service.BaseURL + "/" + imageName

	request := service.Client.Post(url)
	request.Set("Authorization", "Bearer "+service.APIKey)
	request.Set("Content-Type", getContentType(image.Ext))

	request.BodyStream(image.File, -1)

	status, _, err := request.String()
	if err != nil || status != fiber.StatusOK {
		return "", err[0]
	}

	return imageName, nil
}

func (service *StorageServiceImpl) DeleteImage(imageName string) error {
	url := service.BaseURL + "/" + imageName

	request := service.Client.Delete(url)
	request.Set("Authorization", "Bearer "+service.APIKey)

	status, _, err := request.String()
	if err != nil || status != fiber.StatusOK {
		return err[0]
	}

	return nil
}

func getContentType(ext string) string {
	switch ext {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	default:
		return "application/octet-stream"
	}
}

func NewStorageService() StorageService {
	return &StorageServiceImpl{
		Client:  fiber.AcquireClient(),
		BaseURL: os.Getenv("IMAGE_API_ENDPOINT"),
		APIKey:  os.Getenv("IMAGE_API_ENDPOINT_API_KEY"),
	}
}
