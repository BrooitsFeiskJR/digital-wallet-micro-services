package entities

import (
	"errors"

	"github.com/BrooitsFeiskJR/digital-wallet-user-service/dto"
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/validation"
	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	PhoneNumber string    `json:"phone_number"`
	CPF         string    `json:"cpf"`
}

func NewUser(dto dto.CreateUserDTO) (*User, error) {
	if err := validation.ValidateUserRequiredFields(dto.Name, dto.Email, dto.Password, dto.PhoneNumber, dto.CPF); err != nil {
		return nil, err
	}
	result := validation.ValidateEmailField(dto.Email)
	if !result {
		return nil, errors.New("email is invalid")
	}
	hash, err := validation.HashPassword(dto.Password)
	if err != nil {
		return nil, err
	}
	result = validation.ValidPhoneNumber(dto.PhoneNumber)
	if !result {
		return nil, errors.New("phone number is invalid")
	}
	result = validation.CheckCPF(dto.CPF)
	if !result {
		return nil, errors.New("CPF is invalid")
	}
	return &User{
		ID:          uuid.New(),
		Name:        dto.Name,
		Email:       dto.Email,
		Password:    hash,
		PhoneNumber: dto.PhoneNumber,
		CPF:         dto.CPF,
	}, nil
}
