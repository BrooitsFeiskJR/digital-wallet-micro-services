package config

import (
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/api/handlers"
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/services"
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/infra/repositories"
	"github.com/jmoiron/sqlx"
)

func AuthHandlerSetup(database *sqlx.DB) *handlers.AuthHandler {
	authRepository, err := repositories.NewAuthRepository(database)
	if err != nil {
		panic(err)
	}
	authService, err := services.NewAuthService(authRepository)
	if err != nil {
		panic(err)
	}

	authHandler, err := handlers.NewAuthHandler(authService)
	if err != nil {
		panic(err)
	}
	return authHandler
}
