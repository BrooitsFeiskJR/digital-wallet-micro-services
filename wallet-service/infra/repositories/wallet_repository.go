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

func NewWalletRepository(client *mongo.Client) *WalletRepository {
	return &WalletRepository{
		client:     client,
		collection: client.Database("wallet").Collection("wallets"),
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
	fmt.Printf("Inserted wallet with ID: %v\n", result.InsertedID)
	return result.InsertedID, nil
}

var ErrWalletNotFound = errors.New("wallet not found")

func (wr *WalletRepository) GetWalletByUserID(ctx context.Context, userID string) (*dto.WalletDTO, error) {
	if userID == "" {
		return nil, fmt.Errorf("empty user id")
	}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse user id: %w", err)
	}

	filter := bson.D{{Key: "user_id", Value: userID}}

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
		WalletID: result.WalletID,
		Balance:  result.Balance,
		CreateAt: result.CreateAt,
		UpdateAt: result.UpdateAt,
	}, nil
}

func (wr *WalletRepository) Deposit(dto *dto.DepostiWalletDTO) error {
	fmt.Printf("Attempting deposit for user: %s, amount: %.2f\n", dto.UserID, dto.Amount)
	if dto.UserID == "" {
		return fmt.Errorf("empty user id")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := uuid.Parse(dto.UserID)
	if err != nil {
		return fmt.Errorf("failed to parse user id: %w", err)
	}

	filter := bson.D{{Key: "user_id", Value: dto.UserID}}

	session, err := wr.client.StartSession()
	if err != nil {
		return fmt.Errorf("failed to start session: %w", err)
	}
	defer session.EndSession(ctx)

	var wallet entities.Wallet
	_, err = session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		err := wr.collection.FindOne(sessCtx, filter).Decode(&wallet)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return nil, ErrWalletNotFound
			}
			return nil, fmt.Errorf("failed to get wallet: %w", err)
		}

		err = wallet.Deposit(dto.Amount)
		if err != nil {
			return nil, err
		}

		update := bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "balance", Value: wallet.Balance},
				{Key: "update_at", Value: wallet.UpdateAt},
			}},
		}

		result, err := wr.collection.UpdateOne(sessCtx, filter, update)
		if err != nil {
			return nil, fmt.Errorf("failed to update wallet: %w", err)
		}

		if result.ModifiedCount == 0 {
			return nil, errors.New("wallet was not updated")
		}

		return nil, nil
	})

	if err != nil {
		fmt.Printf("Deposit failed: %v\n", err)
		return err
	}

	fmt.Printf("Deposit successful for user %s, new balance: %.2f\n", dto.UserID, wallet.Balance)
	return nil
}

func (wr *WalletRepository) Withdraw(dto *dto.WithdrawWalletDTO) error {
	fmt.Printf("Attempting withdrawal for user: %s, amount: %.2f\n", dto.UserID, dto.Amount)
	if dto.UserID == "" {
		return fmt.Errorf("empty user id")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := uuid.Parse(dto.UserID)
	if err != nil {
		return fmt.Errorf("failed to parse user id: %w", err)
	}

	filter := bson.D{{Key: "user_id", Value: dto.UserID}}

	session, err := wr.client.StartSession()
	if err != nil {
		return fmt.Errorf("failed to start session: %w", err)
	}
	defer session.EndSession(ctx)

	var wallet entities.Wallet
	_, err = session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		err := wr.collection.FindOne(sessCtx, filter).Decode(&wallet)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return nil, ErrWalletNotFound
			}
			return nil, fmt.Errorf("failed to get wallet: %w", err)
		}

		err = wallet.Withdraw(dto.Amount)
		if err != nil {
			return nil, err
		}

		update := bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "balance", Value: wallet.Balance},
				{Key: "update_at", Value: wallet.UpdateAt},
			}},
		}

		result, err := wr.collection.UpdateOne(sessCtx, filter, update)
		if err != nil {
			return nil, fmt.Errorf("failed to update wallet: %w", err)
		}

		if result.ModifiedCount == 0 {
			return nil, errors.New("wallet was not updated")
		}

		return nil, nil
	})

	if err != nil {
		fmt.Printf("Withdrawal failed: %v\n", err)
		return err
	}

	fmt.Printf("Withdrawal successful for user %s, new balance: %.2f\n", dto.UserID, wallet.Balance)
	return nil
}
