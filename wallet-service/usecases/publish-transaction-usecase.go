package usecases

import (
	"log"

	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/internal/rabbitmq"
)

type PublishTransactionUseCase struct {
	TransactionType string
	UserID          string
	WalletID        string
	ToWalletID      string
	Amount          float64
}

func NewPublishTransactionUseCaseFactory(transactionType, userID, walletID, toWalletID string, amount float64) *PublishTransactionUseCase {
	return &PublishTransactionUseCase{
		TransactionType: transactionType,
		UserID:          userID,
		WalletID:        walletID,
		ToWalletID:      toWalletID,
		Amount:          amount,
	}
}

func (p *PublishTransactionUseCase) Execute() error {
	publisher, err := rabbitmq.NewPublisher()
	if err != nil {
		return err
	}
	defer publisher.Close()

	err = publisher.PulishTransactionCreated(p.TransactionType, p.UserID, p.WalletID, p.ToWalletID, p.Amount)
	if err != nil {
		log.Printf("Failed to publish transaction created message: %v", err)
		return err
	}

	return nil
}
