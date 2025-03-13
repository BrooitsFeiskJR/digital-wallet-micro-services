package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var secretKey = []byte(os.Getenv("SECRECT_KEY"))

func CreateAccessToken(id uuid.UUID, name, email string) (string, error) {
	// TODO: Change algortim to RS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"name":  name,
		"email": email,
		"exp":   time.Now().Add(time.Minute * 5).Unix(),
	})

	accessToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func CreateRefreshAcessToken(id uuid.UUID, name, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"name":  name,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 48).Unix(), // 2 days
	})
	refreshToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}
