package validatiion

import "testing"

func TestValidateSuficientBalance(t *testing.T) {
	tests := []struct {
		name    string
		balance float64
		amount  float64
		wantErr error
	}{
		{
			name:    "sufficient balance",
			balance: 100.0,
			amount:  10.0,
			wantErr: nil,
		},
		{
			name:    "insufficient balance",
			balance: 5.0,
			amount:  10.0,
			wantErr: ErrInsufficientBalance,
		},
		{
			name:    "zero balance",
			balance: 0.0,
			amount:  10.0,
			wantErr: ErrInsufficientBalance,
		},
		{
			name:    "negative balance",
			balance: -5.0,
			amount:  10.0,
			wantErr: ErrInsufficientBalance,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSuficientBalance(tt.balance, tt.amount)
			if err != tt.wantErr {
				t.Errorf("ValidateSufticientBalance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
