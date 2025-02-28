package dto

import "github.com/google/uuid"

type UserDTO struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number" db:"phone_number"`
	CPF         string    `json:"cpf"`
	CreatedAt   string    `json:"created_at" db:"created_at"`
}
