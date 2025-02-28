package entities

import (
	"errors"
	"time"

	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/dto"
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/validation"
	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Email       string    `json:"email" db:"email"`
	Password    string    `json:"password" db:"password"`
	PhoneNumber string    `json:"phone_number" db:"phone_number"`
	CPF         string    `json:"cpf" db:"cpf"`
	CreatedAt   string    `json:"created_at" db:"created_at"`
	UpdateAt    string    `json:"updated_at" db:"updated_at"`
}

func NewUser(dto *dto.CreateUserDTO) (*User, error) {
	if err := validation.ValidateUserRequiredFields(dto.Name, dto.Email, dto.Password, dto.PhoneNumber, dto.CPF); err != nil {
		return nil, err
	}
	result := validation.ValidateEmailField(dto.Email)
	if !result {
		return nil, errors.New("email is invalid")
	}
	result = validation.ValidatePasswordField(dto.Password)
	if !result {
		return nil, errors.New("password length must be greater than 8")
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
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdateAt:    time.Now().Format(time.RFC3339),
	}, nil
}
