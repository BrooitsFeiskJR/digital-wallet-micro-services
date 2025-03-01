package routers

import (
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/api/handlers"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(
	engine *gin.Engine,
	authHandler *handlers.AuthHandler,
) {
	v1 := engine.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.RegisterHandler)
			auth.POST("/login", authHandler.LoginUser)
		}

	}
}
