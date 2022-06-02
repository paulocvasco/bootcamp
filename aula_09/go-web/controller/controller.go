package controller

import (
	"bootcamp/aula_09/go-web/service"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	response := service.Hello()
	c.JSON(200, gin.H{
		"message": response,
	})
}

func GetAll(c *gin.Context) {
	response, err := service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, string(response))
}

func GetFieldValue(c *gin.Context) {
	field, ok := c.Params.Get("field")
	if !ok {
		c.JSON(http.StatusBadRequest, errors.New("missed field parameter on URL"))
		return
	}

	response, err := service.GetFieldValue(field)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, string(response))
}

func GetObjectByID(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	response, err := service.GetObjectByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, response)
}

func AddObject(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = service.AddNewObject(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}
