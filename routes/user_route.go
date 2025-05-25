package routes

import (
	"stockwise-be/handlers"
	"stockwise-be/middlewares"
	"stockwise-be/services"
	"stockwise-be/validator"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterUserRoutes(api fiber.Router, db *gorm.DB) {
	service := services.NewUserService(db)
	validator := validator.NewUserValidator()
	handler := handlers.NewUserHandler(service, validator)
	route := api.Group("/users")

	route.Post("/register", handler.PostUser)
	route.Get("/", middlewares.AuthMiddleware(service), handler.GetUserById)
	route.Put("/", middlewares.AuthMiddleware(service), handler.PutUserById)
	route.Post("/login", handler.VerifyUser)
}
