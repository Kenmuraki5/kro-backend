package repository

import (
	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type CustomerRepository interface {
	CreateUser(entity.Customer) (string, error)
	UpdateUser(restmodel.Customer, string) (string, error)
	GetUserByEmail(string) (*dynamodb.GetItemOutput, error)
	AuthenticateUser(string, string) (bool, error)
}
