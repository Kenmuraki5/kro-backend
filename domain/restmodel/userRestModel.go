package restmodel

import (
	"github.com/Kenmuraki5/kro-backend.git/domain/valueobject"
)

type User struct {
	Email        string              `json:"email"`
	Password     string              `json:"password"`
	FullName     string              `json:"fullName"`
	PhoneNumber  string              `json:"phoneNumber"`
	Address      valueobject.Address `json:"address"`
	ImageProfile []string            `json:"imageProfile"`
}
