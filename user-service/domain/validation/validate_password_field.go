package validation

func ValidatePasswordField(password string) bool {
	return len(password) >= 8
}
