package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Wallet struct {
	ID primitive.ObjectID
}

func TestCreateTransaction(t *testing.T) {
	fromWallet := primitive.NewObjectID()
	toWallet := primitive.NewObjectID()
	amount := 100.0
	tType := Deposit

	transaction, err := CreateTransaction(fromWallet, toWallet, amount, tType)
	assert.NoError(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, fromWallet, transaction.SendWalletID)
	assert.Equal(t, toWallet, transaction.ToWalletID)
	assert.Equal(t, amount, transaction.Amount)
	assert.Equal(t, tType, transaction.Type)
	assert.WithinDuration(t, time.Now(), transaction.CreatedAt, time.Second)

	invalidType := "invalid"
	transaction, err = CreateTransaction(fromWallet, toWallet, amount, invalidType)
	assert.Error(t, err)
	assert.Nil(t, transaction)
}
