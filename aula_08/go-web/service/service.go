package service

import (
	"bootcamp/aula_08/go-web/models"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Vice",
	})
}

func GetAll(c *gin.Context) {
	data := models.GetDummyData()
	response, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, string(response))
}

func GetFieldValue(c *gin.Context) {
	field, ok := c.Params.Get("field")
	if !ok {
		c.JSON(http.StatusBadRequest, errors.New("missed field parameter on URL"))
	}

	var result []models.Transaction
	var data models.Transaction
	var err error

	for _, transaction := range models.GetDummyData() {
		switch field {
		case "id":
			data.ID = transaction.ID
			result = append(result, data)
		case "transaction_code":
			data.Code = transaction.Code
			result = append(result, data)
		case "coin":
			data.Coin = transaction.Coin
			result = append(result, data)
		case "value":
			data.Value = transaction.Value
			result = append(result, data)
		case "issuer":
			data.Issuer = transaction.Issuer
			result = append(result, data)
		case "reciever":
			data.Reciever = transaction.Reciever
			result = append(result, data)
		default:
			err = errors.New("invalid field")
		}
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	response, err := json.Marshal(result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, string(response))
}
