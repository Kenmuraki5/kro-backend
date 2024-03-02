package repository

import (
	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type CustomerRepository interface {
	CreateUser(restmodel.Customer) (string, error)
	UpdateUser(entity.Customer) (string, error)
	GetUserByEmail(string) (*dynamodb.GetItemOutput, error)
	AddToken(string, string) (string, error)
}
