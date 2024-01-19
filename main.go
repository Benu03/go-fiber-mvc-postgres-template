package main

import (
	"github.com/Benu03/go-fiber-mvc-postgres-template/databases"
	"github.com/Benu03/go-fiber-mvc-postgres-template/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	databases.ConnectDB()

	router.SetupRoutes(app)

	app.Listen(":8000")
}
