package api

import (
	"github.com/jihadable/stockwise-be/database"
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

	api := app.Group("/api", middlewares.ErrorHandler())
	db := database.DB()
	routes.RegisterUserRoutes(api, db)
	routes.RegisterProductRoutes(api, db)

	return app
}

var JWT string
var App = TestApp()
