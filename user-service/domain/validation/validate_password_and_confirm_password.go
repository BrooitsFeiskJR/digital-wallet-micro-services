package validation

func ValidatePasswordAndConfirmPassword(password, confirmPassword string) bool {
	return password == confirmPassword
}
