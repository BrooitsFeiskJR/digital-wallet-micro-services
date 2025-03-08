package validatiion

import "errors"

var (
	ErrInsufficientBalance = errors.New("insufficient balance")
)

func ValidateSuficientBalance(balance, amount float64) error {
	if amount > balance {
		return ErrInsufficientBalance
	}
	return nil
}
