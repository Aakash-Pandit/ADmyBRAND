package main

import (
	"log"

	"github.com/Aakash-Pandit/ADmyBRAND/routes"
	"github.com/Aakash-Pandit/ADmyBRAND/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	_, err := storage.NewConnection()
	if err != nil {
		log.Fatal("could not load database")
	}

	app := fiber.New()
	app.Use("", logger.New())

	routes.SetupRoutes(app)

	app.Listen(":8080")
}
