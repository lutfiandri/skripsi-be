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
