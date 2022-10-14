package main

import (
	"dora/internal/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	doraController := controller.NewDoraController()
	router.GET("/", doraController.Create)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}

}
