package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.architecture/core/exceptions"
	"go.architecture/src/controllers/users"
)

func Router(app *fiber.App) {
	app.Get("", func(ctx *fiber.Ctx) error {
		defer exceptions.ErrorHandler(ctx)
		return users.List(ctx)
	})
	app.Post("", func(ctx *fiber.Ctx) error {
		defer exceptions.ErrorHandler(ctx)
		return users.Create(ctx)
	})
}
