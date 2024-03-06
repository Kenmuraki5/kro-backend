package interfaces

import (
	"mime/multipart"

	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
)

type UserService interface {
	GetUserByEmail(string) (*entity.User, error)
	AddUser(restmodel.User) (string, error)
	UpdateUser(restmodel.User, string) (string, error)
	AuthenticateUser(string, string) (string, error)
	UpdateProfilePicture(*multipart.FileHeader, string) (string, error)
}
