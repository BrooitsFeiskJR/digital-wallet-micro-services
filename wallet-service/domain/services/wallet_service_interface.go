package services

import "github.com/BrooitsFeiskJR/digital-wallet-wallet-service/domain/dto"

type WalletServiceInterface interface {
	CreateWallet(dto *dto.CreateWalletDTO) error
	GetWalletByUserID(userID string) (*dto.WalletDTO, error)
	DepositToUserWallet(deposit *dto.DepostiWalletDTO)
	WithdrawUserWallet(withdraw *dto.WithdrawWalletDTO)
}
