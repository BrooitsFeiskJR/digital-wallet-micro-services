package repositories

import "go.mongodb.org/mongo-driver/mongo"

type WalletRepository struct {
	mongoClient *mongo.Client
}

func NewWalletRepository(mongoClient *mongo.Client) *WalletRepository {
	return &WalletRepository{mongoClient: mongoClient}
}
