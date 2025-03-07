package middleware

import (
	"net/http"
	"strings"

	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/api/responses"
	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/infra/config/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func VerifyJwtMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, BEARER_SCHEMA) {
			response := responses.ErrorResponse(http.StatusUnauthorized, "request does not contain a valid access token")
			response.ToJSON(context, http.StatusUnauthorized)
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]

		// Create a claims variable to hold the token claims
		claims := jwt.MapClaims{}

		// Verify the token and parse the claims
		err := auth.VerifyToken(tokenString, claims)
		if err != nil {
			response := responses.ErrorResponse(http.StatusUnauthorized, err.Error())
			response.ToJSON(context, http.StatusUnauthorized)
			return
		}

		// Store the claims in the context for posterior requests
		context.Set("JWT_Claims", claims)

		context.Next()
	}
}
