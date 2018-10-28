from time import sleep
import json
from kafka import KafkaProducer

def main():

    # producer = Producer(
    #     os.environ['KAFKAPORT']
    #     )

    producer = KafkaProducer(bootstrap_servers=['kafka:9092'],
                             value_serializer=lambda x: json.dumps(x).encode('utf-8'))

    for e in range(100):
        data = {'number' : e}
        producer.send('numtest', value=data)
        # sleep(5)

    # for e in range(10):
    #     data = {"number": e}    
    #     producer.send(os.environ['TOPICNAME'], key='kill', value=data)
    #     print("Data sent ---> ", data)
    #     # producer.flush()

    # # Block until all async messages are sent
    producer.flush()

    return

if __name__ == "__main__":
    #Call main function
    main()
