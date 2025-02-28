package validation

import "testing"

func TestValidateUserRequiredFields(t *testing.T) {
	tests := []struct {
		name        string
		fullName    string
		email       string
		password    string
		phoneNumber string
		CPF         string
		expected    bool
	}{
		{
			name:        "Valid user",
			fullName:    "User",
			email:       "example@example.com",
			password:    "password",
			phoneNumber: "1234567890",
			CPF:         "12345678901",
			expected:    true,
		},
		{
			name:        "Empty full name",
			fullName:    "",
			email:       "example@example.com",
			password:    "password",
			phoneNumber: "1234567890",
			CPF:         "12345678901",
			expected:    false,
		},
		{
			name:        "Empty email",
			fullName:    "User",
			email:       "",
			password:    "password",
			phoneNumber: "1234567890",
			CPF:         "12345678901",
			expected:    false,
		},
		{
			name:        "Empty password",
			fullName:    "User",
			email:       "example@example.com",
			password:    "",
			phoneNumber: "1234567890",
			CPF:         "12345678901",
			expected:    false,
		},
		{
			name:        "Empty phone number",
			fullName:    "User",
			email:       "example@example.com",
			password:    "password",
			phoneNumber: "",
			CPF:         "12345678901",
			expected:    false,
		},
		{
			name:        "Empty CPF",
			fullName:    "User",
			email:       "example@example.com",
			password:    "password",
			phoneNumber: "1234567890",
			CPF:         "",
			expected:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUserRequiredFields(tt.fullName, tt.email, tt.password, tt.phoneNumber, tt.CPF)
			if tt.expected {
				if err != nil {
					t.Errorf("Expected %v, got %v", nil, err)
				}
			} else {
				if err == nil {
					t.Errorf("Expected error, got %v", nil)
				}
			}
		})
	}
}
