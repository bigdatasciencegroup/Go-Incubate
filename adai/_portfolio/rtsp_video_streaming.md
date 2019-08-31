---
title: "RTSP Video, Kafka, and Microservices"
excerpt: "Data pipeline: Golang, Python, Docker internal/external networking"
header:
  teaser: assets/images/rtsp_video_streaming_01.jpg
---

<!-- {% include toc %} -->

## Info

A real time streaming protocol ([RTSP](https://en.wikipedia.org/wiki/Real_Time_Streaming_Protocol)) video is streamed from a website using [OpenCV](https://opencv.org/) into a Kafka topic and consumed by a signal processing application. This project serves to highlight and demonstrate various key data engineering concepts. The data pipeline is as follows:

+ `GoProducerRTSP` module: An RTSP video is streamed from a third party website using dockerized [GoCV](https://gocv.io/) (a golang wrapper of openCV) code written in Golang. The video is enqueued in a dockerized Kafka topic. A multistage image build is performed to minimize runtime image size. RTSP is prevalent in security cameras and commercial camera systems.
+ Zookeeper and Kafka modules: Zookeeper and Kafka container instances are created from [Confluent](https://docs.confluent.io/current/installation/docker/docs/image-reference.html) docker images. Besides message passing, Kafka is used for inter-language communication between Golang code and Python code, in this project.
+ `PyConsumerRTSP` module: The video is consumed from the Kafka topic by a Python code running in the host machine (i.e., outside of Docker). The fetched video frames are displayed using OpenCV. Some simple computation is performed on each video frame and results are printed to screen. This simple signal processing code serves as a placeholder for the real signal processing code later.
+ `PyConsumerRTSP2` module: This is a duplicate of `PyConsumerRTSP` module. The consumers of `PyConsumerRTSP` and `PyConsumerRTSP2` belong to two different consumer groups but consume from the same Kafka topic. Running them simultaneously, demonstrates the scalability of the data pipeline.

## Learning Outcome

The [repository](https://github.com/Adaickalavan/RTSP-Video-Streaming) contains the following:

+ Docker-compose file for Zookeeper and Kafka services
+ Go code
  + Dockerized Zookeeper and Kafka images
  + Dockerized Kafka producer with GoCV (openCV) library
+ Python code
  + Kafka consumer for RTSP video
  + OpenCV and dummy signal processing code

At the end of this project, we should be able to:
+ manage Go dependencies using `dep`
+ stream RTSP video
+ use dockerized Zookeeper and Kafka
+ dockerize Golang code
+ perform internal/external Docker networking
+ use dockerized GoCV (i.e., Golang client for OpenCV)
+ use of Sarama library (Golang client for Kafka) and Kafka-Python library (Python client for Kafka)
+ perform docker multistage build to reduce image size
+ use Kafka for inter-language communication (between Golang and Python)
+ building scalable and fault tolerant data pipeline with multiple producers and consumers

## Project Structure

The project structure is as follows:

```text
DataPipeline                # Main folder
├── goproducerrtsp          # To grab frames from RTSP video website and to enqueue in Kafka
│   ├── vendor              # Dependency vendor files generated using 'dep ensure'
│   │   └── ...             #
│   ├── .env                # Environment variables
│   ├── Docker-compose.yml  # To instantiate Docker container
│   ├── Dockerfile          # To build Docker image
│   ├── Gopkg.lock          # Dependency version file generated using 'dep ensure'
│   ├── Gopkg.toml          # Dependency version file generated using 'dep ensure'
│   └── main.go             # Go code producing to Kafka
├── pythonconsumerrtsp      # To consume from Kafka in Python, outside Docker environment
│   ├── dataprocessing      # Template folder for signal processing
│   │   ├── __init__.py     # Package file
│   │   └── alg.py          # Dummy signal processing object: Computes average pixel value
│   ├── kafkapc_python      # Kafka producer and consumer using kafka-python library
│   │   └── __init__.py     #
│   ├── message             # Kafka message handling function
│   │   └── __init__.py     #
│   ├── .env                # Environment variables
│   ├── Docker-compose.yml  # To instantiate Docker container
│   ├── Dockerfile          # To build Docker image
│   ├── main.py             # Python code consuming from Kafka
│   └── requirements.txt    # Imported libraries in the python code
├── pythonconsumerrtsp2     # Duplicate of `pythonconsumerrtsp` to illustrate code scalability
│   ├── ...                 #
│       .                   #
│       .                   #
│       .                   #
│   └── ...                 #
└── zookeeper               # Zookeeper and Kafka
    └── Docker-compose.yml  # To instantiate Docker container for Zookeeper and Kafka
```

## System Design

{: .notice--success}
Further description of the system design will be provided if requested by readers.

The complete system design and data flow of this project is illustrated by the following image.

![pipeline](/assets/images/pipeline_01.jpg){:height="100%" width="100%" .align-center}

Each component of the system will be explored in detail in the following sections.

### Zookeeper and Kafka

Firstly, the Zookeeper and Kafka services are started as containers using docker images from Confluent. Below is the `Docker-compose.yml` file.

```yml
version: '3'
services:

  #Zookeeper
  zookeeper:
    image: confluentinc/cp-zookeeper
    container_name: zookeeper
    restart: always #Ensures that intermittent failures in the Docker environment do not result in unnecessary failures of the service.
    networks:
      - dockerNet
    ports:
      - 2181:2181 #Zookeeper port - access port for external communication
    environment:
      - ZOOKEEPER_SERVER_ID=1
      - ZOOKEEPER_CLIENT_PORT=2181
      - ZOOKEEPER_TICK_TIME=2000

  #Kafka
  kafka:
    image: confluentinc/cp-kafka
    container_name: kafka
    restart: always #Ensures that intermittent failures in the Docker environment do not result in unnecessary failures of the service.
    depends_on:
      - zookeeper
    networks:
      - dockerNet
    ports:
      - 9092:9092 #Kafka broker - access port for external communication
    environment:
      - KAFKA_BROKER_ID=1
      - KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181
      # - KAFKA_ADVERTISED_LISTENERS=PLAINTEXT_HOST://localhost:9092, PLAINTEXT://kafka:29092 #For native Docker (e.g., in Windows 10)
      - KAFKA_ADVERTISED_LISTENERS=PLAINTEXT_HOST://192.168.99.100:9092, PLAINTEXT://kafka:29092 #For Docker Tool (e.g., in Windows 7)
      - KAFKA_LISTENER_SECURITY_PROTOCOL_MAP=PLAINTEXT:PLAINTEXT, PLAINTEXT_HOST:PLAINTEXT
      - KAFKA_INTER_BROKER_LISTENER_NAME=PLAINTEXT
      - KAFKA_JMX_PORT=1099
      - KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1

networks:
  dockerNet:
    driver: bridge
```

Containers are connected via an internal Docker network named `dockerNet` to facilitate internal communication among containers. Containers can access each other by their `container_name`.

Zookeeper port number is `2181` and it is published to localhost making it accessible to programs from outside the Docker environment.

The `KAFKA_ADVERTISED_LISTENERS` variable is set to `192.168.99.100:9092` and `kafka:29092`. This makes Kafka accessible from inside the Docker through `kafka:29092` and also from outside the Docker through `192.168.99.100:9092` by advertising its location on the host machine.

Remember to use your Docker machine IP (e.g., `192.168.99.100` as shown in above code) if you are using Docker Tool (or a VM). Otherwise, if you are using native Docker, please replace the IP addresses with `localhost` as shown commented in the above code.

Issue the following Docker commands in the working directory of the Docker-compose.yml file:

+ To run Zookeeper and Kafka containers
  + `docker-compose up`
+ To tear down all containers and stored volume at the end of operation
  + `docker-compose down -v`
  + `docker system prune`
  + `docker volume prune`
  + `docker network prune`

### GoProducerRTSP

The Kafka producer code, in Golang, to stream RTSP video into Kafka topic `timeseries_1` is shown below. [Sarama](https://github.com/Shopify/sarama) library is used as the Golang client for Kafka producer. A further wrapper for Golang producer (and consumer) built on top of Sarama and [wvanbergen](https://github.com/wvanbergen/kafka) libraries is provided for ease of use in my [kafkapc](https://github.com/Adaickalavan/kafkapc) package.

```go
package main

import (
	"encoding/json"
	"fmt"
	"image"
	"io"
	"log"
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/adaickalavan/kafkapc"
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
	if err != nil {
		panic("Failed to connect to Kafka. Error: " + err.Error())
	}
	//Close producer to flush(i.e., push) all batched messages into Kafka queue
	defer func() { producer.Close() }()

	// Capture video
	webcam, err := gocv.OpenVideoCapture(os.Getenv("RTSPLINK"))
	if err != nil {
		panic("Error in opening webcam: " + err.Error())
	}

	// Stream images from RTSP to Kafka message queue
	frame := gocv.NewMat()
	for {
		if !webcam.Read(&frame) {
			continue
		}

		// Type assert frame into RGBA image
		imgInterface, err := frame.ToImage()
		if err != nil {
			panic(err.Error())
		}
		img, ok := imgInterface.(*image.RGBA)
		if !ok {
			panic("Type assertion of pic (type image.Image interface) to type image.RGBA failed")
		}

		//Form the struct to be sent to Kafka message queue
		doc := Result{
			Pix:      img.Pix,
			Channels: frame.Channels(),
			Rows:     frame.Rows(),
			Cols:     frame.Cols(),
		}

		//Prepare message to be sent to Kafka
		docBytes, err := json.Marshal(doc)
		if err != nil {
			log.Fatal("Json marshalling error. Error:", err.Error())
		}
		msg := &sarama.ProducerMessage{
			Topic:     os.Getenv("TOPICNAME"),
			Value:     sarama.ByteEncoder(docBytes),
			Timestamp: time.Now(),
		}
		//Send message into Kafka queue
		producer.Input() <- msg

		//Print time of receiving each image to show the code is running
		fmt.Fprintf(outputWriter, "---->>>> %v\n", time.Now())
	}
}

//Result represents the Kafka queue message format
type Result struct {
	Pix      []byte `json:"pix"`
	Channels int    `json:"channels"`
	Rows     int    `json:"rows"`
	Cols     int    `json:"cols"`
}
```

Video frames captured by GoCV are of type `gocv.Mat` struct. It is then converted to type `image.Image` interface and type asserted into an RGBA image. The pixel values along with other image information are written into a struct and sent to Kafka topic.

[Dep](https://golang.github.io/dep/) is used for Golang dependency management. To initialize, issue `dep init` command in the project's main working directory. Commands `dep ensure` ensures the project is in sync and `dep ensure -update` updates all the dependencies.

Once `vendor`, `Gopkg.lock`, and `Gopkg.toml` folders have been created by `dep` commands, a docker image is created by executing `docker build -t goproducerrtsp .` on the Dockerfile below.

```yml
FROM denismakogon/gocv-alpine:3.4.2-buildstage as build-stage

#Copy the local package files (from development computer) to the container's workspace (docker image)
COPY . /go/src/app
# Build the app command inside the container
RUN go install app

FROM denismakogon/gocv-alpine:3.4.2-runtime

#Author label
LABEL Author Adaickalavan Meiyappan

#Copy executable binary file
COPY --from=build-stage /go/bin/app /go/bin/app

#Run executable when container starts
ENTRYPOINT ["/go/bin/app"]
```

The `denismakogon/gocv-alpine:3.4.2-buildstage` base image consists of Alpine 3.7, FFMPEG 4.0, Golang 1.10, OpenCV 3.4.2, and OpenCV 3.4.2 contrib packages required by GoCV. A multistage build is performed to minimize runtime image size. Hence, the compiled binary is copied from the build-stage into a basic runtime base image `denismakogon/gocv-alpine:3.4.2-runtime` consisting of Alpine 3.7.

The `goproducerrtsp` service is started by running `docker-compose up` on the following Docker-compose.yml file. To enable internal Docker communication between `goproducerrtsp` and `kafka` containers, set the `KAFKAPORT` address to `kafka:29092`.

```yml
version: '3'

services:
  #Producer to obtain data from RTSP and write into Kafka queue
  goproducerrtsp:
    image: goproducerrtsp
    container_name: goproducerrtsp
    environment:
      - TOPICNAME=timeseries_1
      - KAFKAPORT=kafka:29092
      - RTSPLINK=rtsp://184.72.239.149/vod/mp4:BigBuckBunny_175k.mov
    networks:
      - zookeeper_dockerNet
        # To ensure that the containers in different docker-compose files communicate with each other, we place them on the same network. The complete network name is 'zookeeper_dockerNet'. It is derived by joining the name of the folder from which the network originates (i.e., zookeeper) and the name of the network (i.e., dockerNet).
    privileged: true  

networks:
  zookeeper_dockerNet:
    external: true #Label the other docker-compose network as an external network to the current docker-compose file
```

### PyConsumerRTSP

The video frames are retrieved by a Kafka consumer in Python using the [kafka-python](https://github.com/dpkp/kafka-python) library. A further wrapper for Python consumer (and producer) built on top of kafka-python library is provided for ease of use in my [kafkapc_python](https://github.com/Adaickalavan/RTSP-Video-Streaming/tree/master/pyconsumerrtsp/kafkapc_python) package.

```python
from dotenv import load_dotenv
import kafkapc_python as pc
import os
import cv2
import message
import dataprocessing.alg as alg

def main():
    # Create consumer 
    consumer = pc.Consumer(
        os.getenv('TOPICNAME'),
        os.getenv('KAFKAPORT'),
        os.getenv('CONSUMERGROUP')
        )

    # Prepare openCV window
    print(cv2.__version__)
    windowName = "RTSPvideo1"
    cv2.namedWindow(windowName)
    cv2.resizeWindow(windowName, 240, 160)

    #Instantiate a signal processing model
    model = alg.Model()

    # Start consuming video
    for m in consumer:
        #Read message contents
        val = m.value
        print("Time:",m.timestamp,", Topic:",m.topic)

        #Message handler
        img = message.handler(val)

        #Show image
        cv2.imshow(windowName, img)
        cv2.waitKey(1)

        #Process the message
        model.run(img)

    consumer.close()

    return

if __name__ == "__main__":
    # Load environment variable
    load_dotenv(override=True)
    # Call main function
    main()

```

Environment variables are saved in a `.env` file, as shown below. The [python-dotenv](https://github.com/theskumar/python-dotenv) library is used to load the `.env` file into Python.
```yml
TOPICNAME=timeseries_1
#For native Docker (e.g., in Windows 10)
# KAFKAPORT=localhost:9092 
#For Docker Tool (e.g., in Windows 7)
KAFKAPORT=192.168.99.100:9092 
CONSUMERGROUP=consumerGroup_1
```

The Python code is to be run on the host machine, outside of the Docker environment. Hence, for the Python code to communicate with the `kafka` container living inside the Docker environment, set `KAFKAPORT=192.168.99.100:9092`.

Remember to use your Docker machine IP (e.g., `192.168.99.100` as shown in above code) if you are using Docker Tool (or a VM). Otherwise, if you are using native Docker, please replace the IP addresses with `localhost` as shown commented in the above code.

Slice of bytes from Golang, representing image pixel values, is stored as base-64 string in Kafka topic. The `message.handler()` function converts base-64 string into appropriately shaped 3-channel RGB numpy matrix. Converted numpy matrix is then displayed by [OpenCV-python](https://opencv-python-tutroals.readthedocs.io/) library.

The `dataprocessing.alg` module provides a template to perform any further signal processing steps.

Python project dependencies can be easily managed in a two step process. First, execute `pipreqs` command to generate `requirements.txt` file containing all the required library imports.

```python
pipreqs [options] <path/to/Python/project/folder>

[options]
--force : to overwrite existing file
```

Second, install the dependencies via `pip install -r requirements.txt`.

Although the Python code in this project is meant to be run on the host machine, it may be containerized. Sample [Dockerfile](https://github.com/Adaickalavan/RTSP-Video-Streaming/blob/master/pyconsumerrtsp/Dockerfile) to build the Python image and sample [Docker-compose.yml](https://github.com/Adaickalavan/RTSP-Video-Streaming/blob/master/pyconsumerrtsp/Docker-compose.yml) to instantiate the Python container are provided in the repository.

### PyConsumerRTSP2

The `pyconsumerrtsp2` Python module is simply a duplicate of `pyconsumerrtsp` module. Running both Python modules simultaneously illustrates the scalability and resilience of the Kafka based system. The `pyconsumerrtsp2` module similarly runs on the host machine, outside the Docker environment.

From the `.env` file, we see that `pyconsumerrtsp2` subscribes to the same Kafka topic as `pyconsumerrtsp` but is assigned to a different consumergroup, namely, `consumerGroup_2`.
```yml
TOPICNAME=timeseries_1
#For native Docker (e.g., in Windows 10)
# KAFKAPORT=localhost:9092 
#For Docker Tool (e.g., in Windows 7)
KAFKAPORT=192.168.99.100:9092 
CONSUMERGROUP=consumerGroup_2
```

Although both Python consumers read from the same underlying message queue `timeseries_1`, they work asynchronously by processing data at different speeds and at different times. This enables horizontal scaling of multiple processes on the same data.
