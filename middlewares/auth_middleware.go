package middlewares

import (
	"os"
	"strings"

	"github.com/jihadable/stockwise-be/services"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(userService services.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return fiber.NewError(fiber.StatusUnauthorized, "Token tidak ditemukan")
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			return fiber.NewError(fiber.StatusUnauthorized, "Token tidak valid")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "Gagal membaca token")
		}

		userId := claims["user_id"].(string)

		_, err = userService.GetUserById(userId)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Pengguna tidak terdaftar")
		}

		ctx.Locals("user_id", userId)
		return ctx.Next()
	}
}
