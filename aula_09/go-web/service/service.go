package service

import (
	customErrors "bootcamp/aula_09/go-web/custom_errors"
	"bootcamp/aula_09/go-web/models"
	"encoding/json"
	"errors"
	"time"
)

func Hello() string {
	return "Hello Vice"
}

func GetAll() (string, error) {
	data := models.GetObjects()
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

	for _, transaction := range models.GetObjects() {
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
	if id < 0 || id >= len(models.GetObjects()) {
		return "", errors.New("invalid id")
	}

	allObjects := models.GetObjects()
	response, err := json.Marshal(allObjects[id])
	if err != nil {
		return "", err
	}
	return string(response), nil
}

func AddNewObject(data []byte) error {
	var newObject models.Transaction
	err := json.Unmarshal(data, &newObject)
	if err != nil {
		return err
	}

	// validate request fields
	if newObject.Code == "" {
		return customErrors.ErrorMissingCode
	}
	if newObject.Coin == "" {
		return customErrors.ErrorMissingCoin
	}
	if newObject.Issuer == "" {
		return customErrors.ErrorMissingIssuer
	}
	if newObject.Reciever == "" {
		return customErrors.ErrorMissingReciever
	}
	if newObject.Value == 0 {
		return customErrors.ErrorMissingValue
	}

	id := len(models.GetObjects()) + 1
	newObject.ID = id
	newObject.Date = time.Now().Format("2006-01-02")
	models.AddOject(newObject)

	return nil
}
