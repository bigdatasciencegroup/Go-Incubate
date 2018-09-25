package main

import (
	"document"
	"encoding/json"
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
	data map[string]map[string]document.Word
}

func addFakeData(ds *dataStore) {
	user1 := make(map[string]document.Word)
	user1["Golang"] = document.Word{
		Value:   "Hello",
		Meaning: "Greeting",
	}
	user1["Kafka"] = document.Word{
		Value:   "ByeBye",
		Meaning: "Greeting",
	}
	ds.data = make(map[string]map[string]document.Word)
	ds.data["user1"] = user1
}

func main() {

	ds := &dataStore{}
	addFakeData(ds)

	brokers := []string{os.Getenv("ADVERTISED_HOST") + ":" + os.Getenv("ADVERTISED_PORT")}
	producer, err := kafkasw.CreateKafkaProducer(brokers)
	if err != nil {
		log.Fatal("Failed to connect to Kafka")
	}

	//Ensures that the topic has been created in kafka
	producer.Input() <- &sarama.ProducerMessage{
		Key:       sarama.StringEncoder("init"),
		Topic:     os.Getenv("TOPICNAME"),
		Timestamp: time.Now(),
	}
	log.Println("Creating Topic...")
	time.Sleep(1 * time.Second)

	ConsumerParam := kafkasw.ConsumerParam{
		GroupName: "group.testing",
		Topics:    []string{"TOPICNAME"},
		Zookeeper: []string{os.Getenv("ADVERTISED_HOST") + ":" + os.Getenv("ZOOKEEPER_PORT")},
	}

	go func() {
		kafkasw.ConsumeMessages(ConsumerParam, msgHandler(ds))
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

func msgHandler(ds *dataStore) func(m *sarama.ConsumerMessage) error {
	return func(m *sarama.ConsumerMessage) error {
		// Empty body means it is an init message
		if len(m.Value) == 0 {
			return nil
		}

		word := &document.Word{}
		e := json.Unmarshal(m.Value, word)

		if e != nil {
			return e
		}

		//Simulate processing time
		time.Sleep(1 * time.Second)
		log.Printf("Adding skill %s to user %s", skillMsg.SkillName, skillMsg.ProfileID)

		score := skillScore{
			SkillName:  skillMsg.SkillName,
			Score:      rand.Float32() * 100,
			LastScored: time.Now(),
		}

		ds.WriteData(skillMsg.ProfileID, score)
		return nil
	}
}
