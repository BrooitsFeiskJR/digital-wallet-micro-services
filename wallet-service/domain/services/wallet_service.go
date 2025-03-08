package services

import (
	"context"
	"errors"

	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/domain/dto"
	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/domain/repositories"
)

type WalletService struct {
	repository repositories.WalletRepository
	ctx        context.Context
}

func NewWalletService(repository repositories.WalletRepository, ctx context.Context) *WalletService {
	return &WalletService{
		repository: repository,
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
	return nil
}
