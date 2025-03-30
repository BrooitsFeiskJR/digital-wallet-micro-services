package services

import "github.com/BrooitsFeiskJR/digital-wallet-wallet-service/domain/dto"

type WalletServiceInterface interface {
	CreateWallet(dto *dto.CreateWalletDTO) error
	GetWalletByUserID(userID string) (*dto.WalletDTO, error)
	DepositToUserWallet(deposit *dto.DepostiWalletDTO) error
	WithdrawUserWallet(withdraw *dto.WithdrawWalletDTO) error
	MakeTransaction(dto *dto.TransactionRequestDTO) error
}
