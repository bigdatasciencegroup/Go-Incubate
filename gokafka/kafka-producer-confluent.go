package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// NewProducer creates producers
func NewProducer(broker string) (*kafka.Producer, chan bool, error){
	doneChan := make(chan bool)

	// Create producer
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		return p, doneChan, err
	}
	fmt.Printf("Created Producer %v\n", p)

	go func() {
		defer close(doneChan)
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				m := ev
				if m.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
				} else {
					fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
						*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
				}
				return
			default:
				fmt.Printf("Ignored event: %s\n", ev)
			}
		}
	}()

	return p, doneChan, err
}
