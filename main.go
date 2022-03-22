package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Send a string back for GET calls to the endpoint "/"
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("Its nuking!")
		return err
	})

	// Listen on PORT 300
	app.Listen(":3000")
}
