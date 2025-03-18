package config

import (
	"errors"

	"github.com/BrooitsFeiskJR/digital-wallet-user-service/api/handlers"
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/services"
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/infra/repositories"
	"github.com/jmoiron/sqlx"
)

func UserHandlerSetup(database *sqlx.DB) (*handlers.UserHandler, error) {
	if database == nil {
		return nil, errors.New("database is required")
	}
	userRepository, err := repositories.NewUserRepository(database)
	if err != nil {
		return nil, err
	}
	userService, err := services.NewUserService(userRepository)
	if err != nil {
		return nil, err
	}

	authHandler, err := handlers.NewUserHandler(userService)
	if err != nil {
		return nil, err
	}
	return authHandler, nil
}
