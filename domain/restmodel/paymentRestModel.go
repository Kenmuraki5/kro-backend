package restmodel

import "time"

type Payment struct {
	Name            string     `json:"name"`
	Number          string     `json:"number"`
	ExpirationMonth time.Month `json:"expirationMonth"`
	ExpirationYear  int        `json:"expirationYear"`
	Cvc             string     `json:"cvc"`
}
