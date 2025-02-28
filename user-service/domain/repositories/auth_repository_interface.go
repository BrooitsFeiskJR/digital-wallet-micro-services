package repositories

import (
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/dto"
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/entities"
)

type AuthRepositoryInterface interface {
	Login(email string) (*entities.User, error)
	Register(req *dto.CreateUserDTO) (*dto.UserDTO, error)
}
