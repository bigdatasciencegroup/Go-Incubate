import os
# from kafkapc_python import Producer
from pykafka import KafkaClient
import json

def my_custom_serializer(msg):
    """"Transforms all msg fields into raw bytes"""
    msg.value = json.dumps(msg.value).encode()

def my_custom_deserializer(msg):
    """"Transforms msg fields into python objects"""
    msg.value = json.loads(msg.value.decode())


topic.get_producer(deserializer=my_custom_serializer)
topic.get_consumer(serializer=my_custom_deserializer)

def main():

    # producer = Producer(
    #     os.environ['KAFKAPORT']
    #     )
    # for e in range(1000):
    #     data = {'number': e}
    #     producer.send(os.environ['TOPICNAME'], value=data)
    #     print(data)

    # topics_to_produce_to = [os.environ['TOPICNAME']]
    # topic_producers = {topic.name: topic.get_producer() for topic in topics_to_produce_to}
    # for destination_topic, message in consumed_messages:
    #     topic_producers[destination_topic.name].produce(message)




    client = KafkaClient(os.environ['KAFKAPORT'])
    topic = client.topics[os.environ['TOPICNAME']]
    print(client.topics)

    with topic.get_producer() as producer:
        for e in range(10):
            data = {'number': e}
            producer.produce(data)
            # producer.send(os.environ['TOPICNAME'], value=data)
            print(data)

    return

if __name__ == "__main__":
    #Call main function
    main()
