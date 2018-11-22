package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/Shopify/sarama"
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

func main() {

	//Sarama logger
	sarama.Logger = log.New(outputWriter, "[saramaLog]", log.Ltime)

	// Set up the Kafka consumer parameter
	ConsumerParam := kafkapc.ConsumerParam{
		GroupName: os.Getenv("CONSUMERGROUP"),
		Topics:    []string{os.Getenv("TOPICNAME")},
		Zookeeper: []string{os.Getenv("ZOOKEEPERPORT")},
	}

	// Run the consumer
	kafkapc.ConsumeMessages(ConsumerParam, msgHandler())
}

//Result represents the Kafka queue message format
type Result struct {
	T        time.Time `json:"t"`
	Pix      []uint8   `json:"pix"`
	Channels int       `json:"channels"`
	Rows     int       `json:"rows"`
	Cols     int       `json:"cols"`
	Stride   int       `json:"stride"`
}

//Consumer message handler
func msgHandler() func(m *sarama.ConsumerMessage) error {
	return func(m *sarama.ConsumerMessage) error {
		// Empty body means it is an init message
		if len(m.Value) == 0 {
			return nil
		}

		//Read message into 'doc' struct
		doc := Result{}

		err := json.Unmarshal(m.Value, &doc)
		if err != nil {
			return err
		}
		fmt.Println("Doc received:", doc)

		return nil
	}
}
