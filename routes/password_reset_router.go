package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jihadable/stockwise-be/config"
	"github.com/jihadable/stockwise-be/handlers"
	"github.com/jihadable/stockwise-be/services"
	"github.com/jihadable/stockwise-be/validator"
)

func RegisterPasswordResetRoutes(api fiber.Router, config *config.Config) {
	service := services.NewPasswordResetService(config)
	validator := validator.NewPasswordResetValidator()
	handler := handlers.NewPasswordResetHandler(service, validator)
	route := api.Group("/password-resets")

	route.Post("/send-password-reset-email", handler.SendPasswordResetEmail)
	route.Post("/reset-password", handler.ResetPassword)
}
