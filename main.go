package main

import (
	"admin/controllers"
	"admin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectDatabase()
	router.POST("/create", controllers.CreatePerson)
	router.GET("/list", controllers.FindPerson)
	router.Run("localhost:8080")

}
