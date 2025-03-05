package entities

import (
	"testing"

	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/domain/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateWallet(t *testing.T) {
	tests := []struct {
		name      string
		userID    string
		wantErr   error
		checkFunc func(*Wallet) bool
	}{
		{
			name:    "empty userID",
			userID:  "",
			wantErr: ErrEmptyUserID,
		},
		{
			name:    "invalid userID",
			userID:  "invalidObjectID",
			wantErr: ErrConvertObjectID,
		},
		{
			name:    "valid userID",
			userID:  primitive.NewObjectID().Hex(),
			wantErr: nil,
			checkFunc: func(w *Wallet) bool {
				return w.Balance == 0 && !w.CreateAt.IsZero() && !w.UpdateAt.IsZero()
			},
		},
	}

	for _, tt := range tests {
		dto := &dto.CreateWalletDTO{
			UserID: tt.userID,
		}
		t.Run(tt.name, func(t *testing.T) {
			wallet, err := CreateWallet(dto)
			if err != tt.wantErr {
				t.Errorf("CreateWallet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.checkFunc != nil && !tt.checkFunc(wallet) {
				t.Errorf("CreateWallet() checkFunc failed")
			}
		})
	}
}
