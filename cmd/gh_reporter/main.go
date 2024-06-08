package main

import (
	"context"
	"log"

	"skripsi-be/internal/app/gh_reporter"
	"skripsi-be/internal/config"
	"skripsi-be/internal/infrastructure"
	"skripsi-be/internal/util/helper"

	"google.golang.org/api/homegraph/v1"
)

func main() {
	// kafkaReader := infrastructure.NewKafkaReader(config.KafkaBrokerUris, config.KafkaGroupGhReporter, config.KafkaTopicDeviceState)
	redisClient := infrastructure.NewRedisClient(config.RedisUri)

	homegraphService, err := homegraph.NewService(context.Background())
	helper.PanicIfErr(err)
	homegraphDeviceService := homegraph.NewDevicesService(homegraphService)

	consumer := gh_reporter.NewConsumer(redisClient, homegraphDeviceService)
	consumer.StartConsume()
	log.Println("google home reporter service")

	select {}
}
