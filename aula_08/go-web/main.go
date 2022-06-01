package main

import (
	"bootcamp/aula_08/go-web/models"
	"bootcamp/aula_08/go-web/service"

	"github.com/gin-gonic/gin"
)

func main() {
	models.CreateDummyData()

	r := gin.Default()
	mapRouters(r)
	r.Run()
}

func mapRouters(r *gin.Engine) {
	r.GET("/hello", service.Hello)
	r.GET("/transactions", service.GetAll)
	r.GET("/list/:field", service.GetFieldValue)
}
