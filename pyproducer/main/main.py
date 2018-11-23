import json
import os
import time
from pykafka import KafkaClient
# from kafka import KafkaProducer

def main():

    # producer = Producer(
    #     os.environ['KAFKAPORT']
    #     )

    # producer = KafkaProducer(bootstrap_servers=['kafka:9092'],
    #                          value_serializer=lambda x: json.dumps(x).encode('utf-8'))

    # for e in range(1000):
    #     data = {'number' : e}
    #     producer.send('numtest', value=data)
        # sleep(5)

    # for e in range(10):
    #     data = {"number": e}    
    #     producer.send(os.environ['TOPICNAME'], key='kill', value=data)
    #     print("Data sent ---> ", data)
    #     # producer.flush()

    # # Block until all async messages are sent
    # producer.flush()

    client = KafkaClient(hosts=os.environ['KAFKAPORT'])
    topic = client.topics[os.environ['TOPICNAME']]
    # Using with context, the program waits for the producer 
    # to flush its message queue. producer.produce() ships 
    # messages asynchronously by default.
    with topic.get_sync_producer() as producer:
        for i in range(10):
            data = {"number": i ** 2 + 1} 
            producer.produce(serializer(data))
            print(data)

    return

def serializer(x):
    return json.dumps(x).encode('utf-8')

if __name__ == "__main__":
    #Call main function
    main()
