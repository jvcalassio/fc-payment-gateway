package process_transaction

type TransactionDtoInput struct {
	ID string `json:"id"`
	AccountID string `json:"account_id"`
	CreditCardNumber string `json:"credit_card_number"`
	CreditCardName string `json:"credit_card_name"`
	CreditCardExpMonth int `json:"credit_card_exp_month"`
	CreditCardExpYear int `json:"credit_card_exp_year"`
	CreditCardCVV int `json:"credit_card_cvv"`
	Amount float64 `json:"amount"`
}

type TransactionDtoOutput struct {
	ID string `json:"id"`
	Status string `json:"status"`
	ErrorMessage string `json:"error_message"`
}