package routers

import (
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/api/handlers"
	"github.com/gin-gonic/gin"
)

func Initialize(authHandler *handlers.AuthHandler) {
	router := gin.Default()
	initializeRoutes(router, authHandler)
	router.Run(":8080")
}
