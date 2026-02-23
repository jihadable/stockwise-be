package validator

import (
	"github.com/jihadable/stockwise-be/model/request"

	"github.com/go-playground/validator/v10"
)

type ProductValidator interface {
	ValidatePostProductRequest(request request.ProductRequest) error
	ValidatePutProductRequest(request request.ProductRequest) error
}

type ProductValidatorImpl struct {
	*validator.Validate
}

func (validator *ProductValidatorImpl) ValidatePostProductRequest(request request.ProductRequest) error {
	return validator.Validate.Struct(request)
}

func (validator *ProductValidatorImpl) ValidatePutProductRequest(request request.ProductRequest) error {
	return validator.Validate.Struct(request)
}

func NewProductValidator() ProductValidator {
	return &ProductValidatorImpl{Validate: validator.New()}
}
