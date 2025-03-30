package dto

type TransactionDTO struct {
	ID        string  `json:"id"`
	EventType string  `json:"event_type"`
	UserID    string  `json:"user_id"`
	Amount    float64 `json:"amount"`
	CreatedAt string  `json:"created_at"`
}
