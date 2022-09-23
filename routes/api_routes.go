package routes

import (
	"net/http"

	"github.com/Aakash-Pandit/ADmyBRAND/services"
	"github.com/gin-gonic/gin"
)

func ApiHome(context *gin.Context) {
	context.JSON(http.StatusOK, "This is Home Page of API")
}

func ApiDoc(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{})
}

func SetupRoutes(app *gin.Engine) {
	app.GET("/docs/", ApiDoc)

	api := app.Group("/api/v1")
	api.GET("", ApiHome)

	api.GET("/users", services.GetUsersHandler)
	api.GET("/users/:id", services.GetUserByIDHandler)
	api.POST("/users", services.CreateUserHandler)
	api.PATCH("/users/:id", services.UpdateUserHandler)
	api.DELETE("/users/:id", services.DeleteUserHandler)
}
