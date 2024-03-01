package restmodel

import (
	"time"
)

type Order struct {
	ProductId       string
	Quantity        int
	CustomerId      string
	OrderDate       time.Time
	Subtotal        float64
	ShippingAddress string
	Type            string
}
