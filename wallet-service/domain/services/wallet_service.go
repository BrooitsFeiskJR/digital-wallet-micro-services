package services

import (
	"context"

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
	return ws.repository.GetWalletByUserID(ws.ctx, userID)
}
