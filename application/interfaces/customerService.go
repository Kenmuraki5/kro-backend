package interfaces

import (
	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type CustomerService interface {
	GetUserByEmail(string) (*dynamodb.GetItemOutput, error)
	AddUser(entity.Customer) (string, error)
	UpdateUser(restmodel.Customer, string) (string, error)
	AuthenticateUser(string, string) (string, error)
}
