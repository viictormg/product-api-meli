package usecases

import (
	"fmt"

	"github.com/IBM/sarama"
)

func ConnectProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	return sarama.NewSyncProducer(brokers, config)

}

func PushPrice(topic string, message []byte) error {
	brokers := []string{"localhost:9092"}

	producer, err := ConnectProducer(brokers)
	if err != nil {
		return err
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, oftset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}

	fmt.Println("Message is stored in topic(", msg.Topic, ") partition(", partition, ") and offset(", oftset, ")")

	return nil
}
