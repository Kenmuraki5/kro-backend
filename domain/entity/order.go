package entity

import (
	"time"
)

type Order struct {
	OrderId         string    `json:"orderId"`
	ProductId       string    `json:"productId"`
	Quantity        int       `json:"quantity"`
	CustomerId      string    `json:"customerId"`
	OrderDate       time.Time `json:"orderDate"`
	Status          string    `json:"status"`
	Subtotal        float64   `json:"subtotal"`
	ShippingAddress string    `json:"shippingAddress"`
	Type            string    `json:"type"`
}
