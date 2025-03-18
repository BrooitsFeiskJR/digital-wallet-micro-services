package dto

type UpdateUserDTO struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
}
