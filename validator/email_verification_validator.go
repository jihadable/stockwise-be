package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/jihadable/stockwise-be/model/request"
)

type EmailVerificationValidator interface {
	ValidateVerifyEmailRequest(request *request.VerifyEmailRequest) error
}

type EmailVerificationValidatorImpl struct {
	Validate *validator.Validate
}

func (validator *EmailVerificationValidatorImpl) ValidateVerifyEmailRequest(request *request.VerifyEmailRequest) error {
	err := validator.Validate.Struct(request)
	if err != nil {
		return err
	}
	return nil
}

func NewEmailVerificationValidator() EmailVerificationValidator {
	return &EmailVerificationValidatorImpl{Validate: validator.New()}
}
