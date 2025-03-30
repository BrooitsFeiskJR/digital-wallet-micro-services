package repositories

import "github.com/BrooitsFeiskJR/digital-wallet-transaction-server/domain/dto"

type TransactionRepositoryInterface interface {
	CreateTransaction(dto *dto.CreateTransactionDTO) (*dto.TransactionDTO, error)
}
