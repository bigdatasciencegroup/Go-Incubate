package main

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/Shopify/sarama"
	"github.com/adaickalavan/Go-Rest-Kafka-Mongo/database"
	"github.com/adaickalavan/Go-Rest-Kafka-Mongo/document"
	"github.com/adaickalavan/Go-Rest-Kafka-Mongo/kafkapc"
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

//Instantiate a database dictionary
var dictionary = database.Dictionary{}

func main() {

	//Connect to database
	dictionary.Server = os.Getenv("MONGOPORT")
	dictionary.DatabaseName = os.Getenv("DATABASENAME")
	dictionary.Session = dictionary.Connect()
	//Ensure database index is unique
	dictionary.EnsureIndex([]string{"value"})

	//Sarama logger
	sarama.Logger = log.New(outputWriter, "[saramaLog]", log.Ltime)

	// Set up the Kafka consumer parameter
	ConsumerParam := kafkapc.ConsumerParam{
		GroupName: "databaseWriter",
		Topics:    []string{os.Getenv("TOPICNAME")},
		Zookeeper: []string{os.Getenv("ZOOKEEPERPORT")},
	}

	// Run the consumer
	kafkapc.ConsumeMessages(ConsumerParam, msgHandler(&dictionary))

}

//Consumer message handler
func msgHandler(dictionary *database.Dictionary) func(m *sarama.ConsumerMessage) error {
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

		//Write data into database
		err = dictionary.Insert(*word)
		switch {
		case mgo.IsDup(err):
			log.Println("Key has been duplicated !!!", err.Error())
		case err != nil:
			log.Println("Other error inside msg hnadle", err.Error())
		}

		return nil
	}
}
