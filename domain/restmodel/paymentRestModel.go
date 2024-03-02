package restmodel

import "time"

type Payment struct {
	Name            string     `json:"name"`
	Number          string     `json:"number"`
	ExpirationMonth time.Month `json:"expirationMonth"`
	ExpirationYear  int        `json:"expirationYear"`
	Total           int64      `json:"total"`
}
