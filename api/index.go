package handler

import (
	"fmt"
	"net/http"
	"os"
	"stockwise-be/database"
	"stockwise-be/middlewares"
	"stockwise-be/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	if os.Getenv("VERCEL_ENV") == "" {
		if err := godotenv.Load(".env"); err != nil {
			fmt.Println("Warning: .env file not found, using environment variables instead.")
		}
	}

	app := fiber.New()

	app.Use(cors.New(cors.ConfigDefault))

	api := app.Group("/api", middlewares.ErrorHandler())
	db := database.DB()
	routes.RegisterUserRoutes(api, db)
	routes.RegisterProductRoutes(api, db)

	return adaptor.FiberApp(app)
}
