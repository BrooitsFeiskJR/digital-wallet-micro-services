package dto

type TransactionRequestDTO struct {
	TransactionType string  `json:"transaction_type" validate:"required"`
	WalletFrom      string  `json:"wallet_from" validate:"required"`
	WalletTo        string  `json:"wallet_to" validate:"required"`
	Amount          float64 `json:"amount" validate:"required"`
}
