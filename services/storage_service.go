package services

import (
	"os"
	"stockwise-be/model/request"

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
	if err != nil {
		return "", fiber.NewError(fiber.StatusBadRequest, "Gagal mengunggah gambar")
	}

	if status != fiber.StatusOK {
		return "", fiber.NewError(fiber.StatusBadRequest, "Gagal mengunggah gambar")
	}

	return imageName, nil
}

func (service *StorageServiceImpl) DeleteImage(imageName string) error {
	url := service.BaseURL + "/" + imageName

	request := service.Client.Delete(url)
	request.Set("Authorization", "Bearer "+service.APIKey)

	status, _, err := request.String()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Gagal menghapus gambar")
	}

	if status != fiber.StatusOK {
		return fiber.NewError(fiber.StatusBadRequest, "Gagal bro")
	}

	// [
	// 	map[
	// 		description:description 1
	// 		id:9a12977d-76e1-4393-8abf-0c961d68bbd2
	// 		image:<nil>
	// 		name:product 1
	// 		price:1
	// 		quantity:1
	// 		slug:122f8d1f-c2b4-46c1-9ff1-11a05de53bb0
	// 	]
	// 	map[
	// 		description:description 2
	// 		id:04370c75-0738-4e20-b699-7e81546e5aef
	// 		image:67b6114c-d4dd-4074-a2fd-5ab82066c3ee.png
	// 		name:product 2
	// 		price:2
	// 		quantity:2
	// 		slug:44b360a3-fd13-4559-a46f-ce24e8333297
	// 	]
	// ]

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
