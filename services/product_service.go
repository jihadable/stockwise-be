package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jihadable/stockwise-be/config"
	"github.com/jihadable/stockwise-be/model/entity"
	"github.com/jihadable/stockwise-be/model/request"

	"github.com/redis/go-redis/v9"

	"gorm.io/gorm"
)

type ProductService interface {
	AddProduct(image request.ImageRequest, product *entity.Product) (*entity.Product, error)
	GetProductsByUser(userId string) ([]*entity.Product, error)
	GetProductById(id string) (*entity.Product, error)
	UpdateProductById(id string, image request.ImageRequest, product *entity.Product) (*entity.Product, error)
	DeleteProductById(id string) error
}

type ProductServiceImpl struct {
	DB             *gorm.DB
	Redis          *redis.Client
	StorageService StorageService
}

func (service *ProductServiceImpl) AddProduct(image request.ImageRequest, product *entity.Product) (*entity.Product, error) {
	if image.File != nil {
		imagePath, err := service.StorageService.AddImage(image)
		if err != nil {
			return nil, err
		}
		product.Image = &imagePath
	}

	err := service.DB.Create(product).Error
	if err != nil {
		return nil, err
	}

	productJSON, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}

	service.Redis.Set(context.Background(), "product:"+product.Id, productJSON, time.Hour)

	return product, nil
}

func (service *ProductServiceImpl) GetProductsByUser(userId string) ([]*entity.Product, error) {
	products := []*entity.Product{}

	err := service.DB.Where("user_id = ?", userId).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (service *ProductServiceImpl) GetProductById(id string) (*entity.Product, error) {
	ctx := context.Background()
	redisKey := "product:" + id
	product := entity.Product{}

	productInRedis, err := service.Redis.Get(ctx, redisKey).Result()
	if err == nil && productInRedis != "" {
		err = json.Unmarshal([]byte(productInRedis), &product)

		if err == nil {
			return &product, nil
		}
	}

	err = service.DB.Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}

	productJSON, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}
	service.Redis.Set(ctx, redisKey, productJSON, time.Hour)

	return &product, nil
}

func (service *ProductServiceImpl) UpdateProductById(id string, image request.ImageRequest, product *entity.Product) (*entity.Product, error) {
	savedProduct := entity.Product{}

	result := service.DB.Where("id = ?", id).First(&savedProduct)

	err := result.Error
	if err != nil {
		return nil, err
	}

	if image.File != nil {
		if savedProduct.Image != nil {
			err = service.StorageService.DeleteImage(*savedProduct.Image)
			if err != nil {
				return nil, err
			}
		}

		imagePath, err := service.StorageService.AddImage(image)
		if err != nil {
			return nil, err
		}
		savedProduct.Image = &imagePath
	}

	savedProduct.Name = product.Name
	savedProduct.Price = product.Price
	savedProduct.Quantity = product.Quantity
	savedProduct.Category = product.Category
	savedProduct.Description = product.Description

	err = service.DB.Where("id = ?", id).Updates(&savedProduct).Error
	if err != nil {
		return nil, err
	}

	service.Redis.Del(context.Background(), "product:"+id)

	return &savedProduct, nil
}

func (service *ProductServiceImpl) DeleteProductById(id string) error {
	savedProduct := entity.Product{}

	err := service.DB.Where("id = ?", id).First(&savedProduct).Error
	if err != nil {
		return err
	}

	if savedProduct.Image != nil {
		err = service.StorageService.DeleteImage(*savedProduct.Image)
		if err != nil {
			return err
		}
	}

	err = service.DB.Delete(&savedProduct).Error
	if err != nil {
		return err
	}

	redisKey := "product:" + id
	service.Redis.Del(context.Background(), redisKey)

	return nil
}

func NewProductService(config *config.Config) ProductService {
	return &ProductServiceImpl{DB: config.DB, Redis: config.Redis, StorageService: NewStorageService()}
}
