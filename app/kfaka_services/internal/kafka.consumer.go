package internal

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	libkafka "kafka-demoset/app/_lib/kafka"
	"kafka-demoset/app/conf"
	"kafka-demoset/app/internal/logger"
)

func NewKafkaEventConsumer() {
	newKafkaConsumer(libkafka.Test_Kafka_Message_Topic, handleTestMessageKafkaEvent)
}

func newKafkaConsumer(topic string, handle func(message *kafka.Message) error) error {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": conf.Kafka().Addr,
		"client.id":         "goapp_consumer",
		"group.id":          "goapp_group",
	}

	consumer, err := kafka.NewConsumer(configMap)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	topics := []string{topic}
	err = consumer.SubscribeTopics(topics, nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			logger.Sugar.Error(err)
			return err
		}

		go handle(msg)
	}
}

func handleTestMessageKafkaEvent(message *kafka.Message) error {
	defer func() {}()
	logger.Sugar.Debug(libkafka.Test_Kafka_Message_Topic)

	msg := &libkafka.TestKafkaMessageData{}
	if err := json.Unmarshal(message.Value, msg); err != nil {
		logger.Sugar.Error(err)
		return err
	}

	if msg.Message != "" {
		go func() {
			fmt.Println("send to telegram")

			msgV := fmt.Sprintf("Kafka Message : %s", msg.Message)
			_, err := TelegramBot.SendMessage(msgV)
			if err != nil {
				logger.Sugar.Error(err)
				return
			}

			logger.Sugar.Debug("send to telegram success")
		}()
	}

	return nil
}
