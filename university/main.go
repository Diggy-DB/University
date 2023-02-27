package main

import (
	"university/database"
	"university/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Database
	database.ConnectDB()
	database.Migrate()

	// Initiate Fiber
	app := fiber.New()
	route.AccessRoute(app)
	app.Listen(":3000")
}
