package entities

import (
	"errors"
	"time"

	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/domain/dto"
	validatiion "github.com/BrooitsFeiskJR/digital-wallet-wallet-service/domain/validation"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrEmptyUserID     = errors.New("empty user id")
	ErrConvertObjectID = errors.New("error convert object id")
	ErrInsuffientFound = errors.New("insufficient found")
	ErrInvalidAmount   = errors.New("invalid amount")
)

type Wallet struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	WalletID string             `bson:"wallet_id,omitempty" json:"wallet_id"` // Change to string
	UserID   string             `bson:"user_id,omitempty" json:"user_id"`     // Change to string
	Balance  float64            `bson:"balance,omitempty" json:"balance"`
	CreateAt time.Time          `bson:"create_at,omitempty" json:"create_at"`
	UpdateAt time.Time          `bson:"update_at,omitempty" json:"update_at"`
}

func CreateWallet(dto *dto.CreateWalletDTO) (*Wallet, error) {
	if dto.UserID == "" {
		return nil, ErrEmptyUserID
	}

	_, err := uuid.Parse(dto.UserID)
	if err != nil {
		return nil, ErrConvertObjectID
	}

	return &Wallet{
		WalletID: uuid.New().String(),
		UserID:   dto.UserID,
		Balance:  0,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}, nil
}
func (w *Wallet) Deposit(amount float64) error {
	err := validatiion.ValidateAmount(amount)
	if err != nil {
		return ErrInvalidAmount
	}
	w.Balance += amount
	w.UpdateAt = time.Now()
	return nil
}

func (w *Wallet) Withdraw(amount float64) error {
	err := validatiion.ValidateAmount(amount)
	if err != nil {
		return ErrInsuffientFound
	}
	err = validatiion.ValidateSuficientBalance(w.Balance, amount)
	if err != nil {
		return err
	}
	w.Balance -= amount
	w.UpdateAt = time.Now()
	return nil
}
