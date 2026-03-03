package routes

import (
	"github.com/jihadable/stockwise-be/config"
	"github.com/jihadable/stockwise-be/handlers"
	"github.com/jihadable/stockwise-be/middlewares"
	"github.com/jihadable/stockwise-be/services"
	"github.com/jihadable/stockwise-be/validator"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(api fiber.Router, config *config.Config) {
	service := services.NewUserService(config)
	validator := validator.NewUserValidator()
	handler := handlers.NewUserHandler(service, validator)
	route := api.Group("/users")

	route.Post("/register", handler.PostUser)
	route.Get("/", middlewares.AuthMiddleware(), handler.GetUserById)
	route.Put("/", middlewares.AuthMiddleware(), handler.PutUserById)
	route.Post("/login", handler.VerifyUser)
	route.Post("/change-password", middlewares.AuthMiddleware(), handler.UpdatePassword)
}
