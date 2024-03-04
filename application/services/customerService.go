package services

import (
	"fmt"

	"github.com/Kenmuraki5/kro-backend.git/application/services/auth"
	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/repository"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
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

func (s *CustomerService) AddUser(user entity.Customer) (string, error) {
	id, err := s.customerRepository.CreateUser(user)
	fmt.Println(id)
	if err != nil {
		return "", err
	}
	token, err := s.authService.GenerateToken(user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *CustomerService) UpdateUser(user restmodel.Customer, email string) (string, error) {
	id, err := s.customerRepository.UpdateUser(user, email)
	if err != nil {
		return "", err
	}
	token, err := s.authService.GenerateToken(id)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *CustomerService) GetUserByEmail(email string) (*entity.Customer, error) {
	data, err := s.customerRepository.GetUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("failed to update customer data: %w", err)
	}

	return data, err
}

func (s *CustomerService) AuthenticateUser(email, password string) (string, error) {
	_, err := s.customerRepository.AuthenticateUser(email, password)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	token, err := s.authService.GenerateToken(email)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	return token, nil
}
