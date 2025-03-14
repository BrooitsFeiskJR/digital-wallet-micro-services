package handlers

import (
	"errors"
	"net/http"

	"github.com/BrooitsFeiskJR/digital-wallet-user-service/api/responses"
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/dto"
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/services"
	valueobject "github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/value_object"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service services.AuthServiceInterface
}

func NewAuthHandler(service services.AuthServiceInterface) (*AuthHandler, error) {
	if service == nil {
		return nil, errors.New("service is required")
	}
	return &AuthHandler{service}, nil
}

func (ah *AuthHandler) RegisterHandler(ctx *gin.Context) {
	var createUserDTO dto.CreateUserDTO
	if err := ctx.ShouldBindJSON(&createUserDTO); err != nil {
		response := responses.ErrorResponse(http.StatusBadRequest, err.Error())
		response.ToJSON(ctx, http.StatusBadRequest)
		return
	}

	userDTO, err := ah.service.Register(&createUserDTO)
	if err != nil {
		response := responses.ErrorResponse(http.StatusInternalServerError, err.Error())
		response.ToJSON(ctx, http.StatusInternalServerError)
		return
	}

	response := responses.SuccessResponse(userDTO, http.StatusOK)
	response.ToJSON(ctx, http.StatusOK)
}

func (ah *AuthHandler) LoginUser(ctx *gin.Context) {
	type responseToken struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
	var loginRequest valueobject.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		response := responses.ErrorResponse(http.StatusBadRequest, err.Error())
		response.ToJSON(ctx, http.StatusBadRequest)
		return
	}
	accessToken, refreshToken, err := ah.service.Login(loginRequest)
	if err != nil {
		response := responses.ErrorResponse(http.StatusBadRequest, err.Error())
		response.ToJSON(ctx, http.StatusBadRequest)
		return
	}
	tokens := responseToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	response := responses.SuccessResponse(tokens, http.StatusOK)
	response.ToJSON(ctx, http.StatusOK)
}
