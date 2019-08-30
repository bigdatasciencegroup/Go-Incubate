---
title: "Docker Images"
excerpt: "Golang, GoCV, Librdkafka, Confluent-Kafka-Go"
header:
  teaser: assets/images/gopher_docker.jpg
---

{% include toc %}

## Info

Below, I have listed several helpful Docker images which I maintain at [DockerHub](https://hub.docker.com/u/adaickalavan).

## Repository

Automated docker build images and their contents are shown below.

1. **GoKafka-Alpine**

    Repository [link](https://github.com/Adaickalavan/gokafka-alpine)

    + FROM golang:alpine
    + Add
      + Librdkafka
      + Confluent-Kafka-Go

    Output image: adaickalavan/gokafka-alpine

2. **GoCV-Alpine**

    Repository [link](https://github.com/Adaickalavan/gocv-alpine)

    + FROM denismakogon/ffmpeg-alpine:4.0-golang
    + Add
      + OpenCV 4.0.1
      + GUI - highgui providing display functionalities for OpenCV

    Output image: adaickalavan/gocv-alpine
  
3. **GoCVKafka-Alpine**

    Repository [link](https://github.com/Adaickalavan/gocvkafka-alpine)

    + FROM adaickalavan/gocv-alpine
    + Add
      + GoCV
      + Librdkafka
      + Confluent-Kafka-Go

    Output image: adaickalavan/gocvkafka-alpine
