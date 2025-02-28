package validation

import "testing"

func TestPhoneNumberValidation(t *testing.T) {
	tests := []struct {
		name     string
		phone    string
		expected bool
	}{
		{
			name:     "Valid phone number",
			phone:    "08123456789",
			expected: true,
		},
		{
			name:     "Empty phone number",
			phone:    "",
			expected: false,
		},
		{
			name:     "Invalid phone number",
			phone:    "0812345678",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidPhoneNumber(tt.phone)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
