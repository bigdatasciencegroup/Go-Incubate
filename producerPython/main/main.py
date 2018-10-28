# from time import sleep
# import json
# from kafka import KafkaProducer
from pykafka import KafkaClient
import time


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

    client = KafkaClient(hosts="kafka:9092")
    topic = client.topics['my_test']
    with topic.get_sync_producer() as producer:
        for i in range(600):
            producer.produce(('test message ' + str(i ** 2)).encode())
            print('test message ' + str(i ** 2))

    time.sleep(10)

    return

if __name__ == "__main__":
    #Call main function
    main()
