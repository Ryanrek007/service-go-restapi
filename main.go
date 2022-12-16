package main

import (
	"service-go-restapi/controllers"
	"service-go-restapi/models"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/endpoint", controllers.Index)

	router.Run()

}
