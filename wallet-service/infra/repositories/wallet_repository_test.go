package repositories

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/domain/dto"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupMongoTest(t *testing.T) (*mongo.Client, func()) {
	ctx := context.Background()

	mongodbContainer := testcontainers.ContainerRequest{
		Image:        "mongo:7.0",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForListeningPort("27017/tcp"),
		Cmd:          []string{"mongod", "--bind_ip_all"},
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: mongodbContainer,
		Started:          true,
	})
	require.NoError(t, err)

	mappedPort, err := container.MappedPort(ctx, "27017")
	require.NoError(t, err)

	host, err := container.Host(ctx)
	require.NoError(t, err)

	connectionString := fmt.Sprintf("mongodb://%s:%s", host, mappedPort.Port())

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	require.NoError(t, err)

	err = client.Ping(ctx, nil)
	require.NoError(t, err)

	return client, func() {
		err := client.Database("wallet").Collection("wallets").Drop(ctx)
		if err != nil {
			t.Logf("Failed to drop collection: %v", err)
		}

		if err := client.Disconnect(ctx); err != nil {
			t.Logf("Failed to disconnect client: %v", err)
		}

		if err := container.Terminate(ctx); err != nil {
			t.Logf("Failed to terminate container: %v", err)
		}
	}
}

func TestSaveWallet(t *testing.T) {
	client, cleanup := setupMongoTest(t)
	defer cleanup()

	repo := NewWalletRepository(client)
	wallet := &dto.CreateWalletDTO{
		UserID: uuid.New().String(),
	}

	id, err := repo.SaveWallet(wallet)
	require.NoError(t, err)
	require.NotNil(t, id)
}

func TestGetWalletByUserID(t *testing.T) {
	client, cleanup := setupMongoTest(t)
	defer cleanup()

	repo := NewWalletRepository(client)

	userId := uuid.New().String()
	walletData := &dto.CreateWalletDTO{
		UserID: userId,
	}

	id, err := repo.SaveWallet(walletData)
	require.NoError(t, err)
	require.NotNil(t, id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var wallet *dto.WalletDTO
	var getErr error
	maxRetries := 5

	for retries := 0; retries < maxRetries; retries++ {
		wallet, getErr = repo.GetWalletByUserID(ctx, userId)
		if getErr == nil && wallet != nil {
			break
		}
	}

	// ... existing code ...

	require.NoError(t, getErr, "Failed to get wallet by user ID after retries")
	require.NotNil(t, wallet, "Retrieved wallet should not be nil")
	assert.Equal(t, userId, wallet.UserID, "UserID should match")
	assert.Equal(t, 0.0, wallet.Balance, "Initial balance should be 0")
}

func TestDeposit(t *testing.T) {
	client, cleanup := setupMongoTest(t)
	defer cleanup()

	repo := NewWalletRepository(client)

	userId := uuid.New().String()
	walletData := &dto.CreateWalletDTO{
		UserID: userId,
	}
	id, err := repo.SaveWallet(walletData)
	require.NoError(t, err)

	objID, ok := id.(primitive.ObjectID)
	require.True(t, ok)

	wallet, err := repo.GetWalletByUserID(context.Background(), userId)
	require.NoError(t, err)

	depositAmount := 100.0
	err = repo.Deposit(wallet, depositAmount)
	assert.NoError(t, err)

	updatedWallet := &dto.WalletDTO{}
	err = client.Database("wallet").Collection("wallets").FindOne(
		context.Background(),
		bson.D{{Key: "_id", Value: objID}},
	).Decode(updatedWallet)
	require.NoError(t, err)
	assert.Equal(t, depositAmount, updatedWallet.Balance)
}

func TestWithdraw(t *testing.T) {
	client, cleanup := setupMongoTest(t)
	defer cleanup()

	repo := NewWalletRepository(client)

	userId := uuid.New().String()
	walletData := &dto.CreateWalletDTO{
		UserID: userId,
	}
	id, err := repo.SaveWallet(walletData)
	require.NoError(t, err)

	objID, ok := id.(primitive.ObjectID)
	require.True(t, ok)

	wallet, err := repo.GetWalletByUserID(context.Background(), userId)
	require.NoError(t, err)

	depositAmount := 200.0
	err = repo.Deposit(wallet, depositAmount)
	require.NoError(t, err)

	withdrawAmount := 50.0
	_, err = repo.Withdraw(wallet, withdrawAmount)
	assert.NoError(t, err)

	updatedWallet := &dto.WalletDTO{}
	err = client.Database("wallet").Collection("wallets").FindOne(
		context.Background(),
		bson.D{{Key: "_id", Value: objID}},
	).Decode(updatedWallet)
	require.NoError(t, err)
	assert.Equal(t, 150.0, updatedWallet.Balance)

	_, err = repo.Withdraw(updatedWallet, 200.0)
	assert.Error(t, err)
}

func TestSaveWalletWithInvalidDTO(t *testing.T) {
	client, cleanup := setupMongoTest(t)
	defer cleanup()

	repo := NewWalletRepository(client)
	wallet := &dto.CreateWalletDTO{}

	id, err := repo.SaveWallet(wallet)
	require.Error(t, err)
	require.Nil(t, id)
}

func TestGetWalletByUserIDWithEmptyID(t *testing.T) {
	client, cleanup := setupMongoTest(t)
	defer cleanup()

	repo := NewWalletRepository(client)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	wallet, err := repo.GetWalletByUserID(ctx, "")
	assert.EqualError(t, err, "empty user id")
	require.Error(t, err)
	require.Nil(t, wallet)
}
