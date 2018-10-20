import os
# from kafkapc_python import Producer
# from pykafka import KafkaClient
from kafka import KafkaProducer
import json

class Producer(KafkaProducer):
    def __init__(self, kafkaPort):
        KafkaProducer.__init__(
            self,
            bootstrap_servers=[kafkaPort],
            value_serializer=lambda x: json.dumps(x).encode(),
            key_serializer=str.encode
            )

def main():

    producer = Producer(
        os.environ['KAFKAPORT']
        )
    for e in range(1000):
        data = {'number': e}
        producer.send(os.environ['TOPICNAME'], value=data)
        print(data)

    # topics_to_produce_to = [os.environ['TOPICNAME']]
    # topic_producers = {topic.name: topic.get_producer() for topic in topics_to_produce_to}
    # for destination_topic, message in consumed_messages:
    #     topic_producers[destination_topic.name].produce(message)

    # client = KafkaClient(os.environ['KAFKAPORT'])
    # topic = client.topics[os.environ['TOPICNAME']]
    # print(client.topics)

    # msg = {'value':2}
    # my_custom_serializer(msg)
    # print(msg)

    # with topic.get_producer(serializer=my_custom_serializer) as producer:
    #     for e in range(10):
    #         data = e
    #         # producer.send(os.environ['TOPICNAME'], value=data)
    #         print(data)

    # return

if __name__ == "__main__":
    #Call main function
    main()
