package utils

import (
	"stockwise-be/model/entity"
	"stockwise-be/model/request"
	"stockwise-be/model/response"
)

func RequestToProduct(request *request.ProductRequest) *entity.Product {
	return &entity.Product{
		Name:        request.Name,
		Price:       request.Price,
		Quantity:    request.Quantity,
		Description: request.Description,
		UserId:      request.UserId,
	}
}

func ProductToResponse(product *entity.Product) *response.ProductResponse {
	return &response.ProductResponse{
		Id:          product.Id,
		Slug:        product.Slug,
		Name:        product.Name,
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
