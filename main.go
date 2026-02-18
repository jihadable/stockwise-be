package main

import (
	"github.com/jihadable/stockwise-be/config"
	"github.com/jihadable/stockwise-be/middlewares"
	"github.com/jihadable/stockwise-be/model/entity"
	"github.com/jihadable/stockwise-be/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.local")
	if err != nil {
		panic("Failed to read .env file: " + err.Error())
	}

	app := fiber.New()

	app.Use(cors.New(cors.ConfigDefault))

	api := app.Group("/api", middlewares.ErrorHandler())
	config := &config.Config{
		DB:    config.DB(),
		Redis: config.Redis(),
	}

	err = config.DB.AutoMigrate(&entity.User{}, &entity.Product{}, &entity.EmailVerification{})
	if err != nil {
		panic(err)
	}

	routes.RegisterUserRoutes(api, config)
	routes.RegisterProductRoutes(api, config)
	routes.RegisterEmailVerificationRoutes(api, config)

	err = app.Listen(":3000")
	if err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
