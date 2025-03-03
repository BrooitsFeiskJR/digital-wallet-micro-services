package validatiion

import (
	"errors"
)

const (
	Deposit    = "deposit"
	Withdrawal = "withdrawal"
	Transfer   = "transfer"
)

func ValidateType(transactionType string) error {
	switch transactionType {
	case Deposit, Withdrawal, Transfer:
		return nil
	default:
		return errors.New("invalid transaction type")
	}
}
