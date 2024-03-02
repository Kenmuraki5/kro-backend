package restmodel

import (
	"time"
)

type Order struct {
	ProductId       string    `json:"ProductId"`
	Quantity        int       `json:"Quantity"`
	CustomerId      string    `json:"CustomerId"`
	OrderDate       time.Time `json:"OrderDate"`
	Subtotal        float64   `json:"Subtotal"`
	ShippingAddress string    `json:"ShippingAddress"`
	ShippingMethod  string    `json:"ShippingMethod"`
	Type            string    `json:"Type"`
}
