package handlers

import (
	"context"
	"log"

	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/domain/services"
	rabbitmq "github.com/BrooitsFeiskJR/digital-wallet-wallet-service/internal/handlers"
)

type WalletHandler struct {
	service *services.WalletService
}

func NewWalletHandler(service *services.WalletService) *WalletHandler {
	return &WalletHandler{
		service: service,
	}
}

func (wh *WalletHandler) NewWalletForNewUserHandler(ctx context.Context) {
	if err := rabbitmq.StartConsumerService(ctx, wh.service); err != nil {
		log.Fatalf("Failed to start consumer service: %v", err)
	}
}
