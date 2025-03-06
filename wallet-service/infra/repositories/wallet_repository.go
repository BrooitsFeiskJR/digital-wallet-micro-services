package repositories

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/domain/dto"
	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/domain/entities"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type WalletRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewWalletRepository(mongoClient *mongo.Client) *WalletRepository {
	return &WalletRepository{
		client:     mongoClient,
		collection: mongoClient.Database("wallet").Collection("wallets"),
	}
}

func (wr *WalletRepository) SaveWallet(wallet *dto.CreateWalletDTO) (any, error) {
	w, err := entities.CreateWallet(wallet)
	if err != nil {
		return nil, err
	}
	result, err := wr.collection.InsertOne(context.TODO(), w)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

var ErrWalletNotFound = errors.New("wallet not found")

func (wr *WalletRepository) GetWalletByUserID(ctx context.Context, userID string) (*dto.WalletDTO, error) {
	if userID == "" {
		return nil, fmt.Errorf("empty user id")
	}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse user id: %w", err)
	}
	filter := bson.D{{Key: "user_id", Value: uid}}

	var result entities.Wallet
	err = wr.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrWalletNotFound
		}
		return nil, fmt.Errorf("failed to get wallet: %w", err)
	}

	return &dto.WalletDTO{
		ID:       result.ID,
		UserID:   result.UserID,
		Balance:  result.Balance,
		CreateAt: result.CreateAt,
		UpdateAt: result.UpdateAt,
	}, nil
}

func (wr *WalletRepository) Deposit(wallet *dto.WalletDTO, amount float64) error {
	filter := bson.D{{Key: "_id", Value: wallet.ID}}

	var result entities.Wallet
	err := wr.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return err
	}
	result.Deposit(amount)
	_, err = wr.collection.UpdateOne(context.TODO(), filter, bson.D{{Key: "$set", Value: result}})
	if err != nil {
		return err
	}
	return nil
}

func (wr *WalletRepository) Withdraw(wallet *dto.WalletDTO, amount float64) (any, error) {
	filter := bson.D{{Key: "_id", Value: wallet.ID}}

	var result entities.Wallet
	err := wr.collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	err = result.Withdraw(amount)
	if err != nil {
		return nil, err
	}
	mongoResult, err := wr.collection.UpdateOne(context.Background(), filter, bson.D{{Key: "$set", Value: result}})
	if err != nil {
		return nil, err
	}

	return mongoResult.UpsertedID, nil
}
