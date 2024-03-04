package restmodel

import (
	"time"
)

type Order struct {
	ProductId       string    `json:"productId"`
	Quantity        int       `json:"auantity"`
	Email           string    `json:"email"`
	OrderDate       time.Time `json:"orderDate"`
	Subtotal        float64   `json:"subtotal"`
	ShippingAddress string    `json:"shippingAddress"`
	ShippingMethod  string    `json:"shippingMethod"`
	Type            string    `json:"type"`
}
