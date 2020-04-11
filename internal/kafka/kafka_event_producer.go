package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
)

type KafkaEventProducer interface {
	Produce(event interface{}) error
	Close()
}

type DefaultKafkaEventProducer struct {
	brokerHosts []string
	topic       string
	producer    sarama.SyncProducer
}

func NewKafkaEventProducer(brokerHosts []string, topic string) KafkaEventProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer(brokerHosts, config)
	if err != nil {
		// Should not reach here
		panic(err)
	}
	return &DefaultKafkaEventProducer{
		brokerHosts: brokerHosts,
		topic:       topic,
		producer:    producer,
	}
}

func (k *DefaultKafkaEventProducer) Produce(event interface{}) error {
	bytes, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: k.topic,
		Value: sarama.ByteEncoder(bytes),
	}

	partition, offset, err := k.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", k.topic, partition, offset)
	return nil
}

func (k *DefaultKafkaEventProducer) Close() {
	if err := k.producer.Close(); err != nil {
		// Should not reach here
		panic(err)
	}
}
