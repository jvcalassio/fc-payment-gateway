package entity

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTransaction_Amount(t *testing.T) {
	transaction := NewTransaction()
	// transaction.ID = "1"
	// transaction.AccountID = "1"
	transaction.Amount = 1001

	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "no limit for this transaction", err.Error())

	transaction.Amount = 0.50
	err = transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "minimum amount for transaction is 1", err.Error())

	transaction.Amount = 550
	err = transaction.IsValid()
	assert.Nil(t, err)
}