package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.architecture/core/configs"
	"go.architecture/routes"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New(configs.Server())
	app.Use(cors.New())
	app.Static("/", "./public")
	configs.Database()
	configs.Limiter(app)
	routes.Router(app)
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
