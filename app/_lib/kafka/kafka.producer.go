package kafka

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"kafka-demoset/app/conf"
	"kafka-demoset/app/internal/logger"
)

var Producer *kafka.Producer

func InitKafkaProducer() {
	if Producer == nil {
		initKafkaProducer()
	}
}

func initKafkaProducer() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   conf.Kafka().Addr,
		"delivery.timeout.ms": "1",
		"acks":                "all",
		"enable.idempotence":  "true",
	}

	producer, err := kafka.NewProducer(configMap)
	if err != nil {
		logger.Sugar.Error(err)
		panic(err)
	}

	logger.Sugar.Info("Kafka Producer started ... ")

	Producer = producer
}

func Publish(msg any, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	bytes, err := json.Marshal(msg)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	message := &kafka.Message{
		Value: bytes,
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key: key,
	}
	err = producer.Produce(message, deliveryChan)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	return nil
}

func DeliveryReport(deliveryChan chan kafka.Event) {
	for e := range deliveryChan {
		switch e.(type) {
		case *kafka.Message:
			e := <-deliveryChan
			msg := e.(*kafka.Message)

			if msg.TopicPartition.Error != nil {
				logger.Sugar.Error("Error in topic partition ... ")
			} else {
				logger.Sugar.Debug("Message event : ", msg.TopicPartition)
				// @TODO something what you want over msg
			}
		}
	}
}
