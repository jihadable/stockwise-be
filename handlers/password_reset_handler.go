package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jihadable/stockwise-be/model/request"
	"github.com/jihadable/stockwise-be/services"
	"github.com/jihadable/stockwise-be/validator"
)

type PasswordResetHandler interface {
	SendPasswordResetEmail(ctx *fiber.Ctx) error
	ResetPassword(ctx *fiber.Ctx) error
}

type PasswordResetHandlerImpl struct {
	Service   services.PasswordResetService
	Validator validator.PasswordResetValidator
}

func (handler *PasswordResetHandlerImpl) SendPasswordResetEmail(ctx *fiber.Ctx) error {
	requestBody := request.SendPasswordResetEmailRequest{}

	err := ctx.BodyParser(&requestBody)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	err = handler.Validator.ValidateSendPasswordResetRequest(requestBody)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	err = handler.Service.SendPasswordResetEmail(requestBody.Email)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Fail to send password reset email")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}

func (handler *PasswordResetHandlerImpl) ResetPassword(ctx *fiber.Ctx) error {
	requestBody := request.ResetPasswordRequest{}

	err := ctx.BodyParser(&requestBody)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	err = handler.Validator.ValidateResetPasswordRequest(requestBody)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	err = handler.Service.ResetPassword(requestBody.Token, requestBody.NewPassword)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Fail to reset password")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}

func NewPasswordResetHandler(service services.PasswordResetService, validator validator.PasswordResetValidator) PasswordResetHandler {
	return &PasswordResetHandlerImpl{Service: service, Validator: validator}
}
