package main

import (
	apihandler "findindices/api_handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	apihandler.SetupRoutes(router)

	router.Run(":8080")
}
