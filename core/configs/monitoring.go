package configs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"os"
)

func Monitor(app *fiber.App) {
	if os.Getenv("APP_ENV") == "development" {
		app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))
	}
}
