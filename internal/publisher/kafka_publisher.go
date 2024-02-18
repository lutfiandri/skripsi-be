package publisher

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type KafkaPublisher interface {
	Publish(ctx context.Context, topic string, message interface{}) error
}

type kafkaPublisher struct {
	writer *kafka.Writer
}

func NewKafkaPublisher(writer *kafka.Writer) KafkaPublisher {
	return &kafkaPublisher{
		writer: writer,
	}
}

func (publisher *kafkaPublisher) Publish(ctx context.Context, topic string, message interface{}) error {
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	if err := publisher.writer.WriteMessages(ctx, kafka.Message{
		Topic: topic,
		Key:   []byte(uuid.NewString()),
		Value: messageBytes,
	}); err != nil {
		return err
	}

	return nil
}
