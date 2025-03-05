package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WalletDTO struct {
	ID       primitive.ObjectID `bson:"id,omitempty" json:"id"`
	UserID   primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
	Balance  float64            `bson:"balance,omitempty" json:"balance"`
	CreateAt time.Time          `bson:"create_at,omitempty" json:"create_at"`
	UpdateAt time.Time          `bson:"update_at,omitempty" json:"update_at"`
}
