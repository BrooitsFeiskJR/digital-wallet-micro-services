package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var secretKey = []byte(os.Getenv("SECRECT_KEY"))

func CreateToken(id uuid.UUID, name, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"name":  name,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 8).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string, claimsParam jwt.Claims) error {
	token, err := jwt.ParseWithClaims(tokenString, claimsParam, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("invalid token claims")
	}

	expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
	if time.Now().UTC().After(expirationTime) {
		return fmt.Errorf("token has expired")
	}
	return nil
}
