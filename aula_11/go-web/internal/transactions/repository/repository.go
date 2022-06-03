package repository

import (
	customErrors "bootcamp/aula_11/go-web/internal/custom_errors"
	"time"
)

const dateFormat = "2006-01-02 15:04:05"

type Repository interface {
	GetAll() []Transaction
	GetByID(int) (Transaction, error)
	Store(Transaction)
	Update(int, Transaction) error
}

type Transaction struct {
	ID       int     `json:"id,omitempty"`
	Code     string  `json:"transaction_code,omitempty"`
	Coin     string  `json:"coin,omitempty"`
	Value    float64 `json:"value,omitempty"`
	Issuer   string  `json:"issuer,omitempty"`
	Reciever string  `json:"reciever,omitempty"`
	Date     string  `json:"date,omitempty"`
}

type Transactions []Transaction

var repository Transactions

func (t *Transactions) GetAll() []Transaction {
	return *t
}

func (t *Transactions) GetByID(id int) (Transaction, error) {
	if id < 0 || id >= len(repository) {
		return Transaction{}, customErrors.ErrorInvalidID
	}

	return (*t)[id], nil
}

func (t *Transactions) Store(newTransaction Transaction) {
	newTransaction.ID = len(repository) + 1
	newTransaction.Date = time.Now().Format(dateFormat)
	*t = append(*t, newTransaction)
}

func (t *Transactions) Update(id int, newValues Transaction) error {
	if id < 0 || id > len(repository) {
		return customErrors.ErrorInvalidID
	}

	newValues.ID = id
	newValues.Date = time.Now().Format(dateFormat)
	(*t)[id] = newValues
	return nil
}

func NewRepository() Repository {
	return &repository
}

func Quantity() int {
	return len(repository)
}
