package main

import (
	"golang-crud-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.RegisterBookRoutes(router)

	router.Run(":8080")
}
