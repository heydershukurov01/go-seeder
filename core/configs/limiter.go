package configs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
)

func Limiter(app *fiber.App) {
	app.Use(limiter.New(limiter.Config{
		Expiration: 3 * time.Second,
		Max:        20,
	}))
}
