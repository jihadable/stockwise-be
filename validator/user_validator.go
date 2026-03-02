package validator

import (
	"github.com/jihadable/stockwise-be/model/request"

	"github.com/go-playground/validator/v10"
)

type UserValidator interface {
	ValidatePostUserRequest(request request.UserRequest) error
	ValidatePutUserRequest(request request.UpdateUserRequest) error
	ValidateVerifyUserRequest(request request.LoginRequest) error
	ValidateUpdatePasswordRequest(request request.UpdatePasswordRequest) error
}

type UserValidatorImpl struct {
	*validator.Validate
}

func (validator *UserValidatorImpl) ValidatePostUserRequest(request request.UserRequest) error {
	return validator.Validate.Struct(request)
}

func (validator *UserValidatorImpl) ValidatePutUserRequest(request request.UpdateUserRequest) error {
	return validator.Validate.Struct(request)
}

func (validator *UserValidatorImpl) ValidateVerifyUserRequest(request request.LoginRequest) error {
	return validator.Validate.Struct(request)
}

func (validator *UserValidatorImpl) ValidateUpdatePasswordRequest(request request.UpdatePasswordRequest) error {
	return validator.Validate.Struct(request)
}

func NewUserValidator() UserValidator {
	return &UserValidatorImpl{Validate: validator.New()}
}
