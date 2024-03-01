package entity

import (
	"time"
)

type Order struct {
	OrderId         string
	ProductId       string
	Quantity        int
	CustomerId      string
	OrderDate       time.Time
	Status          string
	Subtotal        float64
	ShippingAddress string
	Type            string
}
