package main

import (
	"fmt"
	"os"

	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/infra/db"
)

func main() {
	fmt.Println("Hello, Wallet Service!")
	uri, err := db.GetMongoURI()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get mongo uri: %v\n", err)
	}
	_, err = db.ConnectMongoDB(uri)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to connect to mongo: %v\n", err)
	}
}
