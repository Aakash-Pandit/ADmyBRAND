package routes

import (
	"github.com/Aakash-Pandit/ADmyBRAND/services"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func ApiHome(context *gin.Context) {
	context.JSON(200, "This is Home Page of API")
}

func ApiDoc(context *fiber.Ctx) error {
	return context.Render("index.html", fiber.Map{})
}

func SetupRoutes(app *gin.Engine) {
	api := app.Group("/api/v1")
	api.GET("", ApiHome)

	api.GET("/users", services.GetUsers)
	api.GET("/users/:id", services.GetUserByID)
	api.POST("/users", services.CreateUser)
	api.PATCH("/users/:id", services.UpdateUser)
	api.DELETE("/users/:id", services.DeleteUser)
}
