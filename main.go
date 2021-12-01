package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/syncodd/simple-todoapp-fiber-go/config"
	"github.com/syncodd/simple-todoapp-fiber-go/routes"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "endpoint",
		})
	})

	api := app.Group("/api")

	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "api endpoint",
		})
	})

	routes.TodoRoute(api.Group("/todos"))
}

func main() {

	app := fiber.New()
	app.Use(logger.New())

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	config.ConnectDatabase()

	setupRoutes(app)

	err = app.Listen(":8000")

	if err != nil {
		panic(err)
	}

}
