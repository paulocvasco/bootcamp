package main

import (
	"bootcamp/aula_10/go-web/authentication"
	"bootcamp/aula_10/go-web/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	mapRouters(r)
	r.Run()
}

func mapRouters(r *gin.Engine) {
	r.GET("/hello", controller.Hello)
	r.GET("/token", controller.GetToken)
	r.GET("/transactions", authentication.ValidateToken(), controller.GetAll)
	r.GET("/list/:field", authentication.ValidateToken(), controller.GetFieldValue)
	r.GET("/info/:id", authentication.ValidateToken(), controller.GetObjectByID)
	r.POST("/new", authentication.ValidateToken(), controller.AddObject)
}
