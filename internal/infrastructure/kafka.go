package infrastructure

import (
	"strings"

	"github.com/segmentio/kafka-go"
)

func NewKafkaWriter(brokerUris string) *kafka.Writer {
	brokers := strings.Split(brokerUris, ",")

	return &kafka.Writer{
		Addr:                   kafka.TCP(brokers...),
		AllowAutoTopicCreation: true,
	}
}

func NewKafkaReader(brokerUris string, groupId, topic string) *kafka.Reader {
	brokers := strings.Split(brokerUris, ",")

	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		GroupID: groupId,
		Topic:   topic,
	})
}
