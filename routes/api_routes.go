package routes

import (
	"github.com/Aakash-Pandit/ADmyBRAND/services"
	"github.com/gofiber/fiber/v2"
)

func ApiHome(context *fiber.Ctx) error {
	return context.SendString("This is API Home Page")
}

func ApiDoc(context *fiber.Ctx) error {
	return context.Render("index.html", fiber.Map{})
}

func SetupRoutes(app *fiber.App) {
	app.Get("/docs", ApiDoc)

	api := app.Group("/api/v1")
	api.Get("", ApiHome)

	api.Get("/users", services.GetUsers)
	api.Get("/users/:id", services.GetUserByID)
	api.Post("/users", services.CreateUser)
	api.Patch("/users/:id", services.UpdateUser)
	api.Delete("/users/:id", services.DeleteUser)
}
