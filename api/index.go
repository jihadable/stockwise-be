package handler

import (
	"net/http"
	"stockwise-be/database"
	"stockwise-be/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/joho/godotenv"
)

// Handler is the main entry point of the application. Think of it like the main() method
func Handler(w http.ResponseWriter, r *http.Request) {
	// This is needed to set the proper request path in `*fiber.Ctx`
	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}

// building the fiber application
func handler() http.HandlerFunc {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Failed to read .env file: " + err.Error())
	}

	app := fiber.New()

	api := app.Group("/api")
	db := database.DB()
	routes.RegisterUserRoutes(api, db)
	routes.RegisterProductRoutes(api, db)

	return adaptor.FiberApp(app)
}
