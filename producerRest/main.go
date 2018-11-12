package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
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
	// log.Print("Creating Topic...")
	// producer.Input() <- &sarama.ProducerMessage{
	// 	Key:       sarama.StringEncoder("init"),
	// 	Topic:     os.Getenv("TOPICNAME"),
	// 	Timestamp: time.Now(),
	// }
	// time.Sleep(1 * time.Second)
	// log.Print(" ...done")

	//Run the REST API server
	if err := run(); err != nil {
		log.Fatal(err.Error())
	}

}

//Create and run REST API server
func run() error {
	mux := makeMuxRouter()
	httpAddr := os.Getenv("LISTENINGADDR")
	log.Println("Listening on ", httpAddr)
	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func checkError(err error) bool {
	if err != nil {
		fmt.Fprintln(outputWriter, err.Error())
		return true
	}
	return false
}
