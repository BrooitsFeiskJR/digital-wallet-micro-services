package validatiion

import "errors"

var (
	ErrInsufficientBalance = errors.New("insufficient balance")
)

func ValidateSuficientBalance(balance, amount float64) error {
	if balance < amount {
		return ErrInsufficientBalance
	}
	return nil
}
