package repositories

import (
	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/domain/dto"
	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/domain/entities"
)

type WalletRepository interface {
	SaveWallet(wallet *dto.CreateWalletDTO) error
	GetWalletByUserID(userID string) (*dto.WalletDTO, error)
	Deposit(wallet *entities.Wallet, amount float64) (*dto.WalletDTO, error)
	Withdraw(wallet *entities.Wallet, amount float64) (*dto.WalletDTO, error)
}
