package main

import (
	"bootcamp/aula_07/go-web/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Vice",
		})
	})
	r.GET("/transactions", service.GetAll)
	r.Run()
}
