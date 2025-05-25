package validator

import (
	"stockwise-be/model/request"
	"stockwise-be/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserValidator interface {
	ValidatePostUserRequest(request *request.UserRequest) error
	ValidatePutUserRequest(request *request.UpdateUserRequest) error
	ValidateVerifyUserRequest(request *request.LoginRequest) error
}

type UserValidatorImpl struct {
	Validate *validator.Validate
}

func (validator *UserValidatorImpl) ValidatePostUserRequest(request *request.UserRequest) error {
	err := validator.Validate.Struct(request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, utils.ParseValidationErrors(err))
	}
	return nil
}

func (validator *UserValidatorImpl) ValidatePutUserRequest(request *request.UpdateUserRequest) error {
	err := validator.Validate.Struct(request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, utils.ParseValidationErrors(err))
	}
	return nil
}

func (validator *UserValidatorImpl) ValidateVerifyUserRequest(request *request.LoginRequest) error {
	err := validator.Validate.Struct(request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, utils.ParseValidationErrors(err))
	}
	return nil
}

func NewUserValidator() UserValidator {
	return &UserValidatorImpl{Validate: validator.New()}
}
