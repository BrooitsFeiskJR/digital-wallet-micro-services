package dto

import "time"

type CreateTransactionDTO struct {
	EventType string    `json:"event_type"`
	UserID    string    `json:"user_id"`
	WalletID  string    `json:"wallet_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
