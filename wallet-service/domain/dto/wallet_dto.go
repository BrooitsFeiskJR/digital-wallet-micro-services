package dto

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WalletDTO struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID   uuid.UUID          `bson:"user_id" json:"user_id"`
	Balance  float64            `bson:"balance,omitempty" json:"balance"`
	CreateAt time.Time          `bson:"create_at,omitempty" json:"create_at"`
	UpdateAt time.Time          `bson:"update_at,omitempty" json:"update_at"`
}
