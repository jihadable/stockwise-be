package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/jihadable/stockwise-be/model/request"
)

type EmailVerificationValidator interface {
	ValidateSendEmailVerificationRequest(request *request.SendEmailVerificationRequest) error
	ValidateVerifyEmailRequest(request *request.VerifyEmailRequest) error
}

type EmailVerificationValidatorImpl struct {
	Validate *validator.Validate
}

func (validator *EmailVerificationValidatorImpl) ValidateSendEmailVerificationRequest(request *request.SendEmailVerificationRequest) error {
	err := validator.Validate.Struct(request)
	if err != nil {
		return err
	}
	return nil
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
