package utils

import (
	"github.com/jihadable/stockwise-be/model/entity"
	"github.com/jihadable/stockwise-be/model/request"
	"github.com/jihadable/stockwise-be/model/response"
)

func RequestToProduct(request *request.ProductRequest) *entity.Product {
	return &entity.Product{
		Name:        request.Name,
		Category:    request.Category,
		Price:       request.Price,
		Quantity:    request.Quantity,
		Description: request.Description,
		UserId:      request.UserId,
	}
}

func ProductToResponse(product *entity.Product) *response.ProductResponse {
	return &response.ProductResponse{
		Id:          product.Id,
		Name:        product.Name,
		Category:    product.Category,
		Price:       product.Price,
		Quantity:    product.Quantity,
		Image:       product.Image,
		Description: product.Description,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func ProductsToResponses(products []entity.Product) []response.ProductResponse {
	responses := make([]response.ProductResponse, len(products))
	for i, p := range products {
		responses[i] = *ProductToResponse(&p)
	}
	return responses
}
