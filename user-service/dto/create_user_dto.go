package dto

type CreateUserDTO struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	CPF         string `json:"cpf"`
}
