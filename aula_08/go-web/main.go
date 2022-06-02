package main

import (
	"bootcamp/aula_08/go-web/controller"
	"bootcamp/aula_08/go-web/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.CreateDummyData()

	r := gin.Default()
	mapRouters(r)
	r.Run()
}

func mapRouters(r *gin.Engine) {
	r.GET("/hello", controller.Hello)
	r.GET("/transactions", controller.GetAll)
	r.GET("/list/:field", controller.GetFieldValue)
	r.GET("/info/:id", controller.GetObjectByID)
}
