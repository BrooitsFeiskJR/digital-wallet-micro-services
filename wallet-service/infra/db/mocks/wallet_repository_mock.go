package mocks

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/domain/dto"
	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/domain/entities"
	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/infra/repositories"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MockWalletRepository implements the repository interface for testing
type MockWalletRepository struct {
	wallets map[string]*entities.Wallet
	mu      sync.RWMutex

	SaveWalletFunc        func(wallet *dto.CreateWalletDTO) (any, error)
	GetWalletByUserIDFunc func(ctx context.Context, userID string) (*dto.WalletDTO, error)
	DepositFunc           func(dto *dto.DepostiWalletDTO) error
	WithdrawFunc          func(dto *dto.WithdrawWalletDTO) error
}

func NewMockWalletRepository() *MockWalletRepository {
	mock := &MockWalletRepository{
		wallets: make(map[string]*entities.Wallet),
	}

	mock.SaveWalletFunc = mock.defaultSaveWallet
	mock.GetWalletByUserIDFunc = mock.defaultGetWalletByUserID
	mock.DepositFunc = mock.defaultDeposit
	mock.WithdrawFunc = mock.defaultWithdraw

	return mock
}

func (m *MockWalletRepository) SaveWallet(wallet *dto.CreateWalletDTO) (any, error) {
	return m.SaveWalletFunc(wallet)
}

func (m *MockWalletRepository) GetWalletByUserID(ctx context.Context, userID string) (*dto.WalletDTO, error) {
	return m.GetWalletByUserIDFunc(ctx, userID)
}

func (m *MockWalletRepository) Deposit(dto *dto.DepostiWalletDTO) error {
	return m.DepositFunc(dto)
}

func (m *MockWalletRepository) Withdraw(dto *dto.WithdrawWalletDTO) error {
	return m.WithdrawFunc(dto)
}

// Default implementations

func (m *MockWalletRepository) defaultSaveWallet(wallet *dto.CreateWalletDTO) (any, error) {
	if wallet.UserID == "" {
		return nil, fmt.Errorf("empty user id")
	}

	_, err := uuid.Parse(wallet.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse user id: %w", err)
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	// Check if wallet already exists
	if _, exists := m.wallets[wallet.UserID]; exists {
		return nil, fmt.Errorf("wallet already exists for user %s", wallet.UserID)
	}

	// Create new wallet
	newWallet, err := entities.CreateWallet(wallet)
	if err != nil {
		return nil, err
	}

	// Generate a MongoDB-like ObjectID
	objID := primitive.NewObjectID()
	newWallet.ID = objID

	m.wallets[wallet.UserID] = newWallet
	return objID, nil
}

func (m *MockWalletRepository) defaultGetWalletByUserID(ctx context.Context, userID string) (*dto.WalletDTO, error) {
	if userID == "" {
		return nil, fmt.Errorf("empty user id")
	}

	_, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse user id: %w", err)
	}

	m.mu.RLock()
	defer m.mu.RUnlock()

	wallet, exists := m.wallets[userID]
	if !exists {
		return nil, repositories.ErrWalletNotFound
	}

	return &dto.WalletDTO{
		ID:       wallet.ID,
		WalletID: wallet.WalletID,
		UserID:   wallet.UserID,
		Balance:  wallet.Balance,
		CreateAt: wallet.CreateAt,
		UpdateAt: wallet.UpdateAt,
	}, nil
}

func (m *MockWalletRepository) defaultDeposit(dto *dto.DepostiWalletDTO) error {
	if dto.UserID == "" {
		return fmt.Errorf("empty user id")
	}

	_, err := uuid.Parse(dto.UserID)
	if err != nil {
		return fmt.Errorf("failed to parse user id: %w", err)
	}

	if dto.Amount <= 0 {
		return fmt.Errorf("amount must be positive")
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	wallet, exists := m.wallets[dto.UserID]
	if !exists {
		return repositories.ErrWalletNotFound
	}

	err = wallet.Deposit(dto.Amount)
	if err != nil {
		return err
	}

	wallet.UpdateAt = time.Now()
	m.wallets[dto.UserID] = wallet
	return nil
}

func (m *MockWalletRepository) defaultWithdraw(dto *dto.WithdrawWalletDTO) error {
	if dto.UserID == "" {
		return fmt.Errorf("empty user id")
	}

	_, err := uuid.Parse(dto.UserID)
	if err != nil {
		return fmt.Errorf("failed to parse user id: %w", err)
	}

	if dto.Amount <= 0 {
		return fmt.Errorf("amount must be positive")
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	wallet, exists := m.wallets[dto.UserID]
	if !exists {
		return repositories.ErrWalletNotFound
	}

	err = wallet.Withdraw(dto.Amount)
	if err != nil {
		return err
	}

	wallet.UpdateAt = time.Now()
	m.wallets[dto.UserID] = wallet
	return nil
}

// Helper methods for tests

func (m *MockWalletRepository) Reset() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.wallets = make(map[string]*entities.Wallet)
}

func (m *MockWalletRepository) GetWalletCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.wallets)
}

func (m *MockWalletRepository) SetWallet(wallet *entities.Wallet) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.wallets[wallet.UserID] = wallet
}

func (m *MockWalletRepository) SetupMockFunctions(
	saveWalletFn func(*dto.CreateWalletDTO) (any, error),
	getWalletByUserIDFn func(context.Context, string) (*dto.WalletDTO, error),
	depositFn func(*dto.DepostiWalletDTO) error,
	withdrawFn func(*dto.WithdrawWalletDTO) error,
) {
	if saveWalletFn != nil {
		m.SaveWalletFunc = saveWalletFn
	}
	if getWalletByUserIDFn != nil {
		m.GetWalletByUserIDFunc = getWalletByUserIDFn
	}
	if depositFn != nil {
		m.DepositFunc = depositFn
	}
	if withdrawFn != nil {
		m.WithdrawFunc = withdrawFn
	}
}
