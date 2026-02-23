package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/jihadable/stockwise-be/model/request"
)

type EmailVerificationValidator interface {
	ValidateVerifyEmailRequest(request request.VerifyEmailRequest) error
}

type EmailVerificationValidatorImpl struct {
	*validator.Validate
}

func (validator *EmailVerificationValidatorImpl) ValidateVerifyEmailRequest(request request.VerifyEmailRequest) error {
	return validator.Validate.Struct(request)
}

func NewEmailVerificationValidator() EmailVerificationValidator {
	return &EmailVerificationValidatorImpl{Validate: validator.New()}
}
