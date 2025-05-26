package routes

import (
	"github.com/jihadable/stockwise-be/handlers"
	"github.com/jihadable/stockwise-be/middlewares"
	"github.com/jihadable/stockwise-be/services"
	"github.com/jihadable/stockwise-be/validator"

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
