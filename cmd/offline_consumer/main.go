package main

import (
	"log"

	"skripsi-be/internal/app/offline_consumer"
	"skripsi-be/internal/config"
	"skripsi-be/internal/infrastructure"
)

func main() {
	mongo := infrastructure.NewMongoDatabase(config.MongoUri, config.MongoDbName)
	redisClient := infrastructure.NewRedisClient(config.RedisUri)

	// Set the notify-keyspace-events configuration to enable expiration events
	_, err := redisClient.Do("CONFIG", "SET", "notify-keyspace-events", "Ex").Result()
	if err != nil {
		log.Panicf("Unable to set keyspace events: %v\n", err)
	}

	repository := offline_consumer.NewRepository(mongo)
	consumer := offline_consumer.NewConsumer(repository, redisClient)
	consumer.StartConsume()

	select {}
}
