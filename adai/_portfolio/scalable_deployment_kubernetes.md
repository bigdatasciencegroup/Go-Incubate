---
title: "Scalable Deployment in Kubernetes"
excerpt: "Video Analytics, TensorFlow Serving, Golang, Kafka, Docker, REST"
header:
  teaser: /assets/images/scalable_deployment_kubernetes_01.jpg
feature_row:
  - image_path: /assets/images/scalable_deployment_kubernetes_03.jpg
  - image_path: /assets/images/scalable_deployment_kubernetes_04.jpg
---

{% include feature_row %}

{% include toc %}

## Introduction

In this project, we will develop a real-time video analytic pipeline (i.e., machine learning on a video stream). Video feed will be streamed from IP cameras into a Kafka queue, processed by a machine learning model, predictions are imprinted onto the frames, processed frames are queued again into another Kafka topic, and then streamed to web for display. Microservices architecture is adopted, where all code is containerized using Dockers and orchestrated using Kubernetes.

The following tools will be used in this project:
+ Kubernetes
+ Docker
+ Golang
+ Kafka & Zookeeper
+ REST
+ TensorFlow Serving API

In this project, the video analytics will focus on a frame-by-frame object classification task, similar to the ImageNet competition. In fact, we will deploy a TensorFlow Serving saved model which is originally trained on ImageNet database.

The complete system design and data flow of this project is illustrated by the following image.

{% capture fig_pipeline %}
![pipeline](/assets/images/scalable_deployment_kubernetes_02.jpg){:height="100%" width="100%" .align-center}
{% endcapture %}

<figure id="pipeline">
  {{ fig_pipeline | markdownify | remove: "<p>" | remove: "</p>" }}
  <figcaption>Scalable microservices architecture for video analytics pipeline.</figcaption>
</figure>

## Learning Outcome

