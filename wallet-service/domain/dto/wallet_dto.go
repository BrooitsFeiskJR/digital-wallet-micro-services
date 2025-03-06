package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WalletDTO struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID   string             `bson:"user_id" json:"user_id"`     // Change to string
	WalletID string             `bson:"wallet_id" json:"wallet_id"` // Change to string
	Balance  float64            `bson:"balance,omitempty" json:"balance"`
	CreateAt time.Time          `bson:"create_at,omitempty" json:"create_at"`
	UpdateAt time.Time          `bson:"update_at,omitempty" json:"update_at"`
}
