package service

import (
	"bootcamp/aula_07/go-web/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	response, err := createModels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, string(response))
}

func createModels() ([]byte, error) {
	var values []models.Transaction

	newModel := models.Transaction{
		ID:       "3445435",
		Code:     "xxxxxxxxxx",
		Coin:     "dolar",
		Value:    78.34,
		Issuer:   "A",
		Reciever: "B",
		Date:     time.Now(),
	}
	values = append(values, newModel)

	newModel = models.Transaction{
		ID:       "1111111111",
		Code:     "fffffff",
		Coin:     "dolar",
		Value:    547.89,
		Issuer:   "C",
		Reciever: "D",
		Date:     time.Now(),
	}
	values = append(values, newModel)

	data, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}

	return data, nil
}
