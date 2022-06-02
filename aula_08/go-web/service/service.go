package service

import (
	"bootcamp/aula_08/go-web/models"
	"encoding/json"
	"errors"
)

func Hello() string {
	return "Hello Vice"
}

func GetAll() (string, error) {
	data := models.GetDummyData()
	response, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(response), nil
}

func GetFieldValue(field string) (string, error) {
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
		return "", err
	}

	response, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(response), nil
}

func GetObjectByID(id int) (string, error) {
	if id < 0 || id >= len(models.GetDummyData()) {
		return "", errors.New("invalid id")
	}

	allObjects := models.GetDummyData()
	response, err := json.Marshal(allObjects[id])
	if err != nil {
		return "", err
	}
	return string(response), nil
}
