package entities

import (
	"errors"
	"time"

	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/domain/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrEmptyUserID     = errors.New("empty user id")
	ErrConvertObjectID = errors.New("error convert object id")
)

type Wallet struct {
	ID       primitive.ObjectID `bson:"id,omitempty" json:"id"`
	UserID   primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
	Balance  float64            `bson:"balance,omitempty" json:"balance"`
	CreateAt time.Time          `bson:"create_at,omitempty" json:"create_at"`
	UpdateAt time.Time          `bson:"update_at,omitempty" json:"update_at"`
}

func CreateWallet(dto *dto.CreateWalletDTO) (*Wallet, error) {
	if dto.UserID == "" {
		return nil, ErrEmptyUserID
	}
	uID, err := primitive.ObjectIDFromHex(dto.UserID)
	if err != nil {
		return nil, ErrConvertObjectID
	}
	return &Wallet{
		ID:       primitive.NewObjectID(),
		UserID:   uID,
		Balance:  0,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}, nil
}
