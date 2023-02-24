package exceptions

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx) {
	if r := recover(); r != nil {
		ctx.Status(500).JSON(fiber.Map{"error": "Error occurred: " + fmt.Sprint(r)})
	}
}
