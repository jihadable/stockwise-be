package routes

import (
	"github.com/jihadable/stockwise-be/config"
	"github.com/jihadable/stockwise-be/handlers"
	"github.com/jihadable/stockwise-be/middlewares"
	"github.com/jihadable/stockwise-be/services"
	"github.com/jihadable/stockwise-be/validator"

	"github.com/gofiber/fiber/v2"
)

func RegisterProductRoutes(api fiber.Router, config *config.Config) {
	service := services.NewProductService(config)
	validator := validator.NewProductValidator()
	handler := handlers.NewProductHandler(service, validator)
	route := api.Group("/products", middlewares.AuthMiddleware())

	route.Post("/", middlewares.GetImageFile(), handler.PostProduct)
	route.Get("/", handler.GetProductsByUser)
	route.Get("/:id", handler.GetProductById)
	route.Put("/:id", middlewares.GetImageFile(), handler.PutProductById)
	route.Delete("/:id", handler.DeleteProductById)
}
