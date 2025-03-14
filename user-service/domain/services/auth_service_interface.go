package services

import (
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/dto"
	valueobject "github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/value_object"
)

type AuthServiceInterface interface {
	Login(req valueobject.LoginRequest) (string, string, error)
	Register(dto *dto.CreateUserDTO) (*dto.UserDTO, error)
}
