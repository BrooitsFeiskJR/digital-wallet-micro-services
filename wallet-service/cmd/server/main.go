package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/api/handlers"
	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/api/routers"
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
	walletService := services.NewWalletService(walletRepo, ctx)
	walletHandler := handlers.NewWalletHandler(walletService)

	go func() {
		if err := walletHandler.NewWalletForNewUserHandler(ctx); err != nil {
			log.Printf("Error in RabbitMQ consumer: %v", err)
		}
	}()

	router := routers.Initialize(walletHandler)

	srv := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	go func() {
		log.Println("Starting server on :8081")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down server...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
