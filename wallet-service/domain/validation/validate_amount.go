package validatiion

import "errors"

var (
	ErrInvalidAmount = errors.New("invalid amount: amount must be greater than 0")
)

func ValidateAmount(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	return nil
}
