package main

import (
	"bootcamp/aula_10/go-web/cmd/server/authentication"
	"bootcamp/aula_10/go-web/cmd/server/controller"
	"bootcamp/aula_10/go-web/internal/transactions/repository"
	"bootcamp/aula_10/go-web/internal/transactions/services"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := repository.NewRepository()
	serv := services.NewService(repo)
	controller := controller.NewController(serv)

	r := gin.Default()
	mapRouters(r, controller)
	r.Run()
}

func mapRouters(r *gin.Engine, c controller.Controller) {
	r.GET("/token", c.GetToken)
	tr := r.Group("/transactions")
	tr.POST("/", authentication.ValidateToken(), c.Store)
	tr.GET("/all", authentication.ValidateToken(), c.GetAll)
	tr.GET("/:id", authentication.ValidateToken(), c.GetByID)
}
