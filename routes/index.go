package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Route("/api", func(api fiber.Router) {
		api.Route("/v1", func(v1 fiber.Router) {
			v1.Route("/users", Users)
		})
	})
}
