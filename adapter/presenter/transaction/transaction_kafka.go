package transaction

import (
	"github.com/jvcalassio/fc-payment-gateway/usecase/process_transaction"
	"encoding/json"
)

type KafkaPresenter struct {
	ID string `json:"id"`
	Status string `json:"status"`
	ErrorMessage string `json:"error_message"`
}
func NewTransactionKafkaPresenter() *KafkaPresenter {
	return &KafkaPresenter{}
}

func (t *KafkaPresenter) Bind(dtoOutput interface{}) error {
	t.ID = dtoOutput.(process_transaction.TransactionDtoOutput).ID
	t.Status = dtoOutput.(process_transaction.TransactionDtoOutput).Status
	t.ErrorMessage = dtoOutput.(process_transaction.TransactionDtoOutput).ErrorMessage
	return nil
}

func (t *KafkaPresenter) Show() ([] byte, error) {
	j, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return j, nil
}