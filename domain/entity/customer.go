package entity

import (
	"github.com/Kenmuraki5/kro-backend.git/domain/valueobject"
)

type Customer struct {
	Email       string              `json:"email"`
	FullName    string              `json:"fullName"`
	PhoneNumber string              `json:"phoneNumber"`
	Address     valueobject.Address `json:"address"`
}
