package kafkasw

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
	"github.com/wvanbergen/kafka/consumergroup"
)

type messageHandler func(*sarama.ConsumerMessage) error

//ConsumerGroup contains the consumer connection parameters
type ConsumerGroup struct {
	GroupName string   //ConsumerGroup name
	Topics    []string //Topic name
	Zookeeper []string //Zookeeper IP:Port address
}

//ConsumeMessages creates consumer group to consume the messages
func ConsumeMessages(consumerGroup ConsumerGroup, handler messageHandler) {
	log.Println("Starting Consumer")
	config := consumergroup.NewConfig()
	config.Offsets.Initial = sarama.OffsetOldest
	config.Offsets.ProcessingTimeout = 10 * time.Second

	consumer, err := consumergroup.JoinConsumerGroup(
		consumerGroup.GroupName,
		consumerGroup.Topics,
		consumerGroup.Zookeeper,
		config)
	if err != nil {
		log.Fatal("Failed to join consumer group", consumerGroup, err)
	}

	//Relay incoming signals to channel 'c'
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	//Terminate the consumer gracefully upon receiving a kill signal.
	go func() {
		<-c
		if err := consumer.Close(); err != nil {
			log.Println("Error closing the consumer", err)
		}

		log.Println("Consumer closed")
		os.Exit(0)
	}()

	//Read from the Errors() channel to avoid producer deadlock
	go func() {
		for err := range consumer.Errors() {
			log.Println(err)
		}
	}()

	log.Println("Waiting for messages")
	for message := range consumer.Messages() {
		log.Printf("Topic: %s\t Partition: %v\t Offset: %v\n", message.Topic, message.Partition, message.Offset)

		e := handler(message)
		if e != nil {
			log.Fatal(e)
			consumer.Close()
		} else {
			consumer.CommitUpto(message)
		}
	}
}
