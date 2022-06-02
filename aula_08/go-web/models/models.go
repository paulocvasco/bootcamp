package models

import (
	"time"
)

type Transaction struct {
	ID       int     `json:"id,omitempty"`
	Code     string  `json:"transaction_code,omitempty"`
	Coin     string  `json:"coin,omitempty"`
	Value    float64 `json:"value,omitempty"`
	Issuer   string  `json:"issuer,omitempty"`
	Reciever string  `json:"reciever,omitempty"`
	Date     string  `json:"date,omitempty"`
}

var dummyData []Transaction

func GetDummyData() []Transaction {
	return dummyData
}

func CreateDummyData() {
	newModel := Transaction{
		ID:       1,
		Code:     "xxxxxxxxxx",
		Coin:     "dolar",
		Value:    78.34,
		Issuer:   "A",
		Reciever: "B",
		Date:     time.Now().Format("2006-01-02"),
	}
	dummyData = append(dummyData, newModel)

	newModel = Transaction{
		ID:       2,
		Code:     "fffffff",
		Coin:     "dolar",
		Value:    547.89,
		Issuer:   "C",
		Reciever: "D",
		Date:     time.Now().Format("2006-01-02"),
	}
	dummyData = append(dummyData, newModel)
}
