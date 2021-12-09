package main

import (
	"database/sql"
	"encoding/json"
	"github.com/jvcalassio/fc-payment-gateway/adapter/broker/kafka"
	"github.com/jvcalassio/fc-payment-gateway/adapter/factory"
	"github.com/jvcalassio/fc-payment-gateway/adapter/presenter/transaction"
	"github.com/jvcalassio/fc-payment-gateway/usecase/process_transaction"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"fmt"
)

func main() {
	// db
	db, err := sql.Open("sqlite3", "transactions.db")
	if err != nil {
		log.Fatal(err)
	}
	// repository
	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)
	repository := repositoryFactory.CreateTransactionRepository()
	// producer + cfgmap
	configMapProducer := &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
	}
	kafkaPresenter := transaction.NewTransactionKafkaPresenter()
	producer := kafka.NewKafkaProducer(configMapProducer, kafkaPresenter)
	// consumer + cfgmap
	var msgChan = make(chan *ckafka.Message)
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
		"client.id": "goapp",
		"group.id": "goapp",
	}
	topics := []string{"transactions"}
	consumer := kafka.NewConsumer(configMapConsumer, topics)
	go consumer.Consume(msgChan)
	// usecase
	usecase := process_transaction.NewProcessTransaction(repository, producer, "transactions_result")

	fmt.Println("Ready.")
	
	for msg := range msgChan {
		var input process_transaction.TransactionDtoInput
		json.Unmarshal(msg.Value, &input)
		usecase.Execute(input)
	}
}