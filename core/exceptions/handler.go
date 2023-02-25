package exceptions

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError
	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	// Send the error message in the response
	return ctx.Status(code).JSON(fiber.Map{
		"message": fmt.Sprintf("Error: %s", err.Error()),
		"status":  code,
	})
}
