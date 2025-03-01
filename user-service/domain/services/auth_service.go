package services

import (
	"errors"

	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/dto"
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/repositories"
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/validation"
	valueobject "github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/value_object"
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/infra/auth"
)

var (
	ErrNilRepository = errors.New("nil repository")
)

type AuthService struct {
	repository repositories.AuthRepositoryInterface
}

func NewAuthService(repository repositories.AuthRepositoryInterface) (*AuthService, error) {
	if repository == nil {
		return nil, ErrNilRepository
	}
	return &AuthService{repository: repository}, nil
}

func (as *AuthService) Login(req valueobject.LoginRequest) (string, error) {
	user, err := as.repository.Login(req.Email)
	if err != nil {
		return "user not found", err
	}
	if err := validation.CheckHashPassword(user.Password, req.Password); err != nil {
		return "", errors.New("invalid password")
	}
	tokenString, err := auth.CreateToken(user.ID, user.Name, user.Email)
	if err != nil {
		return "", errors.New("failed to create token")
	}
	return tokenString, nil
}

func (as *AuthService) Register(dto *dto.CreateUserDTO) (*dto.UserDTO, error) {
	if dto == nil {
		return nil, errors.New("dto is required")
	}
	return as.repository.Register(dto)
}
