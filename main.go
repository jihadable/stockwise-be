package main

import (
	"stockwise-be/database"
	"stockwise-be/middlewares"
	"stockwise-be/model/entity"
	"stockwise-be/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Failed to read .env file: " + err.Error())
	}

	app := fiber.New()

	api := app.Group("/api", middlewares.ErrorHandler())
	db := database.DB()

	err = db.AutoMigrate(&entity.User{}, &entity.Product{})
	if err != nil {
		panic(err)
	}

	routes.RegisterUserRoutes(api, db)
	routes.RegisterProductRoutes(api, db)

	err = app.Listen(":3000")
	if err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
