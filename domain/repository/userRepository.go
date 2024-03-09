package repository

import (
	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
)

type UserRepository interface {
	CreateUser(restmodel.User) (string, error)
	UpdateUser(restmodel.User, string) (string, error)
	GetUserByEmail(string) (*entity.User, error)
	AuthenticateUser(string, string) (string, error)
	GetAllUser() ([]*entity.User, error)
}
