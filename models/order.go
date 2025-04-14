package models

import (
	"time"
)

type Order struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	OrderNumber   string    `json:"order_number"`
	CustomerID    uint      `json:"customer_id"`
	Customer      Customer  `json:"customer"` // <- ini penting!
	Status        string    `json:"status"`
	Currency      string    `json:"currency"`
	Country       string    `json:"country"`
	StreetAddress string    `json:"street_address"`
	City          string    `json:"city"`
	State         string    `json:"state"`
	Zip           string    `json:"zip"`
	CreatedAt     time.Time `json:"created_at"`
}
