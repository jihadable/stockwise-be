package handlers

import (
	"github.com/gofiber/fiber/v2"
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
	panic("")
}

func (handler *PasswordResetHandlerImpl) ResetPassword(ctx *fiber.Ctx) error {
	panic("")
}

func NewPasswordResetHandler(service services.PasswordResetService, validator validator.PasswordResetValidator) PasswordResetHandler {
	return &PasswordResetHandlerImpl{Service: service, Validator: validator}
}
