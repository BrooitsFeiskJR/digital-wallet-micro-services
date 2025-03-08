package routers

import (
	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/api/handlers"
	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/api/middleware"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(
	engine *gin.Engine,
	walletHandler *handlers.WalletHandler,
) {
	v1 := engine.Group("/api/v1")
	{
		wallet := v1.Group("/wallet").Use(middleware.VerifyJwtMiddleware())
		{
			wallet.GET("/", walletHandler.GetWallet)
			wallet.POST("/deposit", walletHandler.DepositHandler)
			wallet.POST("/withdraw", walletHandler.WithdrawHandler)
		}
	}
}
