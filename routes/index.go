package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.architecture/core/exceptions"
	"go.architecture/src/controllers/users"
)

func Router(app *fiber.App) {
	app.Get("users", func(ctx *fiber.Ctx) error {
		defer exceptions.ErrorHandler(ctx)
		return users.List(ctx)
	})
	app.Post("users", func(ctx *fiber.Ctx) error {
		defer exceptions.ErrorHandler(ctx)
		return users.Create(ctx)
	})
	app.Put("users/:id", func(ctx *fiber.Ctx) error {
		defer exceptions.ErrorHandler(ctx)
		return users.Update(ctx)
	})
	app.Delete("users/:id", func(ctx *fiber.Ctx) error {
		defer exceptions.ErrorHandler(ctx)
		return users.Delete(ctx)
	})
}
