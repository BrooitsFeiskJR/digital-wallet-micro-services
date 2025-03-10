package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/api/responses"
	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/domain/dto"
	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/domain/services"
	rabbitmq "github.com/BrooitsFeiskJR/digital-wallet-wallet-service/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type WalletHandler struct {
	service services.WalletServiceInterface
}

func NewWalletHandler(service services.WalletServiceInterface) *WalletHandler {
	return &WalletHandler{
		service: service,
	}
}

func (wh *WalletHandler) NewWalletForNewUserHandler(ctx context.Context) error {
	return rabbitmq.StartConsumerService(ctx, wh.service)
}

func GetUserIdHelper(ctx *gin.Context) (string, *responses.APIResponse) {
	claimsInterface, exists := ctx.Get("JWT_Claims")
	if !exists {
		return "", responses.ErrorResponse(http.StatusUnauthorized, "claims not found")
	}
	claims, ok := claimsInterface.(jwt.MapClaims)
	if !ok {
		return "", responses.ErrorResponse(http.StatusInternalServerError, "invalid claims type")
	}

	idStr, ok := claims["id"].(string)
	if !ok {
		return "", responses.ErrorResponse(http.StatusInternalServerError, "invalid user ID in token")
	}
	return idStr, nil
}

func (wh *WalletHandler) GetWallet(ctx *gin.Context) {
	idStr, claimsResponse := GetUserIdHelper(ctx)
	if claimsResponse != nil {
		claimsResponse.ToJSON(ctx, http.StatusUnauthorized)
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

func (wh *WalletHandler) DepositHandler(ctx *gin.Context) {
	var request dto.DepositRequestDTO
	idStr, claimsResponse := GetUserIdHelper(ctx)
	if claimsResponse != nil {
		claimsResponse.ToJSON(ctx, http.StatusUnauthorized)
		return
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		response := responses.ErrorResponse(http.StatusBadRequest, err.Error())
		response.ToJSON(ctx, http.StatusBadRequest)
		return
	}
	dto := &dto.DepostiWalletDTO{
		UserID: idStr,
		Amount: request.Amount,
	}

	if err := wh.service.DepositToUserWallet(dto); err != nil {
		response := responses.ErrorResponse(http.StatusInternalServerError, err.Error())
		response.ToJSON(ctx, http.StatusInternalServerError)
		return
	}
	response := responses.SuccessResponse(nil, http.StatusOK)
	response.ToJSON(ctx, http.StatusOK)
}

var (
	ErrWalletNotFound = errors.New("wallet not found")
)

func (h *WalletHandler) WithdrawHandler(c *gin.Context) {
	userID, apiError := GetUserIdHelper(c)
	if apiError != nil {
		c.JSON(http.StatusUnauthorized, apiError)
		return
	}

	var withdrawRequest dto.WithdrawRequestDTO
	if err := c.ShouldBindJSON(&withdrawRequest); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(http.StatusBadRequest, "Invalid request body"))
		return
	}

	withdrawDTO := &dto.WithdrawWalletDTO{
		UserID: userID,
		Amount: withdrawRequest.Amount,
	}

	err := h.service.WithdrawUserWallet(withdrawDTO)
	if err != nil {
		fmt.Printf("Withdraw error: %v\n", err)
		statusCode := http.StatusInternalServerError

		if errors.Is(err, ErrWalletNotFound) {
			statusCode = http.StatusNotFound
		} else if strings.Contains(err.Error(), "invalid") {
			statusCode = http.StatusBadRequest
		}

		c.JSON(statusCode, responses.ErrorResponse(statusCode, err.Error()))
		return
	}

	wallet, err := h.service.GetWalletByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse(http.StatusInternalServerError, "Failed to get updated wallet"))
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse(wallet, http.StatusOK, "Withdrawal successful"))
}
