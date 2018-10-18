package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/adaickalavan/Go-Rest-Kafka-Mongo/document"
	"github.com/adaickalavan/Go-Rest-Kafka-Mongo/kafkapc"
	"github.com/joho/godotenv"
)

//Hooks that may be overridden for testing
var inputReader io.Reader = os.Stdin
var outputWriter io.Writer = os.Stdout

func init() {
	//Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

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
	producer.Input() <- &sarama.ProducerMessage{
		Key:       sarama.StringEncoder("init"),
		Topic:     os.Getenv("TOPICNAME"),
		Timestamp: time.Now(),
	}
	time.Sleep(1 * time.Second)
	log.Print(" ...done")

	for ii := 0; ii <= 1000; ii++ {
		matrix := &document.Matrix{}
		matrix.Value = float64(ii)
		//Prepare message to be sent to Kafka
		docBytes, _ := json.Marshal(*matrix)
		msg := &sarama.ProducerMessage{
			Topic:     os.Getenv("TOPICNAME"),
			Value:     sarama.ByteEncoder(docBytes),
			Timestamp: time.Now(),
		}

		//Send message into Kafka queue
		producer.Input() <- msg
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
