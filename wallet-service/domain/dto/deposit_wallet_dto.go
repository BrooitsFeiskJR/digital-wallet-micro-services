package dto

type DepostiWalletDTO struct {
	UserID string  `bson:"user_id" json:"user_id"`
	Amount float64 `bson:"amount,omitempty" json:"amount"`
}
