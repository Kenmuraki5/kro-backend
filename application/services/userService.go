package services

import (
	"fmt"

	"github.com/Kenmuraki5/kro-backend.git/application/services/auth"
	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/repository"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
)

type UserService struct {
	userRepository repository.UserRepository
	authService    auth.AuthService
}

func NewUserService(
	userRepository repository.UserRepository,
	authService auth.AuthService,
) *UserService {
	return &UserService{
		userRepository: userRepository,
		authService:    authService,
	}
}

// register
func (s *UserService) AddUser(user restmodel.User) (string, error) {
	email, err := s.userRepository.CreateUser(user)
	fmt.Println(email)
	if err != nil {
		return "", err
	}
	token, err := s.authService.GenerateToken(email, "customer")
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *UserService) UpdateUser(user restmodel.User, email string) (string, error) {
	email, err := s.userRepository.UpdateUser(user, email)
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}

	return email, nil
}

func (s *UserService) GetUserByEmail(email string) (*entity.User, error) {
	data, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("failed to get customer data: %w", err)
	}

	return data, err
}

func (s *UserService) GetAllUser() ([]*entity.User, error) {
	data, err := s.userRepository.GetAllUser()
	if err != nil {
		return nil, fmt.Errorf("failed to get user data: %w", err)
	}

	return data, err
}

func (s *UserService) AuthenticateUser(email, password string) (string, error) {
	role, err := s.userRepository.AuthenticateUser(email, password)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	fmt.Println(role)
	token, err := s.authService.GenerateToken(email, role)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	return token, nil
}
