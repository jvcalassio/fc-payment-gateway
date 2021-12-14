package main

import (
	"encoding/json"
	"database/sql"
	"github.com/jvcalassio/fc-payment-gateway/adapter/broker/kafka"
	"github.com/jvcalassio/fc-payment-gateway/adapter/factory"
	"github.com/jvcalassio/fc-payment-gateway/adapter/presenter/transaction"
	"github.com/jvcalassio/fc-payment-gateway/adapter/repository/fixture"
	"github.com/jvcalassio/fc-payment-gateway/usecase/process_transaction"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"os"
	"log"
)

func main() {
	// db
	db, err := sql.Open("mysql", os.Getenv("MYSQL_USERNAME")+":"+os.Getenv("MYSQL_PASSWORD")+"@tcp("+os.Getenv("MYSQL_HOST")+":3306)/"+os.Getenv("MYSQL_DATABASE"))
	if err != nil {
		log.Fatal(err)
	}
	migrationsDir := os.DirFS("adapter/repository/fixture/sql")
	migrationsInstance := fixture.NewMigrationsDb(db)
	db = migrationsInstance.Up(migrationsDir)
	// repository
	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)
	repository := repositoryFactory.CreateTransactionRepository()
	// producer + cfgmap
	configMapProducer := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("BOOTSTRAP_SERVERS"),
	}
	kafkaPresenter := transaction.NewTransactionKafkaPresenter()
	producer := kafka.NewKafkaProducer(configMapProducer, kafkaPresenter)
	// consumer + cfgmap
	var msgChan = make(chan *ckafka.Message)
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("BOOTSTRAP_SERVERS"),
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