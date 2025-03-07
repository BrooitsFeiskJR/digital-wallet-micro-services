package handlers

import (
	"context"
	"net/http"

	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/api/responses"
	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/domain/services"
	rabbitmq "github.com/BrooitsFeiskJR/digital-wallet-wallet-service/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type WalletHandler struct {
	service *services.WalletService
}

func NewWalletHandler(service *services.WalletService) *WalletHandler {
	return &WalletHandler{
		service: service,
	}
}

func (wh *WalletHandler) NewWalletForNewUserHandler(ctx context.Context) error {
	return rabbitmq.StartConsumerService(ctx, wh.service)
}

func (wh *WalletHandler) GetWallet(ctx *gin.Context) {
	claimsInterface, exists := ctx.Get("JWT_Claims")
	if !exists {
		response := responses.ErrorResponse(http.StatusUnauthorized, "claims not found")
		response.ToJSON(ctx, http.StatusUnauthorized)
		return
	}
	claims, ok := claimsInterface.(jwt.MapClaims)
	if !ok {
		response := responses.ErrorResponse(http.StatusInternalServerError, "invalid claims type")
		response.ToJSON(ctx, http.StatusInternalServerError)
		return
	}

	idStr, ok := claims["id"].(string)
	if !ok {
		response := responses.ErrorResponse(http.StatusInternalServerError, "invalid user ID in token")
		response.ToJSON(ctx, http.StatusInternalServerError)
		return
	}

	dto, err := wh.service.GetWalletByUserID(idStr)
	if err != nil {
		response := responses.ErrorResponse(http.StatusInternalServerError, err.Error())
		response.ToJSON(ctx, http.StatusInternalServerError)
		return
	}
	response := responses.SuccessResponse(dto, http.StatusOK)
	response.ToJSON(ctx, http.StatusOK)
}
