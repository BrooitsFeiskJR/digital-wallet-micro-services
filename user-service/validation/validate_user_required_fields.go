package validation

import "errors"

func ValidateUserRequiredFields(name, email, password, phoneNumber, CPF string) error {
	requiredFields := map[string]string{
		"full name":    name,
		"email":        email,
		"password":     password,
		"phone number": phoneNumber,
		"CPF":          CPF,
	}

	for fieldName, fieldValue := range requiredFields {
		if !isValidField(fieldValue) {
			return errors.New(fieldName + " is required")
		}
	}
	return nil
}
