package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/jihadable/stockwise-be/model/request"
)

type PasswordResetValidator interface {
	ValidateSendPasswordResetRequest(request request.SendPasswordResetEmailRequest) error
	ValidateResetPasswordRequest(request request.ResetPasswordRequest) error
}

type PasswordResetValidatorImpl struct {
	*validator.Validate
}

func (validator *PasswordResetValidatorImpl) ValidateSendPasswordResetRequest(request request.SendPasswordResetEmailRequest) error {
	return validator.Validate.Struct(request)
}

func (validator *PasswordResetValidatorImpl) ValidateResetPasswordRequest(request request.ResetPasswordRequest) error {
	return validator.Validate.Struct(request)
}

func NewPasswordResetValidator() PasswordResetValidator {
	return &PasswordResetValidatorImpl{Validate: validator.New()}
}
