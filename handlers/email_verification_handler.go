package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jihadable/stockwise-be/model/request"
	"github.com/jihadable/stockwise-be/services"
	"github.com/jihadable/stockwise-be/validator"
)

type EmailVerificationHandler interface {
	SendEmailVerification(ctx *fiber.Ctx) error
	VerifyEmail(ctx *fiber.Ctx) error
}

type EmailVerificationHandlerImpl struct {
	Service   services.EmailVerificationService
	Validator validator.EmailVerificationValidator
}

func (handler *EmailVerificationHandlerImpl) SendEmailVerification(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(string)

	err := handler.Service.SendEmailVerification(userId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Fail to send email verification")
	}

	return nil
}

func (handler *EmailVerificationHandlerImpl) VerifyEmail(ctx *fiber.Ctx) error {
	requestBody := request.VerifyEmailRequest{}

	err := ctx.BodyParser(&requestBody)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	err = handler.Validator.ValidateVerifyEmailRequest(&requestBody)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	err = handler.Service.VerifyEmail(requestBody.Token)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Fail to verify email")
	}

	return nil
}

func NewEmailVerificationHandler(service services.EmailVerificationService) EmailVerificationHandler {
	return &EmailVerificationHandlerImpl{Service: service}
}
