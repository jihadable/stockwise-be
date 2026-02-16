package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jihadable/stockwise-be/helper/mapper"
	"github.com/jihadable/stockwise-be/model/entity"
	"github.com/jihadable/stockwise-be/model/request"
	"github.com/jihadable/stockwise-be/services"
	"github.com/jihadable/stockwise-be/validator"
)

type ProductHandler interface {
	PostProduct(ctx *fiber.Ctx) error
	GetProductsByUser(ctx *fiber.Ctx) error
	GetProductById(ctx *fiber.Ctx) error
	PutProductById(ctx *fiber.Ctx) error
	DeleteProductById(ctx *fiber.Ctx) error
}

type ProductHandlerImpl struct {
	Service   services.ProductService
	Validator validator.ProductValidator
}

func (handler *ProductHandlerImpl) PostProduct(ctx *fiber.Ctx) error {
	image := ctx.Locals("image").(request.ImageRequest)

	requestBody := request.ProductRequest{}

	err := ctx.BodyParser(&requestBody)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	requestBody.UserId = ctx.Locals("user_id").(string)

	err = handler.Validator.ValidatePostProductRequest(requestBody)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	product, err := handler.Service.AddProduct(image, &entity.Product{
		Name:        requestBody.Name,
		Category:    requestBody.Category,
		Price:       requestBody.Price,
		Quantity:    requestBody.Quantity,
		Description: requestBody.Description,
		UserId:      requestBody.UserId,
	})
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Fail to create product")
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"product": mapper.ProductToResponse(product),
		},
	})
}

func (handler *ProductHandlerImpl) GetProductsByUser(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(string)
	products, err := handler.Service.GetProductsByUser(userId)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Products not found")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"products": mapper.ProductsToResponses(products),
		},
	})
}

func (handler *ProductHandlerImpl) GetProductById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	product, err := handler.Service.GetProductById(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Product not found")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"product": mapper.ProductToResponse(product),
		},
	})
}

func (handler *ProductHandlerImpl) PutProductById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	image := ctx.Locals("image").(request.ImageRequest)
	requestBody := request.ProductRequest{}

	err := ctx.BodyParser(&requestBody)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	err = handler.Validator.ValidatePutProductRequest(requestBody)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	product, err := handler.Service.UpdateProductById(id, image, &entity.Product{
		Name:        requestBody.Name,
		Category:    requestBody.Category,
		Price:       requestBody.Price,
		Quantity:    requestBody.Quantity,
		Description: requestBody.Description,
	})
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Fail to update product")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"product": mapper.ProductToResponse(product),
		},
	})
}

func (handler *ProductHandlerImpl) DeleteProductById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := handler.Service.DeleteProductById(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Fail to delete product")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}

func NewProductHandler(service services.ProductService, validator validator.ProductValidator) ProductHandler {
	return &ProductHandlerImpl{Service: service, Validator: validator}
}
