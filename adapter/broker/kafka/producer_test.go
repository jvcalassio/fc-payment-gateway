package kafka

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/jvcalassio/fc-payment-gateway/domain/entity"
	"github.com/jvcalassio/fc-payment-gateway/usecase/process_transaction"
	"github.com/jvcalassio/fc-payment-gateway/adapter/presenter/transaction"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func TestProducer_Publish(t *testing.T) {
	expectedOutput := process_transaction.TransactionDtoOutput{
		ID: "1",
		Status: entity.REJECTED,
		ErrorMessage: "no limit for this transaction",
	}
	// outputJson, _ := json.Marshal(expectedOutput)

	configMap := ckafka.ConfigMap{
		"test.mock.num.brokers": 3,
	}
	producer := NewKafkaProducer(&configMap, transaction.NewTransactionKafkaPresenter())
	err := producer.Publish(expectedOutput, []byte("1"), "test")
	assert.Nil(t, err)
}