package entity

import (
	"regexp"
	"errors"
	"time"
)

type CreditCard struct {
	number string
	name string
	expMonth int
	expYear int
	cvv int
}

func NewCreditCard(number string, name string, expMonth int, expYear int, cvv int) (*CreditCard, error) {
	cc := &CreditCard{
		number,
		name,
		expMonth,
		expYear,
		cvv,
	}

	err := cc.IsValid()

	if err != nil {
		return nil, err
	}

	return cc, nil
}

func (c *CreditCard) IsValid() error {
	err := c.validateNumber()
	if err != nil {
		return err
	}

	err = c.validateMonth()
	if err != nil {
		return err
	}

	err = c.validateYear()
	if err != nil {
		return err
	}

	if time.Now().Year() == c.expYear && int(time.Now().Month()) <= c.expMonth {
		return errors.New("expired credit card")
	}

	return nil
}

func (c *CreditCard) validateNumber() error {
	re := regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$`)
	if re.MatchString(c.number) {
		return nil
	}
	return errors.New("invalid credit card number")
}

func (c *CreditCard) validateMonth() error {
	if c.expMonth > 0 && c.expMonth < 13 {
		return nil
	}
	return errors.New("invalid expiration month")
}
func (c *CreditCard) validateYear() error {
	if c.expYear >= time.Now().Year() {
		return nil
	}
	return errors.New("invalid expiration year")
}