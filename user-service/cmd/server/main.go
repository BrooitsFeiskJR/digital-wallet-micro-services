package main

import (
	"fmt"

	"github.com/BrooitsFeiskJR/digital-wallet-user-service/api/config"
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/api/routers"
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/infra/db"
)

func main() {
	connString := db.ConnectionString()
	db, err := db.GetDBConnection(connString)
	if err != nil {
		fmt.Printf("error connecting to database: %v", err)
		return
	}
	authHandler := config.AuthHandlerSetup(db)
	routers.Initialize(authHandler)
}
