package entities

import (
	"time"

	validatiion "github.com/BrooitsFeiskJR/digital-wallet-transaction-server/domain/validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	Deposit    = "deposit"
	Withdrawal = "withdrawal"
	Transfer   = "transfer"
)

type Transaction struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	SendWalletID primitive.ObjectID `bson:"send_wallet_id" json:"send_wallet_id"`
	ToWalletID   primitive.ObjectID `bson:"to_wallet_id" json:"to_wallet_id"`
	Amount       float64            `bson:"amount" json:"amount"`
	Type         string             `bson:"type" json:"type"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
}

func CreateTransaction(fromWallet, toWallet primitive.ObjectID, amount float64, tType string) (*Transaction, error) {
	if err := validatiion.ValidateType(tType); err != nil {
		return nil, err
	}
	return &Transaction{
		ID:           primitive.NewObjectID(),
		SendWalletID: fromWallet,
		ToWalletID:   toWallet,
		Amount:       amount,
		Type:         tType,
		CreatedAt:    time.Now(),
	}, nil
}
