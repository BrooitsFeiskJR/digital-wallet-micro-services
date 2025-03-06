package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/api/handlers"
	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/domain/services"
	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/infra/db"
	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/infra/repositories"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable is required")
	}

	mongoClient, err := db.ConnectMongoDB(mongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	walletRepo := repositories.NewWalletRepository(mongoClient)

	// Initialize services
	walletService := services.NewWalletService(walletRepo, ctx)

	// Initialize handlers
	walletHandler := handlers.NewWalletHandler(walletService)

	go walletHandler.NewWalletForNewUserHandler(ctx)

	blockUntilShutdown()
}

func blockUntilShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

}
