package repositories

import (
	"context"

	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/domain/dto"
)

type WalletRepository interface {
	SaveWallet(wallet *dto.CreateWalletDTO) (any, error)
	GetWalletByUserID(ctx context.Context, userID string) (*dto.WalletDTO, error)
	Deposit(dto *dto.DepostiWalletDTO) error
	Withdraw(dto *dto.WithdrawWalletDTO) error
}
