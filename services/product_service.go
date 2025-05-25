package services

import (
	"stockwise-be/model/entity"
	"stockwise-be/model/request"
	"stockwise-be/model/response"
	"stockwise-be/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductService interface {
	AddProduct(image request.ImageRequest, product entity.Product) (response.ProductResponse, error)
	GetProductsByUser(userId string) ([]response.ProductResponse, error)
	GetProductById(id string) (response.ProductResponse, error)
	UpdateProductById(id string, image request.ImageRequest, product entity.Product) (response.ProductResponse, error)
	DeleteProductById(id string) error
}

type ProductServiceImpl struct {
	DB             *gorm.DB
	StorageService StorageService
}

func (service *ProductServiceImpl) AddProduct(image request.ImageRequest, product entity.Product) (response.ProductResponse, error) {
	if image.File != nil {
		imagePath, err := service.StorageService.AddImage(image)
		if err != nil {
			return response.ProductResponse{}, err
		}
		product.Image = &imagePath
	}
	product.Id = uuid.NewString()
	result := service.DB.Create(&product)

	err := result.Error
	if err != nil {
		return response.ProductResponse{}, fiber.NewError(fiber.StatusBadRequest, "Gagal bro")
	}

	return *utils.ProductToResponse(&product), nil
}

func (service *ProductServiceImpl) GetProductsByUser(userId string) ([]response.ProductResponse, error) {
	products := []entity.Product{}

	result := service.DB.Where("user_id = ?", userId).Find(&products)

	err := result.Error
	if err != nil {
		return []response.ProductResponse{}, fiber.NewError(fiber.StatusBadRequest, "Gagal mendapatkan produk")
	}

	return utils.ProductsToResponses(products), nil
}

func (service *ProductServiceImpl) GetProductById(id string) (response.ProductResponse, error) {
	product := entity.Product{}

	result := service.DB.Where("id = ?", id).First(&product)

	err := result.Error
	if err != nil {
		return response.ProductResponse{}, fiber.NewError(fiber.StatusNotFound, "Gagal mendapatkan produk. Produk tidak ditemukan")
	}

	return *utils.ProductToResponse(&product), nil
}

func (service *ProductServiceImpl) UpdateProductById(id string, image request.ImageRequest, product entity.Product) (response.ProductResponse, error) {
	savedProduct := entity.Product{}

	result := service.DB.Where("id = ?", id).First(&savedProduct)

	err := result.Error
	if err != nil {
		return response.ProductResponse{}, fiber.NewError(fiber.StatusNotFound, "Gagal memperbarui produk. Produk tidak ditemukan")
	}

	if image.File != nil {
		if savedProduct.Image != nil {
			err = service.StorageService.DeleteImage(*savedProduct.Image)
			if err != nil {
				return response.ProductResponse{}, err
			}
		}

		imagePath, err := service.StorageService.AddImage(image)
		if err != nil {
			return response.ProductResponse{}, err
		}
		savedProduct.Image = &imagePath
	}

	savedProduct.Name = product.Name
	savedProduct.Price = product.Price
	savedProduct.Quantity = product.Quantity
	savedProduct.Description = product.Description

	result = service.DB.Where("id = ?", id).Updates(&savedProduct)

	err = result.Error
	if err != nil {
		return response.ProductResponse{}, fiber.NewError(fiber.StatusNotFound, "Gagal memperbarui produk")
	}

	return *utils.ProductToResponse(&savedProduct), nil
}

func (service *ProductServiceImpl) DeleteProductById(id string) error {
	savedProduct := entity.Product{}

	result := service.DB.Where("id = ?", id).First(&savedProduct)

	err := result.Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Gagal menghapus produk. Produk tidak ditemukan")
	}

	if savedProduct.Image != nil {
		err = service.StorageService.DeleteImage(*savedProduct.Image)
		if err != nil {
			return err
		}
	}

	result = service.DB.Where("id = ?", id).Delete(&entity.Product{})

	err = result.Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Gagal menghapus produk")
	}

	return nil
}

func NewProductService(db *gorm.DB) ProductService {
	return &ProductServiceImpl{DB: db, StorageService: NewStorageService()}
}
