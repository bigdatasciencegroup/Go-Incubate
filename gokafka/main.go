package main

import (
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	broker := os.Getenv("KAFKAPORT")
	topic := os.Getenv("TOPICNAME")

	p, doneChan, err := NewProducer(broker)
	if err != nil {
		os.Exit(1)
	}

	for ii := 0; ii < 7; ii++ {
		value := "Hello Go it is a success!"
		p.ProduceChannel() <- &kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}, Value: []byte(value)}

		// wait for delivery report goroutine to finish
		_ = <-doneChan
	}

	p.Close()
}
