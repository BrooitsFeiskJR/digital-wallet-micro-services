package services

import (
	"errors"

	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/dto"
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/repositories"
	"github.com/google/uuid"
)

var ErrNilUserRepository = errors.New("user repository is required")

type UserService struct {
	repository repositories.UserRepositoryInterface
}

func NewUserService(repository repositories.UserRepositoryInterface) (*UserService, error) {
	if repository == nil {
		return nil, ErrNilUserRepository
	}
	return &UserService{repository: repository}, nil
}

func (us *UserService) GetUserById(id string) (*dto.UserDTO, error) {
	dto, err := us.repository.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return dto, nil
}

func (us *UserService) UpdateUserById(id string, dto *dto.UpdateUserDTO) (*dto.UpdatedUserDTO, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	userId, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid user id")
	}
	return us.repository.UpdateUserById(userId, dto)
}
