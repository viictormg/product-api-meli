package producer

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/viictormg/product-api-meli/config"
)

type KafkaProducer struct {
	Producer sarama.SyncProducer
	Brokers  []string
	Topic    string
}

func NewKafkaProducer(config *config.Config) *KafkaProducer {
	kafkaConfig := config.GeKafkaConfg()

	fmt.Println("DDD", kafkaConfig.Topic)

	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Retry.Max = kafkaConfig.Retry

	producer, err := sarama.NewSyncProducer(kafkaConfig.Brokers, cfg)

	if err != nil {
		fmt.Println("Error connecting to Kafka:", err)
	}

	return &KafkaProducer{
		Producer: producer,
		Brokers:  config.Brokers,
		Topic:    config.Topic,
	}
}
