package main

import (
	"fmt"
	"io"
	"kafkasw"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/Shopify/sarama"
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

	// Seed for fake skill score
	rand.Seed(time.Now().Unix())
}

//Results is data store
type dataStore struct {
	data map[string]map[string]int
}

func addFakeData(ds *dataStore) {
	user1 := make(map[string]int)
	user1["Golang"] = 3

	user1["Kafka"] = 2

	ds.data = make(map[string]map[string]int)
	ds.data["user1"] = user1
}

func main() {

	ds := &dataStore{}
	addFakeData(ds)

	kafkaConn := os.Getenv("ADVERTISED_HOST") + ":" + os.Getenv("ADVERTISED_PORT")
	producer, err := kafkasw.createKafkaProducer(kafkaConn)
	if err != nil {
		log.Fatal("Failed to connect to Kafka")
	}

	//Ensures that the topic has been created in kafka
	producer.Input() <- &sarama.ProducerMessage{
		Key:       sarama.StringEncoder("init"),
		Topic:     topicName,
		Timestamp: time.Now(),
	}

	go func() {
		consumeMessages("127.0.0.1:2181", msgHandler(ds))
	}()

	if err := run(); err != nil {
		log.Fatal(err.Error())
	}

}

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
