package routers

import (
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/api/handlers"
	"github.com/gin-gonic/gin"
)

func Initialize(authHandler *handlers.AuthHandler, userHandler *handlers.UserHandler) {
	router := gin.Default()
	initializeRoutes(router, authHandler, userHandler)
	router.Run(":8080")
}
