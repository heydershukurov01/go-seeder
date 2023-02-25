package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.architecture/src/controllers/users"
)

func Users(router fiber.Router) {
	router.Get("", users.List)
	router.Post("", users.Create)
	router.Put("/:id", users.Update)
	router.Delete("/:id", users.Delete)
}
