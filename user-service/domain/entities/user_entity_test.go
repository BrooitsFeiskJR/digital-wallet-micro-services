package entities

import (
	"testing"

	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/dto"
	"github.com/stretchr/testify/assert"
)

func TestUserConstructor(t *testing.T) {
	dto := dto.CreateUserDTO{
		Name:            "example",
		Email:           "example@example.com",
		Password:        "password",
		ConfirmPassword: "password",
		PhoneNumber:     "43999999999",
		CPF:             "12345678909",
	}
	user, err := NewUser(&dto)
	if err != nil {
		t.Errorf("Error creating user: %v", err)
	}
	if user == nil {
		t.Errorf("Expected user, got nil")
		return
	}
	assert.Equal(t, dto.Name, user.Name)
	assert.Equal(t, dto.Email, user.Email)
	assert.Equal(t, dto.PhoneNumber, user.PhoneNumber)
	assert.Equal(t, dto.CPF, user.CPF)
}

func TestUserConstructorInvalidPhoneNumber(t *testing.T) {
	dto := dto.CreateUserDTO{
		Name:            "example",
		Email:           "example@example.com",
		Password:        "password",
		ConfirmPassword: "password",
		PhoneNumber:     "invalid-phone-number",
		CPF:             "12345678909",
	}
	user, err := NewUser(&dto)
	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func TestUserConstructorInvalidEmail(t *testing.T) {
	dto := dto.CreateUserDTO{
		Name:            "example",
		Email:           "invalid-email",
		Password:        "password",
		ConfirmPassword: "password",
		PhoneNumber:     "43999999999",
		CPF:             "12345678909",
	}
	user, err := NewUser(&dto)
	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func TestUserConstructorInvalidPassword(t *testing.T) {
	dto := dto.CreateUserDTO{
		Name:            "example",
		Email:           "example@example.com",
		Password:        "short",
		ConfirmPassword: "short",
		PhoneNumber:     "43999999999",
		CPF:             "12345678909",
	}
	user, err := NewUser(&dto)
	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func TestUserConstructorInvalidCPF(t *testing.T) {
	dto := dto.CreateUserDTO{
		Name:            "example",
		Email:           "example@example.com",
		Password:        "password",
		ConfirmPassword: "password",
		PhoneNumber:     "43999999999",
		CPF:             "invalid-cpf",
	}
	user, err := NewUser(&dto)
	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func TestUserConstructorMissingFields(t *testing.T) {
	dto := dto.CreateUserDTO{
		Name:            "",
		Email:           "",
		Password:        "",
		ConfirmPassword: "",
		PhoneNumber:     "",
		CPF:             "",
	}
	user, err := NewUser(&dto)
	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func TestCreatedAndUpteAtFields(t *testing.T) {
	dto := dto.CreateUserDTO{
		Name:            "example",
		Email:           "example@example.com",
		Password:        "password",
		ConfirmPassword: "password",
		PhoneNumber:     "43999999999",
		CPF:             "12345678909",
	}
	user, err := NewUser(&dto)
	if err != nil {
		t.Errorf("Error creating user: %v", err)
	}
	if user == nil {
		t.Errorf("Expected user, got nil")
		return
	}
	assert.NotNil(t, user.CreatedAt)
	assert.NotNil(t, user.UpdateAt)
}

func ValidatePasswordAndConfirmPassword(t *testing.T) {
	dto := dto.CreateUserDTO{
		Name:            "example",
		Email:           "example@example.com",
		Password:        "password",
		ConfirmPassword: "password",
		PhoneNumber:     "43999999999",
		CPF:             "12345678909",
	}
	user, err := NewUser(&dto)
	if err != nil {
		t.Errorf("Error creating user: %v", err)
	}
	if user == nil {
		t.Errorf("Expected user, got nil")
		return
	}
	assert.Equal(t, dto.Password, user.Password)
}

func TestUserConstructorPasswordAndConfirmPasswordMismatch(t *testing.T) {
	dto := dto.CreateUserDTO{
		Name:            "example",
		Email:           "example@example.com",
		Password:        "password",
		ConfirmPassword: "password-mismatch",
		PhoneNumber:     "43999999999",
		CPF:             "12345678909",
	}
	user, err := NewUser(&dto)
	assert.NotNil(t, err)
	assert.Nil(t, user)
}
