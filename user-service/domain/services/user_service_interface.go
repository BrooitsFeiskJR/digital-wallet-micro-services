package services

import "github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/dto"

type UserServiceInterface interface {
	GetUserById(id string) (*dto.UserDTO, error)
	UpdateUserById(id string, dto *dto.UpdateUserDTO) (*dto.UpdatedUserDTO, error)
}
