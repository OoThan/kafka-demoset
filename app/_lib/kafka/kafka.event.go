package kafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

func TestKafkaMessageEvent(message string) {
	deliveryChan := make(chan kafka.Event)
	go DeliveryReport(deliveryChan)

	Publish(&TestKafkaMessageData{Message: message}, Test_Kafka_Message_Topic, Producer, nil, deliveryChan)
}
