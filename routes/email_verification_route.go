package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jihadable/stockwise-be/config"
	"github.com/jihadable/stockwise-be/handlers"
	"github.com/jihadable/stockwise-be/middlewares"
	"github.com/jihadable/stockwise-be/services"
	"github.com/jihadable/stockwise-be/validator"
)

func RegisterEmailVerificationRoutes(api fiber.Router, config *config.Config) {
	service := services.NewEmailVerificationService(config)
	validator := validator.NewEmailVerificationValidator()
	handler := handlers.NewEmailVerificationHandler(service, validator)
	route := api.Group("/email-verifications")

	route.Post("/send-email-verification", middlewares.AuthMiddleware(), handler.SendEmailVerification)
	route.Post("/verify-email", handler.VerifyEmail)
}
