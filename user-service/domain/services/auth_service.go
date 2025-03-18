package services

import (
	"errors"

	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/dto"
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/repositories"
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/validation"
	valueobject "github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/value_object"
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/infra/auth"
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/internal/rabbitmq"
)

var (
	ErrNilRepository = errors.New("user repository is required")
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

func (as *AuthService) Login(req valueobject.LoginRequest) (string, string, error) {
	user, err := as.repository.Login(req.Email)
	if err != nil {
		return "", "", err
	}
	if err := validation.CheckHashPassword(user.Password, req.Password); err != nil {
		return "", "", errors.New("invalid password")
	}
	accessToken, err := auth.CreateAccessToken(user.ID, user.Name, user.Email)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := auth.CreateAccessToken(user.ID, user.Name, user.Email)
	if err != nil {
		return "", "", errors.New("failed to create refresh token")
	}
	return accessToken, refreshToken, nil
}

func (as *AuthService) Register(dto *dto.CreateUserDTO) (*dto.UserDTO, error) {
	if dto == nil {
		return nil, errors.New("dto is required")
	}
	userDTO, err := as.repository.Register(dto)
	if err != nil {
		return nil, err
	}
	publisher, err := rabbitmq.NewPublisher()
	if err != nil {
		return nil, errors.New("failed to create publisher")
	}
	defer publisher.Close()
	err = publisher.PublishUserCreated(userDTO.ID.String(), userDTO.Email)
	if err != nil {
		return nil, errors.New("failed to publish user created")
	}
	return userDTO, nil
}
