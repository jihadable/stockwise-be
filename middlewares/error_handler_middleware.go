package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		err := ctx.Next()

		if err != nil {
			if fiberErr, ok := err.(*fiber.Error); ok {
				return ctx.Status(fiberErr.Code).JSON(fiber.Map{
					"status":  "fail",
					"message": fiberErr.Message,
				})
			}

			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
			})
		}

		return nil
	}
}
