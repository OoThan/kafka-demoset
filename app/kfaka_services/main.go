package main

import (
	_ "kafka-demoset/app/conf"
	"kafka-demoset/app/internal/logger"
	"kafka-demoset/app/kfaka_services/internal"
)

func main() {
	internal.NewKafkaService()
	internal.NewKafkaEventConsumer()

	logger.Sugar.Info("Kafka Consumer started ... ")
	select {}
}
