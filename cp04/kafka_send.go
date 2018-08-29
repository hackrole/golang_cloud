package main

import (
	"github.com/Shopify/sarama"
	"strings"
)

func main() {
	brokerList := os.GetEnv("KAFKA_BROKERS")
	if brokerList == "" {
		brokerList = "localhost:9092"
	}
	brokers := strings.Split(brokerList, ",")

	config := sarama.NewConfig()
	brokers := []string{"localhost:9029"}
	client, err := saram.NewClient(brokers, config)

	if err != nil {
		panic(err)
	}

	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		panic(err)
	}
}
