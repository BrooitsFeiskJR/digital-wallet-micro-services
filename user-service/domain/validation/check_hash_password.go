package validation

import "golang.org/x/crypto/bcrypt"

func CheckHashPassword(hashPassword, providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
