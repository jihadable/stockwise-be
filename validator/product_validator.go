package validator

import (
	"github.com/jihadable/stockwise-be/model/request"

	"github.com/go-playground/validator/v10"
)

type ProductValidator interface {
	ValidatePostProductRequest(productRequest request.ProductRequest) error
	ValidatePutProductRequest(productRequest request.ProductRequest) error
}

type ProductValidatorImpl struct {
	Validate *validator.Validate
}

func (validator *ProductValidatorImpl) ValidatePostProductRequest(request request.ProductRequest) error {
	err := validator.Validate.Struct(request)
	if err != nil {
		return err
	}
	return nil
}

func (validator *ProductValidatorImpl) ValidatePutProductRequest(request request.ProductRequest) error {
	err := validator.Validate.Struct(request)
	if err != nil {
		return err
	}
	return nil
}

func NewProductValidator() ProductValidator {
	return &ProductValidatorImpl{Validate: validator.New()}
}
