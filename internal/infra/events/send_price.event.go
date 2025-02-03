package events

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/viictormg/product-api-meli/internal/application/price/ports"
	"github.com/viictormg/product-api-meli/internal/infra/clients/producer"
)

type PriceEvent struct {
	provider producer.KafkaProducer
}

func NewPriceEvent(provider *producer.KafkaProducer) ports.PriceEventyIF {
	return &PriceEvent{
		provider: *provider,
	}
}

func (s *PriceEvent) SendPriceEvent(message []byte) {
	msg := &sarama.ProducerMessage{
		Topic: s.provider.Topic,
		Value: sarama.StringEncoder(message),
	}

	partition, oftset, err := s.provider.Producer.SendMessage(msg)

	if err != nil {
		fmt.Println("Error sending message to Kafka:", err)
	}

	fmt.Println(fmt.Printf("Message is stored in topic(%s) partition(%d) and offset(%d)", msg.Topic, partition, oftset))
}

func (s *PriceEvent) Close() {
	s.provider.Producer.Close()
}
