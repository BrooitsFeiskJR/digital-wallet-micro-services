package validatiion

import (
	"testing"
)

func TestValidateAmount(t *testing.T) {
	tests := []struct {
		name    string
		amount  float64
		wantErr error
	}{
		{
			name:    "valid amount",
			amount:  10.0,
			wantErr: nil,
		},
		{
			name:    "zero amount",
			amount:  0.0,
			wantErr: ErrInvalidAmount,
		},
		{
			name:    "negative amount",
			amount:  -5.0,
			wantErr: ErrInvalidAmount,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateAmount(tt.amount)
			if err != tt.wantErr {
				t.Errorf("ValidateAmount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
