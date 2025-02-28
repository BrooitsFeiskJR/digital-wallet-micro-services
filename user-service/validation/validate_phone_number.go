package validation

import "regexp"

func ValidPhoneNumber(phoneNumber string) bool {
	re := regexp.MustCompile(`^\d{2}\d{5}\d{4}$`)
	return re.MatchString(phoneNumber)
}