Find the source code in the [repository](https://github.com/Adaickalavan/Scalable-Deployment-Kubernetes).

At the end of this project, we should be able to:
+ Write deployment.yml files in Kubernetes and docker-compose.yml files in Docker
+ Run containerized Confluent Zookeeper and Confluent Kafka in Kubernetes
+ Always retrieve the latest message in the Kafka topic-partition queue
+ Use GoCV (Golang client for OpenCV) to stream video and to manipulate images in Golang
+ Use TensorFlow Serving, with REST API, to deploy machine learning models
+ Form MJPEG from JPEG images and broadcast the video to web
+ Recover and restart a goroutine which panicked

## Project Structure

The project structure is as follows:

```text
Scalable-Deployment-Kubernetes                     # Main folder
├── goconsumer                                     # Process video frames using machine learning models
│   ├── assets                                     #
│   │   └── imagenetLabels.json                    # Labels for imagenet classes
│   ├── vendor                                     # Dependency files
│   │   ├── confluentkafkago                       #
│   │   │   ├── consumer.go                        # Confluent-Kafka-Go consumer wrapper
│   │   │   └── producer.go                        # Confluent-Kafka-Go producer wrapper
│   │   ├── models                                 #
│   │   │   ├── handler.go                         # Handler for all TensorFlow models
│   │   │   └── imagenet.go                        # Handler for imagenet TensorFlow model
│   │   └── github.com                             #
│   │       └── pkg                                #
│   │           └── profile                        # Golang code profiler (cpu and memory)
│   ├── deployment.yml                             # Kubernetes deployment
│   ├── docker-compose.yml                         # Docker deployment
│   ├── dockerfile                                 # To create Docker container
│   ├── main.go                                    #
│   └── message.go                                 # Processes Kafka messages by calling the appropriate TensorFlow models.handler function
├── goproducer                                     # Streams video from IP camera into Kafka
│   ├── vendor                                     # Dependency files
│   │   └── confluentkafkago                       #
│   │       └── producer.go                        # Confluent-Kafka-Go producer wrapper
│   ├── deployment.yml                             # Kubernetes deployment
│   ├── docker-compose.yml                         # Docker deployment
│   ├── dockerfile                                 # To create Docker container
│   └── main.go                                    #
├── govideo                                        # Publishes video to web
│   ├── vendor                                     # Dependency files
│   │   ├── confluentkafkago                       #
│   │   │   └── consumer.go                        # Confluent-Kafka-Go consumer wrapper
│   │   └── mjpeg                                  #
│   │       └── stream.go                          # Streams JPEG pictures to web to form motion JPEG (MJPEG)
│   ├── deployment.yml                             # Kubernetes deployment
│   ├── docker-compose.yml                         # Docker deployment
│   ├── dockerfile                                 # To create Docker container
│   └── main.go                                    #
├── tfserving                                      # TensorFlow Serving API
│   ├── resnet                                     # Resnet TensorFlow model
│   │   └── 1538687457                             # TensorFlow model version number
│   │       ├── assets                             #
│   │       │   └── imagenetLabels.json            # Labels for imagenet classes
│   │       ├── variables                          # Variables of TensorFlow saved model
│   │       │   ├── variables.data-00000-of-00001  #
│   │       │   └── variables.index                #
│   │       └── saved_model.pb                     # TensorFlow saved model
│   ├── deployment.yml                             # Kubernetes deployment
│   ├── docker-compose.yml                         # Docker deployment
│   └── dockerfile                                 # To create Docker container
└── zookeeper                                      # Zookeeper and Kafka
    ├── deployment.yml                             # Kubernetes deployment
    └── docker-compose.yml                         # Docker deployment
```

## Instructions

This project can be run either in Kubernetes cluster using the provided `deployment.yml` files or in Docker using the provided `docker-compose.yml` files. The pipeline consists of 5 subsystems, namely, (i) Zookeeper & Kafka, (ii) GoProducer, (iii) GoConsumer, (iv) TFServing, and (v) GoVideo. Commands to be executed to run each of the subsystem is provided below for both Kubernetes and Docker environments.

<style type="text/css">
.tg {border-collapse:collapse;border-spacing:0;border-color:#aaa;margin:0px auto;}
.tg td{font-family:Arial, sans-serif;font-size:14px;padding:10px 10px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;border-color:#aaa;color:#333;background-color:#fff;width: 33%}
.tg th{font-family:Arial, sans-serif;font-size:14px;font-weight:bold;padding:10px 5px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;border-color:#aaa;color:#fff;background-color:#f38630;}
.tg .tg-7d57{background-color:#FCFBE3;border-color:inherit;text-align:left;vertical-align:top}
.tg .tg-kwiq{color:#000000;border-color:inherit;text-align:center;vertical-align:top}
.tg .tg-0pky{border-color:inherit;text-align:left;vertical-align:top}
@media screen and (max-width: 767px) {.tg {width: auto !important;}.tg col {width: equal !important;}.tg-wrap {overflow-x: auto;-webkit-overflow-scrolling: touch;margin: auto 0px;}}
</style>
<div class="tg-wrap"><table class="tg">
  <tr>
    <th class="tg-kwiq">Steps</th>
    <th class="tg-kwiq">Kubernetes</th>
    <th class="tg-kwiq">Docker</th>
  </tr>
  <tr>
    <td class="tg-7d57">Start Kubernetes cluster
      <ol style='list-style-type:disc'>
        <li>Sample command given here starts a minikube cluster</li>
		    <li>Run the <code>eval</code> statement at the beginning of each shell. It enables the use of local Docker daemon to build Docker images  within Kubernetes cluster.</li>
	    </ol>
    </td>
    <td class="tg-7d57"><code>$ minikube start</code><br><code>$ eval $(minikube docker-env)</code></td>
    <td class="tg-7d57"></td>
  </tr>
  <tr>
    <td class="tg-0pky">Start Zookeeper and Kafka</td>
    <td class="tg-0pky"><code>$ cd ~/Scalable-Deployment-Kubernetes/zookeeper</code><br><code>$ kubectl apply -f deployment.yml</code></td>
    <td class="tg-0pky"><code>$ cd ~/Scalable-Deployment-Kubernetes/zookeeper</code><br><code>$ docker-compose up</code></td>
  </tr>
  <tr>
    <td class="tg-7d57">Start GoProducer<br>
      <ol style='list-style-type:disc'>
		    <li>Streams video from IP camera into Kafka queue</li>
	    </ol>
    </td>
    <td class="tg-7d57"><code>$ cd ~/Scalable-Deployment-Kubernetes/goproducer</code><br><code>$ docker build -t goproducer .</code><br><code>$ kubectl apply -f deployment.yml</code></td>
    <td class="tg-7d57"><code>$ cd ~/Scalable-Deployment-Kubernetes/goproducer</code><br><code>$ docker build -t goproducer .</code><br><code>$ docker-compose up</code></td>
  </tr>
  <tr>
    <td class="tg-0pky">Start GoConsumer<br>
      <ol style='list-style-type:disc'>
		    <li>Classifies objects in each frame using Resnet machine learning model</li>
	    </ol>
    </td>
    <td class="tg-0pky"><code>$ cd ~/Scalable-Deployment-Kubernetes/goconsumer</code><br><code>$ docker build -t goconsumer .</code><br><code>$ kubectl apply -f deployment.yml</code></td>
    <td class="tg-0pky"><code>$ cd ~/Scalable-Deployment-Kubernetes/goconsumer</code><br><code>$ docker build -t goconsumer .</code><br><code>$ docker-compose up</code></td>
  </tr>
  <tr>
    <td class="tg-7d57">Start TensorFlow Serving<br>
      <ol style='list-style-type:disc'>
		    <li>Answers REST queries to port 8501</li>
	    </ol>
    </td>
    <td class="tg-7d57"><code>$ cd ~/Scalable-Deployment-Kubernetes/tfserving</code><br><code>$ docker build -t tfserving .</code><br><code>$ kubectl apply -f deployment.yml</code></td>
    <td class="tg-7d57"><code>$ cd ~/Scalable-Deployment-Kubernetes/tfserving</code><br><code>$ docker build -t tfserving .</code><br><code>$ docker-compose up</code></td>
  </tr>
  <tr>
    <td class="tg-0pky">Start GoVideo<br>
      <ol style='list-style-type:disc'>
		    <li>Broadcasts video to web</li>
	    </ol>
    </td>
    <td class="tg-0pky"><code>$ cd ~/Scalable-Deployment-Kubernetes/govideo</code><br><code>$ docker build -t govideo .</code><br><code>$ kubectl apply -f deployment.yml</code></td>
    <td class="tg-0pky"><code>$ cd ~/Scalable-Deployment-Kubernetes/govideo</code><br><code>$ docker build -t govideo .</code><br><code>$ docker-compose up</code></td>
  </tr>
</table></div>
<br>
To terminate the microservices
+ in Kubernetes, execute `kubectl delete -f deployment.yml` in the respective folders of each susbsytem.
+ in Docker, simply close the terminal used to run the subsystem.

## System Design

Each subsystem of the [pipeline](#pipeline), is further explored in the following sections. For brevity, only deployment in Kubernetes (i.e., `deployment.yml`) is discussed below, since the discussions are similar for deployment in Docker environment using its counterpart `docker-compose.yml` files.

For beginners in Kubernetes, please see my [post](/guides/guide-to-kubernetes/) on guide to quickly mastering Kubernetes.

### Zookeeper and Kafka

<u>deployment.yml</u>
+ Confluent Zookeeper and Kafka images are used.
+ `kind` of Zookeeper and Kafka are set to `StatefulSet` as they should have persistent states
+ A single Kafka broker with 5 ports (19091, 19092, 19093, 19094, 19095) is used, although this could be easily scaled up to multiple Kafka brokers.
+ IP address of the Kafka pod is obtained by declaring `MY_POD_IP` in the environment variable as follows:
    ```yml
    env
      - name: MY_POD_IP
        valueFrom:
            fieldRef:
              fieldPath: status.podIP
    ```
+ The variable `MY_POD_IP` can be referenced elsewhere in the yml file as `$(MY_POD_IP)`. However, such variables must be declared before being used, for example in `KAFKA_ADVERTISED_LISTENERS`.
+ The variables `KAFKA_REPLICA_FETCH_MAX_BYTES` and `KAFKA_MESSAGE_MAX_BYTES` are set to 100MB to handle larger video frame sizes.
+ To save hard disk memory, we opt to delete logs older than 1 minute by setting `KAFKA_LOG_CLEANUP_POLICY` and `KAFKA_LOG_RETENTION_MIUTES`.
+ Here, `KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR` is set to 1 as there is only one Kafka broker.
+ `Service`'s are used to communicate between pods in Kubernetes

### GoProducer

<u>deployment.yml</u>
+ Environment variable `VIDEOLINK` states the IP address of the IP camera
+ Environment variable `FRAMEINTERVAL` determines the frames-per-second (fps) of the video. For example, 42ms translates into 24 fps.

<u>main.go</u>
+ Streams video from `VIDEOLINK` at `FRAMEINTERVAL` interval and writes them into Kafka topic `TOPICNAME`. 

<u>confluentkafkago.NewProducer()</u>
+ Provides a high level wrapper code for creation of Confluent Kafka producers.
+ `message.max.bytes` is set to 100MB to support large sized messages such as images with high quality or large size

### GoConsumer

<u>deployment.yml</u>
+ `MODELURLS` and `LABELURLS` contain addresses to TensorFlowServing saved model and class labels, respectively.
+ Multiple machine learning models can be applied to the same video.
+ Here, we have simply duplicated the same machine learning model, Resnet, twice to mimic use of different machine learning models.

<u>main.go</u>
+ Consumes video from Kafka topic `TOPICNAMEIN`, and handles valid messages using `main.message()` function.
+ Consists of three goroutines, namely, (i) to consume messages from `TOPICNAMEIN`, (ii) to query TensorFlow model, and (ii) to rewrite processed frames into `TOPICNAMEOUT`.

<u>main.message()</u>  
+ Frame is extracted from received Kafka message
+ Frame is fed to each of the TensorFlow Serving saved models
+ The class predictions from each of the TensorFlow Serving is retrieved and imprinted onto the video.

<u>models.*imagenet.Predict()</u>
+ Encodes `gocv.Mat` type to `JPEG` type, packs the images into a JSON structure and makes a POST request to the TensorFlow Serving pod at port 8501. 
+ Returns the predicted class
+ In the event the goroutine panicks, it is recovered and restarted via the defer function:
    ```go
    defer func() {
      if r := recover(); r != nil {
        log.Println("models.*imagenet.Predict():PANICKED AND RESTARTING")
        log.Println("Panic:", r)
        go imn.Predict()
      }
    }()
    ```

<u>confluentkafkago.NewConsumer()</u>
+ Provides a high level wrapper code for creation of Confluent Kafka consumers.

<u>confluentkafkago.LatestOffset()</u>
+ LatestOffset() resets consumer offset to the latest message in the topic-partition queue, if the difference between the low and high watermark is more than the desired diff. 
+ Setting `diff=0`, Kafka consumer will always consume the latest message.
    ```go
    func LatestOffset(c *kafka.Consumer, diff int) error {

      // Record the current topic-partition assignments
      tpSlice, err := c.Assignment()
      if err != nil {
        return err
      }

      //Obtain the last message offset for all topic-partition
      for index, tp := range tpSlice {
        low, high, err := c.QueryWatermarkOffsets(*(tp.Topic), tp.Partition, 100)
        if err != nil {
          return err
        }
        if high-low < int64(diff) {
          return errors.New("Offset difference between Low and High is smaller than diff")
        }
        tpSlice[index].Offset = kafka.Offset(high)
      }

      //Consume the last message in topic-partition
      err = c.Assign(tpSlice)
      if err != nil {
        return err
      }
      return nil
    }
    ```

<u>assets/imagenetLabels.json</u>
+ Contains class labels in JSON format from ImageNet competition. 

### TFServing

<u>dockerfile</u>
+ `tensorflow/serving` is used as the base image
+ Our TensorFlow saved model is copied from `./resnet` folder into the image's `/model/tfModel` folder
+ Since a TensorFlow Serving can hold multiple models simultaneously, the desired model to be used is specified by setting the environment variable `MODEL_NAME` to that of the desired model's folder name in the image. In our case, desired model's folder name would be `tfmodel`.

<u>deployment.yml</u>
+ The `deployment.yml` runs two `tfserving` replicas in a load balancing manner to serve REST requests from clients at port 8501 and gRPC requests at port 8500.

<u>resnet</u>
+ Please refer to the official [guide](https://www.tensorflow.org/tfx/guide/serving) on how to build and use a TensorFlow Serving model.

### GoVideo

<u>deployment.yml</u>
+ Runs one instance of `govideo` and a `govideo-service` service.
+ To view the video output:
  + from outside Kubernetes, go to `<Kubernetes Cluster IP>:<NODEPORT>`. For example, if you are running Minikube at `192.168.99.100`, then go to `192.168.99.100:30163`.
  + from outside Docker, go to `<Docker machine IP>:<NODEPORT>`. For example, if you are running Docker natively, go to `127.0.0.1:30163`. If you are running Docker in a virtual machine at `192.168.99.100`, then go to `192.168.99.100:30163`.

<u>main.go</u>
+ Uses a goroutine to read processed video frames from Kafka topic `TOPICNAME` and calls `mjpeg.*Stream.UpdateJPEG()` to form MJPEG
+ Runs a web server listening at `<POD IP>:<DISPLAYPORT>`, which is mapped to external node (i.e., host machine) at `<Machine IP>:<NODEPORT>` by the `govideo-service` service.

<u>mjpeg.*Stream.UpdateJPEG()</u>
+ Updates JPEG images to form MJPEG, which is then copied into the channels corresponding to each connected web client.

<u>mjpeg.*Stream.ServeHTTP()</u>
+ Provides a `ServeHTTP()` handle pattern to broadcast MJPEG to each connected web client at a rate of 1/`FRAMEINTERVAL`.
+ A sample video output is shown below. It contains the predicted ImageNet object classes.

  ![OutputVideo](/assets/images/scalable_deployment_kubernetes_06.jpg){:height="60%" width="60%" .align-center}