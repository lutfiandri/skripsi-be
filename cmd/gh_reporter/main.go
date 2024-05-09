package main

import (
	"log"

	"skripsi-be/internal/app/gh_reporter"
	"skripsi-be/internal/config"
	"skripsi-be/internal/infrastructure"
)

func main() {
	// kafkaReader := infrastructure.NewKafkaReader(config.KafkaBrokerUris, config.KafkaGroupGhReporter, config.KafkaTopicDeviceState)
	redisClient := infrastructure.NewRedisClient(config.RedisUri)

	consumer := gh_reporter.NewConsumer(redisClient)
	consumer.StartConsume()
	log.Println("google home reporter service")

	select {}
}
