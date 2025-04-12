package services

import (
	"context"
	"errors"
	"log"

	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/domain/dto"
	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/domain/repositories"
	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/internal/rabbitmq"
)

type WalletService struct {
	repository repositories.WalletRepository
	publisher  *rabbitmq.Publisher
	ctx        context.Context
}

func NewWalletService(repository repositories.WalletRepository, rabbit *rabbitmq.Publisher, ctx context.Context) *WalletService {
	return &WalletService{
		repository: repository,
		publisher:  rabbit,
		ctx:        ctx,
	}
}

func (ws *WalletService) CreateWallet(dto *dto.CreateWalletDTO) error {
	ws.repository.SaveWallet(dto)
	return nil
}

func (ws *WalletService) GetWalletByUserID(userID string) (*dto.WalletDTO, error) {
	if userID == "" {
		return nil, errors.New("empty user id")
	}
	return ws.repository.GetWalletByUserID(ws.ctx, userID)
}

func (ws *WalletService) DepositToUserWallet(deposit *dto.DepostiWalletDTO) error {
	if deposit.Amount <= 0 {
		return errors.New("invalid deposit amount")
	}
	if deposit.UserID == "" {
		return errors.New("empty user id")
	}
	if err := ws.repository.Deposit(deposit); err != nil {
		return err
	}

	defer ws.publisher.Close()
	err := ws.publisher.PulishTransactionCreated("deposit", deposit.UserID, deposit.Amount)
	if err != nil {
		log.Printf("Failed to publish deposit created message: %v", err)
	}
	return nil
}

func (ws *WalletService) WithdrawUserWallet(withdraw *dto.WithdrawWalletDTO) error {
	if withdraw.Amount <= 0 {
		return errors.New("invalid withdraw amount")
	}
	if withdraw.UserID == "" {
		return errors.New("empty user id")
	}
	if err := ws.repository.Withdraw(withdraw); err != nil {
		return err
	}
	defer ws.publisher.Close()
	if err := ws.publisher.PulishTransactionCreated("withdraw", withdraw.UserID, withdraw.Amount); err != nil {
		log.Printf("Failed to publish withdraw created message: %v", err)
	}
	return nil
}

func (ws *WalletService) MakeTransaction(dto *dto.TransactionRequestDTO) error {
	return nil
}
