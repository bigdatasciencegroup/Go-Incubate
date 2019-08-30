---
title: "Kafka"
permalink: /docs/kafka/
toc: true
---

## Apache Kafka
+ Apache Kafka is a publish/subscribe messaging system
+ Also known as a distributed commit log
+ Each commit can be time-stamped
+ Messages are queued into Topics
+ Each  Topic may contain multiple Partitions

## Producers
+ Messages may be written to random/specific Partitions
+ Fire-and-forget: Send a message to the server and don’t care if it arrives successfully or not.
+ Synchronous send: Send a message, and wait to see if the send was successful or not.
+ Asynchronous send: Send method with a callback function, which gets triggered when it receives a response from the Kafka broker.

## Consumers
+ Subscribes to one or more topics 
+ No guarantee of message time-ordering across a topic, just within a single partition
+ Commit offsets automatically or at certain intervals
    <br/><img src="/assets/images/wiki/wiki_kafka_01.jpg" width="80%"/>
+ Work as part of a consumer group, where each partition is only consumed by one member
    <br/><img src="/assets/images/wiki/wiki_kafka_02.jpg" width="80%"/>
+ If a consumer fails, the remaining group members will rebalance the partitions
+ Multiple consumer groups may consume same topic independently
    <br/><img src="/assets/images/wiki/wiki_kafka_03.jpg" width="80%"/>
+ Kafka buffers data and allows consumers to operate in asynchronous multirate systems

## Broker
+ A single Kafka server is called a broker.
+ Partitions are replicated in different brokers. Default topic replication factor is 3. Redundancy of messages in case of broker failure.
+ One broker becomes leader for a partition. Leadership changes in case of broker failure.
+ Producers and consumers connect to  the leader
    <br/><img src="/assets/images/wiki/wiki_kafka_04.jpg" width="80%"/>

## Kafka configuration
Several notable settings in Kafka deployment are mentioned below.

+ Producer ConfigMap
    ```go
    "message.max.bytes": 100000000 //100MB. Defaults to 1MB.
    ```
+ Consumer ConfigMap
    ```go
    "group.id": group //Consumer group id 
    "auto.offset.reset": "earliest" //or “latest”  
    ```
+ Kafka Config
    ```bash
    KAFKA_LOG_CLEANUP_POLICY=delete #Cleanup policy for segments beyond the retention window
    KAFKA_LOG_RETENTION_MINUTES=1 #Minutes to keep a log file before deleting. Default 168 hours.
    ```
+ Always retrieve the latest message by resetting the committed offset
<br/><img src="/assets/images/wiki/wiki_kafka_05.jpg" width="80%"/>