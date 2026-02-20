package api

import (
	"github.com/jihadable/stockwise-be/config"
	"github.com/jihadable/stockwise-be/middlewares"
	"github.com/jihadable/stockwise-be/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func TestApp() *fiber.App {
	err := godotenv.Load("../../.env.local")
	if err != nil {
		panic("Failed to read .env file: " + err.Error())
	}

	app := fiber.New()

	app.Get("/assets/logo.png", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("assets/logo.png")
	})

	api := app.Group("/api", middlewares.ErrorHandler())
	config := &config.Config{
		DB:    config.DB(),
		Redis: config.Redis(),
	}
	routes.RegisterUserRoutes(api, config)
	routes.RegisterProductRoutes(api, config)
	routes.RegisterEmailVerificationRoutes(api, config)

	return app
}

var JWT string
var App = TestApp()
