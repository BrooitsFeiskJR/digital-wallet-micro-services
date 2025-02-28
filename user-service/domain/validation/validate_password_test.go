package validation

import "testing"

func TestValidatePasswordField(t *testing.T) {
	password := "password"
	result := ValidatePasswordField(password)
	if !result {
		t.Errorf("ValidatePasswordField(%s) = %t; want true", password, result)
	}
}

func TestInvalidPassword(t *testing.T) {
	password := "short"
	result := ValidatePasswordField(password)
	if result {
		t.Errorf("ValidatePasswordField(%s) = %t; want false", password, result)
	}
}
