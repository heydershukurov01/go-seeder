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

// @title Seeder API
// @version 1.0.0
// @description Seeder application for future applications
// @contact.name Heydar Shukurov
// @contact.email heydar.shukurov@gmail.com
// @host localhost:8080
// @BasePath /
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New(configs.Server())
	app.Use(configs.LogConfig())
	app.Use(cors.New())
	app.Static("/", "./public")
	configs.Database()
	configs.Limiter(app)
	routes.Router(app)
	configs.Monitor(app)
	configs.Swagger(app)
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
