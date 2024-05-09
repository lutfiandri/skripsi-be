package main

import (
	"log"

	"skripsi-be/internal/app/gh_reporter"
	"skripsi-be/internal/config"
	"skripsi-be/internal/infrastructure"
)

func main() {
	kafkaReader := infrastructure.NewKafkaReader(config.KafkaBrokerUris, config.KafkaGroupGhReporter, config.KafkaTopicDeviceState)
	consumer := gh_reporter.NewConsumer(kafkaReader)
	consumer.StartConsume()
	log.Println("google home reporter service")

	select {}
}
