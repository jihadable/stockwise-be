package routes

import (
	"github.com/jihadable/stockwise-be/handlers"
	"github.com/jihadable/stockwise-be/middlewares"
	"github.com/jihadable/stockwise-be/services"
	"github.com/jihadable/stockwise-be/validator"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterProductRoutes(api fiber.Router, db *gorm.DB) {
	service := services.NewProductService(db)
	validator := validator.NewProductValidator()
	handler := handlers.NewProductHandler(service, validator)
	route := api.Group("/products", middlewares.AuthMiddleware(services.NewUserService(db)))

	route.Post("/", middlewares.GetImageFile(), handler.PostProduct)
	route.Get("/", handler.GetProductsByUser)
	route.Get("/:id", handler.GetProductById)
	route.Put("/:id", middlewares.GetImageFile(), handler.PutProductById)
	route.Delete("/:id", handler.DeleteProductById)
}
