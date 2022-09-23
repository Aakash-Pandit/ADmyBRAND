package main

import (
	"log"
	"net/http"

	"github.com/Aakash-Pandit/ADmyBRAND/routes"
	"github.com/Aakash-Pandit/ADmyBRAND/storage"
	"github.com/gin-gonic/gin"
)

func ApiDoc(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{})
}

func main() {

	// doc := redoc.Redoc{
	// 	Title:       "Example API",
	// 	Description: "Example API Description",
	// 	SpecFile:    "./openapi.json",
	// 	SpecPath:    "/openapi.json",
	// 	DocsPath:    "/docs",
	// }

	_, err := storage.NewConnection()
	if err != nil {
		log.Fatal("could not load database")
	}

	route := gin.Default()
	route.LoadHTMLGlob("templates/*.html")
	routes.SetupRoutes(route)
	route.Run()
}
