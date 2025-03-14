package auth

import (
	"crypto/rsa"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func init() {
	privateKeyPEM := strings.ReplaceAll(os.Getenv("JWT_PRIVATE_KEY"), "\\n", "\n")
	publicPEM := strings.ReplaceAll(os.Getenv("JWT_PUBLIC_KEY"), "\\n", "\n")

	var err error
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKeyPEM))
	if err != nil {
		fmt.Printf("Error parsing private key: %v\n", err)
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM([]byte(publicPEM))
	if err != nil {
		fmt.Printf("Error parsing public key: %v\n", err)
	}
}

func CreateAccessToken(id uuid.UUID, name, email string) (string, error) {
	if privateKey == nil {
		return "", fmt.Errorf("private key not initialized")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"id":    id,
		"name":  name,
		"email": email,
		"exp":   time.Now().Add(time.Minute * 5).Unix(),
	})

	accessToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func CreateRefreshAcessToken(id uuid.UUID, name, email string) (string, error) {
	if privateKey == nil {
		return "", fmt.Errorf("private key not initialized")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"id":    id,
		"name":  name,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 48).Unix(), // 2 days
	})
	refreshToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}

func VerifyToken(tokenString string, claimsParam jwt.Claims) error {
	token, err := jwt.ParseWithClaims(tokenString, claimsParam, func(token *jwt.Token) (any, error) {
		return publicKey, nil
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
