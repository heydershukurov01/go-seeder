package configs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func LogConfig() fiber.Handler {
	return logger.New(logger.Config{
		Format:     "[${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Baku",
	})
}
