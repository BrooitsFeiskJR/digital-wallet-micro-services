package services

import (
	"context"
	"errors"
	"testing"

	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/domain/dto"
	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/infra/db/mocks"
	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/infra/repositories"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestNewWalletService(t *testing.T) {
	mockRepo := mocks.NewMockWalletRepository()
	ctx := context.Background()
	service := NewWalletService(mockRepo, ctx)

	assert.NotNil(t, service)
	assert.Same(t, mockRepo, service.repository)
	assert.Equal(t, ctx, service.ctx)
}

func TestCreateWallet(t *testing.T) {
	mockRepo := mocks.NewMockWalletRepository()
	ctx := context.Background()
	service := NewWalletService(mockRepo, ctx)

	tests := []struct {
		name      string
		dto       *dto.CreateWalletDTO
		setupMock func()
		wantErr   bool
		errMsg    string
	}{
		{
			name: "Success",
			dto: &dto.CreateWalletDTO{
				UserID: uuid.New().String(),
			},
			setupMock: func() {
				mockRepo.SaveWalletFunc = func(wallet *dto.CreateWalletDTO) (any, error) {
					return primitive.NewObjectID(), nil
				}
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset mock before each test
			mockRepo.Reset()
			if tt.setupMock != nil {
				tt.setupMock()
			}

			err := service.CreateWallet(tt.dto)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetWalletByUserID(t *testing.T) {
	mockRepo := mocks.NewMockWalletRepository()
	ctx := context.Background()
	service := NewWalletService(mockRepo, ctx)

	validUserID := uuid.New().String()
	expectedWallet := &dto.WalletDTO{
		WalletID: uuid.New().String(),
		UserID:   validUserID,
		Balance:  100.0,
	}

	tests := []struct {
		name        string
		userID      string
		setupMock   func()
		wantErr     bool
		errMsg      string
		checkWallet func(*testing.T, *dto.WalletDTO)
	}{
		{
			name:   "Success",
			userID: validUserID,
			setupMock: func() {
				mockRepo.GetWalletByUserIDFunc = func(ctx context.Context, userID string) (*dto.WalletDTO, error) {
					if userID == validUserID {
						return expectedWallet, nil
					}
					return nil, repositories.ErrWalletNotFound
				}
			},
			wantErr: false,
			checkWallet: func(t *testing.T, wallet *dto.WalletDTO) {
				assert.Equal(t, expectedWallet, wallet)
			},
		},
		{
			name:    "Empty UserID",
			userID:  "",
			wantErr: true,
			errMsg:  "empty user id",
		},
		{
			name:   "Wallet Not Found",
			userID: uuid.New().String(),
			setupMock: func() {
				mockRepo.GetWalletByUserIDFunc = func(ctx context.Context, userID string) (*dto.WalletDTO, error) {
					return nil, repositories.ErrWalletNotFound
				}
			},
			wantErr: true,
			errMsg:  "wallet not found",
		},
		{
			name:   "Invalid UUID",
			userID: "invalid-uuid",
			setupMock: func() {
				mockRepo.GetWalletByUserIDFunc = func(ctx context.Context, userID string) (*dto.WalletDTO, error) {
					return nil, errors.New("failed to parse user id")
				}
			},
			wantErr: true,
			errMsg:  "failed to parse user id",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset mock before each test
			mockRepo.Reset()
			if tt.setupMock != nil {
				tt.setupMock()
			}

			wallet, err := service.GetWalletByUserID(tt.userID)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
				assert.Nil(t, wallet)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, wallet)
				if tt.checkWallet != nil {
					tt.checkWallet(t, wallet)
				}
			}
		})
	}
}

func TestDepositToUserWallet(t *testing.T) {
	mockRepo := mocks.NewMockWalletRepository()
	ctx := context.Background()
	service := NewWalletService(mockRepo, ctx)

	validUserID := uuid.New().String()
	validAmount := 100.0

	tests := []struct {
		name      string
		deposit   *dto.DepostiWalletDTO
		setupMock func()
		wantErr   bool
		errMsg    string
	}{
		{
			name: "Success",
			deposit: &dto.DepostiWalletDTO{
				UserID: validUserID,
				Amount: validAmount,
			},
			setupMock: func() {
				mockRepo.DepositFunc = func(dto *dto.DepostiWalletDTO) error {
					return nil
				}
			},
			wantErr: false,
		},
		{
			name: "Empty UserID",
			deposit: &dto.DepostiWalletDTO{
				UserID: "",
				Amount: validAmount,
			},
			wantErr: true,
			errMsg:  "empty user id",
		},
		{
			name: "Invalid Amount",
			deposit: &dto.DepostiWalletDTO{
				UserID: validUserID,
				Amount: -10.0,
			},
			wantErr: true,
			errMsg:  "invalid deposit amount",
		},
		{
			name: "Zero Amount",
			deposit: &dto.DepostiWalletDTO{
				UserID: validUserID,
				Amount: 0.0,
			},
			wantErr: true,
			errMsg:  "invalid deposit amount",
		},
		{
			name: "Wallet Not Found",
			deposit: &dto.DepostiWalletDTO{
				UserID: validUserID,
				Amount: validAmount,
			},
			setupMock: func() {
				mockRepo.DepositFunc = func(dto *dto.DepostiWalletDTO) error {
					return repositories.ErrWalletNotFound
				}
			},
			wantErr: true,
			errMsg:  "wallet not found",
		},
		{
			name: "Repository Error",
			deposit: &dto.DepostiWalletDTO{
				UserID: validUserID,
				Amount: validAmount,
			},
			setupMock: func() {
				mockRepo.DepositFunc = func(dto *dto.DepostiWalletDTO) error {
					return errors.New("database error")
				}
			},
			wantErr: true,
			errMsg:  "database error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset mock before each test
			mockRepo.Reset()
			if tt.setupMock != nil {
				tt.setupMock()
			}

			err := service.DepositToUserWallet(tt.deposit)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestWithdrawUserWallet(t *testing.T) {
	mockRepo := mocks.NewMockWalletRepository()
	ctx := context.Background()
	service := NewWalletService(mockRepo, ctx)

	validUserID := uuid.New().String()
	validAmount := 50.0

	tests := []struct {
		name      string
		withdraw  *dto.WithdrawWalletDTO
		setupMock func()
		wantErr   bool
		errMsg    string
	}{
		{
			name: "Success",
			withdraw: &dto.WithdrawWalletDTO{
				UserID: validUserID,
				Amount: validAmount,
			},
			setupMock: func() {
				mockRepo.WithdrawFunc = func(dto *dto.WithdrawWalletDTO) error {
					return nil
				}
			},
			wantErr: false,
		},
		{
			name: "Empty UserID",
			withdraw: &dto.WithdrawWalletDTO{
				UserID: "",
				Amount: validAmount,
			},
			wantErr: true,
			errMsg:  "empty user id",
		},
		{
			name: "Invalid Amount",
			withdraw: &dto.WithdrawWalletDTO{
				UserID: validUserID,
				Amount: -10.0,
			},
			wantErr: true,
			errMsg:  "invalid withdraw amount",
		},
		{
			name: "Zero Amount",
			withdraw: &dto.WithdrawWalletDTO{
				UserID: validUserID,
				Amount: 0.0,
			},
			wantErr: true,
			errMsg:  "invalid withdraw amount",
		},
		{
			name: "Wallet Not Found",
			withdraw: &dto.WithdrawWalletDTO{
				UserID: validUserID,
				Amount: validAmount,
			},
			setupMock: func() {
				mockRepo.WithdrawFunc = func(dto *dto.WithdrawWalletDTO) error {
					return repositories.ErrWalletNotFound
				}
			},
			wantErr: true,
			errMsg:  "wallet not found",
		},
		{
			name: "Insufficient Funds",
			withdraw: &dto.WithdrawWalletDTO{
				UserID: validUserID,
				Amount: 500.0,
			},
			setupMock: func() {
				mockRepo.WithdrawFunc = func(dto *dto.WithdrawWalletDTO) error {
					return errors.New("insufficient funds")
				}
			},
			wantErr: true,
			errMsg:  "insufficient funds",
		},
		{
			name: "Repository Error",
			withdraw: &dto.WithdrawWalletDTO{
				UserID: validUserID,
				Amount: validAmount,
			},
			setupMock: func() {
				mockRepo.WithdrawFunc = func(dto *dto.WithdrawWalletDTO) error {
					return errors.New("database error")
				}
			},
			wantErr: true,
			errMsg:  "database error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset mock before each test
			mockRepo.Reset()
			if tt.setupMock != nil {
				tt.setupMock()
			}

			err := service.WithdrawUserWallet(tt.withdraw)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestServiceIntegration(t *testing.T) {
	// This test simulates a complete flow using the mock repository
	mockRepo := mocks.NewMockWalletRepository()
	ctx := context.Background()
	service := NewWalletService(mockRepo, ctx)

	// 1. Create a wallet
	userID := uuid.New().String()
	walletDTO := &dto.CreateWalletDTO{
		UserID: userID,
	}

	err := service.CreateWallet(walletDTO)
	require.NoError(t, err, "Failed to create wallet")

	// 2. Get the created wallet
	wallet, err := service.GetWalletByUserID(userID)
	require.NoError(t, err, "Failed to get wallet")
	assert.Equal(t, userID, wallet.UserID)
	assert.Equal(t, 0.0, wallet.Balance)

	// 3. Deposit funds
	depositDTO := &dto.DepostiWalletDTO{
		UserID: userID,
		Amount: 100.0,
	}
	err = service.DepositToUserWallet(depositDTO)
	require.NoError(t, err, "Failed to deposit funds")

	// 4. Verify balance after deposit
	updatedWallet, err := service.GetWalletByUserID(userID)
	require.NoError(t, err)
	assert.Equal(t, 100.0, updatedWallet.Balance, "Balance not updated after deposit")

	// 5. Withdraw funds
	withdrawDTO := &dto.WithdrawWalletDTO{
		UserID: userID,
		Amount: 30.0,
	}
	err = service.WithdrawUserWallet(withdrawDTO)
	require.NoError(t, err, "Failed to withdraw funds")

	// 6. Verify balance after withdrawal
	finalWallet, err := service.GetWalletByUserID(userID)
	require.NoError(t, err)
	assert.Equal(t, 70.0, finalWallet.Balance, "Balance not updated after withdrawal")

	// 7. Try to withdraw too much
	excessiveWithdrawDTO := &dto.WithdrawWalletDTO{
		UserID: userID,
		Amount: 100.0,
	}
	err = service.WithdrawUserWallet(excessiveWithdrawDTO)
	require.Error(t, err, "Should not allow excessive withdrawal")
	assert.Contains(t, err.Error(), "insufficient balance")
}
