package main

import (
	"bootcamp/aula_09/go-web/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	mapRouters(r)
	r.Run()
}

func mapRouters(r *gin.Engine) {
	r.GET("/hello", controller.Hello)
	r.GET("/transactions", controller.GetAll)
	r.GET("/list/:field", controller.GetFieldValue)
	r.GET("/info/:id", controller.GetObjectByID)
	r.POST("/new", controller.AddObject)
}
