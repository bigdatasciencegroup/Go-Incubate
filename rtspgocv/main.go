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
	"gocv.io/x/gocv"
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
	//Close producer to flush(i.e., push) all batched messages into Kafka queue
	defer func() { producer.Close() }()
	if err != nil {
		log.Fatal("Failed to connect to Kafka. Error:", err.Error())
	}

	// webcam, _ := gocv.VideoCaptureDevice(0)
	// webcam, _ := gocv.VideoCaptureFile("rtsp://127.0.0.1:8554/live.sdp")
	webcam, error := gocv.OpenVideoCapture("rtsp://184.72.239.149/vod/mp4:BigBuckBunny_175k.mov")
	// fmt.Println("Error: ", error)
	// window := gocv.NewWindow("Hello")
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
		// window.IMShow(img)
		// window.WaitKey(1)
		fmt.Printf("%T, %v /n", img.Total(), img.Total())
		doc := Result{T: time.Now()}
		fmt.Println(doc)

		//Prepare message to be sent to Kafka
		docBytes, err := json.Marshal(doc)
		if err != nil {
			log.Fatal("Json marshalling error. Error:", err.Error())
			continue
		}
		msg := &sarama.ProducerMessage{
			Topic:     os.Getenv("TOPICNAME"),
			Value:     sarama.ByteEncoder(docBytes),
			Timestamp: time.Now(),
		}
		//Send message into Kafka queue
		producer.Input() <- msg

		fmt.Println("Yes an image was displayed")
		time.Sleep(1000 * time.Millisecond)
	}
}

//Result is a struct
type Result struct {
	T   time.Time `json:"t"`
	Mat gocv.Mat  `json:"mat"`
}
