package middlewares

import (
	"path/filepath"
	"stockwise-be/model/request"

	"github.com/gofiber/fiber/v2"
)

func GetImageFile() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		imageHeader, err := ctx.FormFile("image")
		if err != nil {
			ctx.Locals("image", request.ImageRequest{})
			return ctx.Next()
		}

		imageFile, err := imageHeader.Open()
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Gagal menambahkan produk")
		}
		defer imageFile.Close()

		imageExt := filepath.Ext(imageHeader.Filename)

		image := request.ImageRequest{
			File: imageFile,
			Ext:  imageExt,
		}

		ctx.Locals("image", image)
		return ctx.Next()
	}
}
