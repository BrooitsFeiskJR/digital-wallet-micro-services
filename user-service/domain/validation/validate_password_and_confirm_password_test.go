package validation

import "testing"

func TestValidationPasswordAndConfirmPassword(t *testing.T) {
	password := "password"
	confirmPassword := "password"
	result := ValidatePasswordAndConfirmPassword(password, confirmPassword)
	if !result {
		t.Errorf("Expected %v, got %v", true, result)
	}
}

func TestValidationPasswordAndConfirmPasswordDifferent(t *testing.T) {
	password := "password"
	confirmPassword := "password123"
	result := ValidatePasswordAndConfirmPassword(password, confirmPassword)
	if result {
		t.Errorf("Expected %v, got %v", false, result)
	}
}
