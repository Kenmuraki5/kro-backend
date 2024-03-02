package services

import (
	"fmt"

	"github.com/Kenmuraki5/kro-backend.git/application/services/auth"
	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/repository"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type CustomerService struct {
	customerRepository repository.CustomerRepository
	authService        auth.AuthService
}

func NewCustomerService(
	customerRepository repository.CustomerRepository,
	authService auth.AuthService,
) *CustomerService {
	return &CustomerService{
		customerRepository: customerRepository,
		authService:        authService,
	}
}

func (s *CustomerService) AddUser(user restmodel.Customer) (string, error) {
	id, err := s.customerRepository.CreateUser(user)
	if err != nil {
		return "", err
	}
	token, err := s.authService.GenerateToken(user.Email)
	if err != nil {
		return "", err
	}

	addtokenToDb, err := s.customerRepository.AddToken(id, token)
	if err != nil {
		return "", err
	}

	return addtokenToDb, nil
}

func (s *CustomerService) UpdateUser(user entity.Customer) (string, error) {
	id, err := s.customerRepository.UpdateUser(user)
	if err != nil {
		return "", err
	}
	token, err := s.authService.GenerateToken(id)
	if err != nil {
		return "", err
	}

	addtokenToDb, err := s.customerRepository.AddToken(id, token)
	if err != nil {
		return "", err
	}

	return addtokenToDb, nil
}

func (s *CustomerService) GetUserByEmail(email string) (*dynamodb.GetItemOutput, error) {
	data, err := s.customerRepository.GetUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("failed to update customer data: %w", err)
	}

	return data, err
}
