package validation

import "testing"

func TestCPFValidation(t *testing.T) {
	tests := []struct {
		name     string
		cpf      string
		expected bool
	}{
		{
			name:     "Valid CPF",
			cpf:      "12345678909",
			expected: true,
		},
		{
			name:     "Invalid CPF",
			cpf:      "12345678901",
			expected: false,
		},
		{
			name:     "Empty CPF",
			cpf:      "",
			expected: false,
		},
		{
			name:     "Invalid CPF with letters",
			cpf:      "1234567890a",
			expected: false,
		},
		{
			name:     "Invalid CPF with special characters",
			cpf:      "1234567890@",
			expected: false,
		},
		{
			name:     "Invalid CPF with spaces",
			cpf:      "1234567890 9",
			expected: false,
		},
		{
			name:     "Invalid CPF with less than 11 digits",
			cpf:      "1234567890",
			expected: false,
		},
		{
			name:     "Invalid CPF with more than 11 digits",
			cpf:      "123456789012",
			expected: false,
		},
		{
			name:     "Invalid CPF with 11 zeros",
			cpf:      "00000000000",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CheckCPF(tt.cpf)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
