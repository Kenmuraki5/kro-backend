package entity

import (
	"github.com/Kenmuraki5/kro-backend.git/domain/valueobject"
)

type Customer struct {
	Id          string              `json:"Id"`
	FullName    string              `json:"fullName"`
	Email       string              `json:"email"`
	PhoneNumber string              `json:"phoneNumber"`
	Address     valueobject.Address `json:"address"`
}
