package repositories

import (
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/dto"
	"github.com/google/uuid"
)

type UserRepositoryInterface interface {
	GetUserById(id string) (*dto.UserDTO, error)
	UpdateUserById(id uuid.UUID, updateData *dto.UpdateUserDTO) (*dto.UpdatedUserDTO, error)
}
