package interfaces

import (
	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
)

type CustomerService interface {
	GetUserByEmail(string) (*entity.Customer, error)
	AddUser(entity.Customer) (string, error)
	UpdateUser(restmodel.Customer, string) (string, error)
	AuthenticateUser(string, string) (string, error)
}
