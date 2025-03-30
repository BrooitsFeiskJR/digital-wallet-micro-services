package services

import (
	"errors"

	"github.com/BrooitsFeiskJR/digital-wallet-transaction-server/domain/dto"
	"github.com/BrooitsFeiskJR/digital-wallet-transaction-server/domain/repositories"
)

type TransactionService struct {
	repository repositories.TransactionRepositoryInterface
}

func NewTransactionService(repository repositories.TransactionRepositoryInterface) (*TransactionService, error) {
	if repository == nil {
		return nil, errors.New("repository is required")
	}
	return &TransactionService{
		repository: repository,
	}, nil
}

func (ts *TransactionService) CreateTransaction(dto *dto.CreateTransactionDTO) (*dto.TransactionDTO, error) {
	if dto == nil {
		return nil, errors.New("dto is required")
	}
	transactionDTO, err := ts.repository.CreateTransaction(dto)
	if err != nil {
		return nil, err
	}
	return transactionDTO, nil
}
