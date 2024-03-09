package interfaces

import (
	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
)

type UserService interface {
	GetUserByEmail(string) (*entity.User, error)
	AddUser(restmodel.User) (string, error)
	UpdateUser(restmodel.User, string) (string, error)
	AuthenticateUser(string, string) (string, error)
	GetAllUser() ([]*entity.User, error)
}
