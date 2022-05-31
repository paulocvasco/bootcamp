package models

import "time"

type Transaction struct {
	ID       string    `json:"id"`
	Code     string    `json:"transaction_code"`
	Coin     string    `json:"coin"`
	Value    float64   `json:"value"`
	Issuer   string    `json:"issuer"`
	Reciever string    `json:"reciever"`
	Date     time.Time `json:"date"`
}
