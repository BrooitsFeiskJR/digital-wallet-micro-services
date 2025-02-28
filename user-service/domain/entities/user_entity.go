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
	if valid, err := validateDTO(dto); !valid {
		return nil, err
	}
	hash, err := validation.HashPassword(dto.Password)
	if err != nil {
		return nil, err
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

func validateDTO(dto *dto.CreateUserDTO) (bool, error) {
	if err := validation.ValidateUserRequiredFields(dto.Name, dto.Email, dto.Password, dto.PhoneNumber, dto.CPF); err != nil {
		return false, err
	}
	if !validation.ValidateEmailField(dto.Email) {
		return false, errors.New("email is invalid")
	}
	if !validation.ValidatePasswordField(dto.Password) {
		return false, errors.New("password length must be greater than 8")
	}
	if !validation.ValidatePasswordAndConfirmPassword(dto.Password, dto.ConfirmPassword) {
		return false, errors.New("password and confirm password must be the same")
	}
	if !validation.ValidPhoneNumber(dto.PhoneNumber) {
		return false, errors.New("phone number is invalid")
	}
	if !validation.CheckCPF(dto.CPF) {
		return false, errors.New("CPF is invalid")
	}
	return true, nil
}
