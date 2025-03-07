package routers

import (
	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/api/handlers"
	"github.com/gin-gonic/gin"
)

// Initialize sets up routes and returns the router
func Initialize(handler *handlers.WalletHandler) *gin.Engine {
	router := gin.Default()
	initializeRoutes(router, handler)
	return router
}
