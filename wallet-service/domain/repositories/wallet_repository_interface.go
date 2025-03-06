package repositories

import (
	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/domain/dto"
	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/domain/entities"
)

type WalletRepository interface {
	SaveWallet(wallet *dto.CreateWalletDTO) (any, error)
	GetWalletByUserID(userID string) (*dto.WalletDTO, error)
	Deposit(wallet *dto.WalletDTO, amount float64) (any, error)
	Withdraw(wallet *entities.Wallet, amount float64) (any, error)
}
