package mocks

import (
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/dto"
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/entities"
)

type MockAuthRepository struct {
	LoginFunc    func(email string) (*entities.User, error)
	RegisterFunc func(req *dto.CreateUserDTO) (*dto.UserDTO, error)
}

func (m *MockAuthRepository) Login(email string) (*entities.User, error) {
	if m.LoginFunc != nil {
		return m.LoginFunc(email)
	}
	return nil, nil
}

func (m *MockAuthRepository) Register(req *dto.CreateUserDTO) (*dto.UserDTO, error) {
	if m.RegisterFunc != nil {
		return m.RegisterFunc(req)
	}
	return nil, nil
}
