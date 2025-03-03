package validatiion

import (
	"testing"
)

func TestValidateType(t *testing.T) {
	tests := []struct {
		name            string
		transactionType string
		wantErr         bool
	}{
		{"ValidDeposit", Deposit, false},
		{"ValidWithdrawal", Withdrawal, false},
		{"ValidTransfer", Transfer, false},
		{"InvalidType", "invalid", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateType(tt.transactionType)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
