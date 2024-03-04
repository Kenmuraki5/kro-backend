package repository

import (
	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
)

type CustomerRepository interface {
	CreateUser(entity.Customer) (string, error)
	UpdateUser(restmodel.Customer, string) (string, error)
	GetUserByEmail(string) (*entity.Customer, error)
	AuthenticateUser(string, string) (bool, error)
}
