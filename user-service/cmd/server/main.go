package main

import (
	"fmt"

	"github.com/BrooitsFeiskJR/digital-wallet-user-service/infra/db"
)

func main() {
	connString := db.ConnectionString()
	_, err := db.ConnectToDB(connString)
	if err != nil {
		fmt.Printf("error connecting to database: %v", err)
		return
	}
	fmt.Println("Hello, Digital Wallet!")
}
