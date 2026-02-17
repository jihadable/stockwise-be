package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jihadable/stockwise-be/model/request"
	"github.com/jihadable/stockwise-be/services"
)

type EmailVerificationHandler interface {
	SendEmailVerification(ctx *fiber.Ctx) error
	VerifyEmail(ctx *fiber.Ctx) error
}

type EmailVerificationHandlerImpl struct {
	Service services.EmailVerificationService
}

func (handler *EmailVerificationHandlerImpl) SendEmailVerification(ctx *fiber.Ctx) error {
	requestBody := request.SendEmailVerificationRequest{}

	err := ctx.BodyParser(&requestBody)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "")
	}
	panic("")
}

func (handler *EmailVerificationHandlerImpl) VerifyEmail(ctx *fiber.Ctx) error {
	panic("")
}

func NewEmailVerificationHandler(service services.EmailVerificationService) EmailVerificationHandler {
	return &EmailVerificationHandlerImpl{Service: service}
}
