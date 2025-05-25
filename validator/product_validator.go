package validator

import (
	"stockwise-be/model/request"
	"stockwise-be/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ProductValidator interface {
	ValidatePostProductRequest(productRequest request.ProductRequest) error
	ValidatePutProductRequest(productRequest request.ProductRequest) error
}

type ProductValidatorImpl struct {
	Validate *validator.Validate
}

func (validator *ProductValidatorImpl) ValidatePostProductRequest(productRequest request.ProductRequest) error {
	err := validator.Validate.Struct(productRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, utils.ParseValidationErrors(err))
	}
	return nil
}

func (validator *ProductValidatorImpl) ValidatePutProductRequest(productRequest request.ProductRequest) error {
	err := validator.Validate.Struct(productRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, utils.ParseValidationErrors(err))
	}
	return nil
}

func NewProductValidator() ProductValidator {
	return &ProductValidatorImpl{Validate: validator.New()}
}
