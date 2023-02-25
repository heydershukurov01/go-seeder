package configs

import (
	"github.com/gofiber/fiber/v2"
	"go.architecture/core/exceptions"
	"os"
)

func Server() fiber.Config {
	return fiber.Config{
		Prefork:           true,
		CaseSensitive:     true,
		StrictRouting:     true,
		ServerHeader:      "Fiber",
		AppName:           os.Getenv("APP_NAME"),
		EnablePrintRoutes: os.Getenv("APP_ENV") == "development",
		ErrorHandler:      exceptions.ErrorHandler,
	}
}
