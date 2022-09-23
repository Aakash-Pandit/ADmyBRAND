package main

import (
	"log"

	"github.com/Aakash-Pandit/ADmyBRAND/routes"
	"github.com/Aakash-Pandit/ADmyBRAND/storage"
	"github.com/gin-gonic/gin"
)

func main() {

	_, err := storage.NewConnection()
	if err != nil {
		log.Fatal("could not load database")
	}

	route := gin.Default()
	routes.SetupRoutes(route)
	route.Run()
}
