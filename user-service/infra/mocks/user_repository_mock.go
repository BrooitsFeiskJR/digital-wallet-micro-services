package mocks

import (
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/dto"
	"github.com/google/uuid"
)

type MockUserRepository struct {
	GetUserByIdFunc    func(id string) (*dto.UserDTO, error)
	UpdateUserByIdFunc func(updateData *dto.UpdateUserDTO) (*dto.UpdatedUserDTO, error)
}

func (m *MockUserRepository) GetUserById(id string) (*dto.UserDTO, error) {
	if m.GetUserByIdFunc != nil {
		return m.GetUserByIdFunc(id)
	}
	return nil, nil
}

func (m *MockUserRepository) UpdateUserById(id uuid.UUID, updateData *dto.UpdateUserDTO) (*dto.UpdatedUserDTO, error) {
	if m.UpdateUserByIdFunc != nil {
		return m.UpdateUserByIdFunc(updateData)
	}
	return nil, nil
}
