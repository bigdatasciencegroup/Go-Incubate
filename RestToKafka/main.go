package main

import (
	"database"
	"document"
	"encoding/json"
	"fmt"
	"io"
	"kafkasw"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/joho/godotenv"
	mgo "gopkg.in/mgo.v2"
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

//Results is data store
type dataStore map[string]document.Word

var ds dataStore
var producer sarama.AsyncProducer
var dictionary = database.Dictionary{}

func main() {

	// ds = make(dataStore)

	//Connect to database
	dictionary.Session = dictionary.Connect()
	//Ensure database index is unique
	dictionary.EnsureIndex([]string{"value"})

	//Create a Kafka producer
	var brokers = []string{os.Getenv("SPEC_KAFKA_PORT")}
	var err error
	producer, err = kafkasw.CreateKafkaProducer(brokers)
	if err != nil {
		log.Fatal("Failed to connect to Kafka. Error:", err.Error())
	}

	// log.Println("Creating Topic...")
	// time.Sleep(1 * time.Second)

	//Create Kafka consumer
	ConsumerParam := kafkasw.ConsumerParam{
		GroupName: "databaseWriter",
		Topics:    []string{os.Getenv("TOPICNAMEPOST")},
		Zookeeper: []string{os.Getenv("SPEC_ZOOKEEPER_PORT")},
	}
	go func() {
		kafkasw.ConsumeMessages(ConsumerParam, msgHandler(&ds))
	}()

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

//Consumer message handler
func msgHandler(ds *dataStore) func(m *sarama.ConsumerMessage) error {
	return func(m *sarama.ConsumerMessage) error {
		// Empty body means it is an init message
		if len(m.Value) == 0 {
			return nil
		}

		//Read message into 'word' struct
		word := &document.Word{}
		err := json.Unmarshal(m.Value, word)
		if err != nil {
			return err
		}

		// //Write data into database
		// (*ds)[word.Value] = *word
		// fmt.Println(ds)

		//Write data into database
		err = dictionary.Insert(*word)
		switch {
		case mgo.IsDup(err):
			log.Println("Key has been duplicated !!! --", err.Error())
		case err != nil:
			log.Println("Other error inside msg hnadle", err.Error())
		}

		return nil
	}
}
