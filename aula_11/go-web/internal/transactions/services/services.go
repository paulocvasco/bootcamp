package services

import (
	customErrors "bootcamp/aula_11/go-web/internal/custom_errors"
	"bootcamp/aula_11/go-web/internal/transactions/repository"
	"encoding/json"
	"strconv"
)

type Service interface {
	GetAll() (string, error)
	GetByID(string) (string, error)
	Store([]byte) error
	Update(string, []byte) error
}

type service struct {
	repository repository.Repository
}

func NewService(r repository.Repository) Service {
	newService := &service{
		repository: r,
	}
	return newService
}

func (s *service) GetAll() (string, error) {
	data := s.repository.GetAll()
	response, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(response), nil
}

func (s *service) GetByID(id string) (string, error) {
	index, err := strconv.Atoi(id)
	if err != nil {
		return "", customErrors.ErrorInvalidIDParameter
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

func (s *service) Store(data []byte) error {
	var newTransaction repository.Transaction
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

func (s *service) Update(id string, data []byte) error {
	value, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	var newTransaction repository.Transaction
	err = json.Unmarshal(data, &newTransaction)
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

	err = s.repository.Update(value, newTransaction)
	if err != nil {
		return err
	}

	return nil
}
