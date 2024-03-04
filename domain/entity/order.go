package entity

import (
	"time"
)

type Order struct {
	OrderId         string    `json:"orderId"`
	ProductId       string    `json:"productId"`
	Quantity        int       `json:"quantity"`
	Email           string    `json:"email"`
	OrderDate       time.Time `json:"orderDate"`
	Status          string    `json:"status"`
	Subtotal        float64   `json:"subtotal"`
	ShippingAddress string    `json:"shippingAddress"`
	ShippingMethod  string    `json:"shippingMethod"`
	Type            string    `json:"type"`
}
