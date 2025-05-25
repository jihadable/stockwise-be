package handlers

import (
	"stockwise-be/model/request"
	"stockwise-be/services"
	"stockwise-be/utils"
	"stockwise-be/validator"

	"github.com/gofiber/fiber/v2"
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

	productRequest := request.ProductRequest{}

	err := ctx.BodyParser(&productRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Gagal menambahkan produk")
	}
	productRequest.UserId = ctx.Locals("user_id").(string)

	err = handler.Validator.ValidatePostProductRequest(productRequest)
	if err != nil {
		return err
	}

	product := *utils.RequestToProduct(&productRequest)

	productResponse, err := handler.Service.AddProduct(image, product)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"product": productResponse,
		},
	})
}

func (handler *ProductHandlerImpl) GetProductsByUser(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(string)
	productsResponse, err := handler.Service.GetProductsByUser(userId)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"products": productsResponse,
		},
	})
}

func (handler *ProductHandlerImpl) GetProductById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	productResponse, err := handler.Service.GetProductById(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"product": productResponse,
		},
	})
}

func (handler *ProductHandlerImpl) PutProductById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	image := ctx.Locals("image").(request.ImageRequest)
	productRequest := request.ProductRequest{}

	err := ctx.BodyParser(&productRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Gagal memperbarui produk")
	}

	err = handler.Validator.ValidatePutProductRequest(productRequest)
	if err != nil {
		return err
	}

	product := *utils.RequestToProduct(&productRequest)

	productResponse, err := handler.Service.UpdateProductById(id, image, product)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"product": productResponse,
		},
	})
}

func (handler *ProductHandlerImpl) DeleteProductById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := handler.Service.DeleteProductById(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}

func NewProductHandler(service services.ProductService, validator validator.ProductValidator) ProductHandler {
	return &ProductHandlerImpl{Service: service, Validator: validator}
}
