package interfaces

import (
	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type CustomerService interface {
	GetUserById(string) (*dynamodb.GetItemOutput, error)
	AddUser(restmodel.Customer) (string, error)
	UpdateUser(entity.Customer) (string, error)
}
