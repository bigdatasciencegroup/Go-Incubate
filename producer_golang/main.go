package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
	"github.com/adaickalavan/Go-Rest-Kafka-Mongo/kafkapc"
)

//Hooks that may be overridden for testing
var inputReader io.Reader = os.Stdin
var outputWriter io.Writer = os.Stdout

// Instantiate a producer
var producer sarama.AsyncProducer

func main() {

	//Sarama logger
	sarama.Logger = log.New(outputWriter, "[saramaLog]", log.Ltime)

	//Create a Kafka producer
	var brokers = []string{os.Getenv("KAFKAPORT")}
	var err error
	producer, err = kafkapc.CreateKafkaProducer(brokers)
	if err != nil {
		log.Fatal("Failed to connect to Kafka. Error:", err.Error())
	}

	//If a consumer accesses the topic before it is created,
	//a 'missing node' error will be thrown.
	//Hence, ensure that the topic has been created in Kafka queue
	//by sending an 'init' message and waiting for a short 1 sec.
	log.Print("Creating Topic...")
	// doc := make(map[string]int)
	// doc["number"] = 324
	doc := ""
	//Prepare message to be sent to Kafka
	docBytes, err := json.Marshal(doc)
	producer.Input() <- &sarama.ProducerMessage{
		Key:       sarama.StringEncoder("init"),
		Topic:     os.Getenv("TOPICNAME"),
		Value:     sarama.ByteEncoder(docBytes),
		Timestamp: time.Now(),
	}
	time.Sleep(4 * time.Second)
	log.Print(" ...done")

	for ii := 0; ii >= -10; ii-- {
		// doc := make(map[string]int)
		// doc["number"] = ii
		doc := "Hihi " + strconv.Itoa(ii)
		//Prepare message to be sent to Kafka
		docBytes, err := json.Marshal(doc)
		msg := &sarama.ProducerMessage{
			Topic:     os.Getenv("TOPICNAME"),
			Value:     sarama.ByteEncoder(docBytes),
			Timestamp: time.Now(),
		}
		if err == nil {
			//Send message into Kafka queue
			producer.Input() <- msg
			fmt.Println(doc, "======", msg)
		} else {
			fmt.Println("WARNING ---------------->>>>>>>>", msg)
		}
	}
	fmt.Println("Completed")
}

func checkError(err error) bool {
	if err != nil {
		fmt.Fprintln(outputWriter, err.Error())
		return true
	}
	return false
}
