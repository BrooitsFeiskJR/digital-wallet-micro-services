package rabbitmq

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/domain/dto"
	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/domain/services"
	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/infra/repositories"
	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/internal/rabbitmq"
)

type UserPayload struct {
	EventType  string    `json:"event_type"`
	UserID     string    `json:"user_id"`
	Email      string    `json:"email"`
	Created_at time.Time `json:"created_at"`
}

// NewUserRegistrationHandler creates a handler that processes new user registration events
func NewUserRegistrationHandler(walletService services.WalletServiceInterface) rabbitmq.MessageHandler {
	return func(msg []byte) error {
		// Parse the message
		var userMsg UserPayload
		if err := json.Unmarshal(msg, &userMsg); err != nil {
			return fmt.Errorf("failed to parse user registration message: %w", err)
		}

		// Validate required fields
		if userMsg.UserID == "" {
			return errors.New("user_id is required in message data")
		}

		existingWallet, err := walletService.GetWalletByUserID(userMsg.UserID)
		if err == nil && existingWallet != nil {
			log.Printf("Wallet already exists for user %s, skipping creation", userMsg.UserID)
			return nil
		}

		if err != nil && !errors.Is(err, repositories.ErrWalletNotFound) {
			return fmt.Errorf("error checking existing wallet: %w", err)
		}

		// Create wallet DTO
		walletDTO := &dto.CreateWalletDTO{
			UserID: userMsg.UserID,
		}

		// Save wallet in database
		err = walletService.CreateWallet(walletDTO)
		if err != nil {
			return fmt.Errorf("failed to create wallet for user %s: %w", userMsg.UserID, err)
		}

		log.Printf("Created wallet for user %v", userMsg.UserID)
		return nil
	}
}

func SetupUserRegistrationConsumer(queueName string, walletService services.WalletServiceInterface) (*rabbitmq.Consumer, error) {
	consumer, err := rabbitmq.NewConsumer(queueName)
	if err != nil {
		return nil, err
	}

	// Register handler for user.registered event
	userRegistrationHandler := NewUserRegistrationHandler(walletService)
	consumer.RegisterHandler("user_created", userRegistrationHandler)

	return consumer, nil
}

// Example usage in your main.go or service init
func StartConsumerService(ctx context.Context, walletService services.WalletServiceInterface) error {
	consumer, err := SetupUserRegistrationConsumer("user_queue", walletService)
	if err != nil {
		return err
	}

	defer consumer.Close()

	if err := consumer.Start(ctx); err != nil {
		return err
	}

	log.Println("User registration consumer started successfully")

	// The consumer runs in a goroutine, so we need to keep the main thread alive
	<-ctx.Done()
	log.Println("Shutting down user registration consumer")

	return nil
}
