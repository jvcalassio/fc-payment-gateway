package entity

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("999999999999999", "Joao Silva", 01, 2099, 999)
	assert.Equal(t, "invalid credit card number", err.Error())

	_, err = NewCreditCard("5438330212460978", "Joao Silva", 01, 2099, 999)
	assert.Nil(t, err)
}

func TestCreditCardExpMonth(t *testing.T) {
	_, err := NewCreditCard("5438330212460978", "Joao Silva", 13, 2099, 999)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("5438330212460978", "Joao Silva", 0, 2099, 999)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("5438330212460978", "Joao Silva", 7, 2099, 999)
	assert.Nil(t, err)
}

func TestCreditCardExpYear(t *testing.T) {
	lastYear := time.Now().AddDate(-1, 0, 0)

	_, err := NewCreditCard("5438330212460978", "Joao Silva", 7, lastYear.Year(), 999)
	assert.Equal(t, "invalid expiration year", err.Error())

	_, err = NewCreditCard("5438330212460978", "Joao Silva", int(time.Now().Month()), time.Now().Year(), 999)
	assert.Equal(t, "expired credit card", err.Error())
	
	_, err = NewCreditCard("5438330212460978", "Joao Silva", int(time.Now().AddDate(0, 1, 0).Month()), time.Now().Year(), 999)
	assert.Nil(t, err)
}