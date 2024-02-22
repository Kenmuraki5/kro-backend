package entity

import (
	"github.com/Kenmuraki5/kro-backend.git/domain/valueobject"
)

type Customer struct {
	Id       string
	fullName string
	email    string
	Address  valueobject.Address
}
