package transactions

import (
	customErrors "bootcamp/aula_10/go-web/internal/custom_errors"
	transactions "bootcamp/aula_10/go-web/internal/transactions/repository"
	"encoding/json"
	"strconv"
)

type ServiceI interface {
	GetAll() (string, error)
	GetByID(string) (string, error)
	Store([]byte) error
}

type Service struct {
	repository transactions.Repository
}

var service Service

func (s *Service) GetAll() (string, error) {
	data := s.repository.GetAll()
	response, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(response), nil
}

func (s *Service) GetByID(id string) (string, error) {
	index, err := strconv.Atoi(id)
	if err != nil {
		return "", customErrors.WrapError(customErrors.ErrorInvalidParameter, err)
	}
	data, err := s.repository.GetByID(index)
	if err != nil {
		return "", err
	}
	response, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(response), nil
}

func (s *Service) Store(data []byte) error {
	var newTransaction transactions.Transaction
	err := json.Unmarshal(data, &newTransaction)
	if err != nil {
		return err
	}

	// validate request fields
	if newTransaction.Code == "" {
		return customErrors.ErrorMissingCode
	}
	if newTransaction.Coin == "" {
		return customErrors.ErrorMissingCoin
	}
	if newTransaction.Issuer == "" {
		return customErrors.ErrorMissingIssuer
	}
	if newTransaction.Reciever == "" {
		return customErrors.ErrorMissingReciever
	}
	if newTransaction.Value == 0 {
		return customErrors.ErrorMissingValue
	}

	s.repository.Store(newTransaction)

	return nil
}
