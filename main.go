package main

import (
	"log"
	"net/http"

	"github.com/Aakash-Pandit/ADmyBRAND/routes"
	"github.com/Aakash-Pandit/ADmyBRAND/storage"
	"github.com/gin-gonic/gin"
	"github.com/mvrilo/go-redoc"
	ginredoc "github.com/mvrilo/go-redoc/gin"
)

func ApiDoc(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{})
}

func main() {

	doc := redoc.Redoc{
		Title:       "Example API",
		Description: "Example API Description",
		SpecFile:    "./Aakash-Pandit-API-Documentation-v1.0.0-resolved.json",
		SpecPath:    "/Aakash-Pandit-API-Documentation-v1.0.0-resolved.json",
		DocsPath:    "/docs",
	}

	_, err := storage.NewConnection()
	if err != nil {
		log.Fatal("could not load database")
	}

	route := gin.Default()
	route.Use(ginredoc.New(doc))
	route.LoadHTMLGlob("templates/*.html")
	routes.SetupRoutes(route)
	route.Run()
}
