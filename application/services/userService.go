package services

import (
	"fmt"
	"mime/multipart"

	"github.com/Kenmuraki5/kro-backend.git/application/services/auth"
	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/repository"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
	s3 "github.com/Kenmuraki5/kro-backend.git/infrastructure/persistence/s3upload"
)

type UserService struct {
	userRepository repository.UserRepository
	authService    auth.AuthService
	s3Service      s3.S3Service
}

func NewUserService(
	userRepository repository.UserRepository,
	authService auth.AuthService,
	s3Service s3.S3Service,
) *UserService {
	return &UserService{
		userRepository: userRepository,
		authService:    authService,
		s3Service:      s3Service,
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
		return nil, fmt.Errorf("failed to update customer data: %w", err)
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

func (s *UserService) UpdateProfilePicture(file *multipart.FileHeader, objKey string) (string, error) {
	uploadedFile, err := file.Open()
	if err != nil {
		return "", err
	}
	defer uploadedFile.Close()

	endpoint, err := s.s3Service.UploadFile(uploadedFile, objKey)
	if err != nil {
		return "", err
	}

	return endpoint, nil
}
