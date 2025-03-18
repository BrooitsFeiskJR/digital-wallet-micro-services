package handlers

import (
	"errors"
	"net/http"

	"github.com/BrooitsFeiskJR/digital-wallet-user-service/api/responses"
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/dto"
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserHandler struct {
	service *services.UserService
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

func NewUserHandler(service *services.UserService) (*UserHandler, error) {
	if service == nil {
		return nil, errors.New("service is required")
	}
	return &UserHandler{service}, nil
}

func (uh *UserHandler) GetProfile(ctx *gin.Context) {
	idStr, claimsResponse := GetUserIdHelper(ctx)
	if claimsResponse != nil {
		claimsResponse.ToJSON(ctx, http.StatusUnauthorized)
		return
	}
	dto, err := uh.service.GetUserById(idStr)
	if err != nil {
		response := responses.ErrorResponse(http.StatusInternalServerError, err.Error())
		response.ToJSON(ctx, http.StatusInternalServerError)
		return
	}
	response := responses.SuccessResponse(dto, http.StatusOK)
	response.ToJSON(ctx, http.StatusOK)
}

func (uh *UserHandler) UpdateProfile(ctx *gin.Context) {
	idStr, claimsResponse := GetUserIdHelper(ctx)
	if claimsResponse != nil {
		claimsResponse.ToJSON(ctx, http.StatusUnauthorized)
		return
	}
	var updateDTO dto.UpdateUserDTO
	if err := ctx.ShouldBindJSON(&updateDTO); err != nil {
		response := responses.ErrorResponse(http.StatusBadRequest, err.Error())
		response.ToJSON(ctx, http.StatusBadRequest)
		return
	}
	dto, err := uh.service.UpdateUserById(idStr, &updateDTO)
	if err != nil {
		response := responses.ErrorResponse(http.StatusInternalServerError, err.Error())
		response.ToJSON(ctx, http.StatusInternalServerError)
		return
	}
	response := responses.SuccessResponse(dto, http.StatusOK)
	response.ToJSON(ctx, http.StatusOK)
}
