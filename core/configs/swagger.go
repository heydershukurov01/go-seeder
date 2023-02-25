package configs

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	_ "go.architecture/docs"
)

func Swagger(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))
}
