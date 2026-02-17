package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jihadable/stockwise-be/config"
	"github.com/jihadable/stockwise-be/handlers"
	"github.com/jihadable/stockwise-be/services"
)

func RegisterEmailVerificationRoutes(api fiber.Router, config *config.Config) {
	service := services.NewEmailVerificationService(config)
	handler := handlers.NewEmailVerificationHandler(service)
	route := api.Group("/email-verification")

	route.Post("/send-email-verification", handler.SendEmailVerification)
	route.Post("/verify-email", handler.VerifyEmail)
}
