package main

import (
	"fiber_intro/database"
	"fiber_intro/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	database.ConnectDB()

	// Setup the router
	router.SetupRoutes(app)

	// Listen on PORT 300
	app.Listen(":3000")
}
