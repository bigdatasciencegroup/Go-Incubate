# import json
# from kafka import KafkaConsumer
# from kafka.structs import OffsetAndMetadata, TopicPartition
# from pykafka import KafkaClient
from kafka import KafkaConsumer
# import kazoo
import os

# class Consumer(multiprocessing.Process):
#     def __init__(self):
#         multiprocessing.Process.__init__(self)
#         self.stop_event = multiprocessing.Event()
        
#     def stop(self):
#         self.stop_event.set()
        
#     def run(self):
#         consumer = KafkaConsumer(bootstrap_servers='localhost:9092',
#                                  auto_offset_reset='earliest',
#                                  consumer_timeout_ms=1000)
#         consumer.subscribe(['my-topic'])

#         while not self.stop_event.is_set():
#             for message in consumer:
#                 print(message)
#                 if self.stop_event.is_set():
#                     break

#         consumer.close()

def main():

    # consumer = Consumer(
    #     os.environ['TOPICNAME'],
    #     os.environ['KAFKAPORT'],
    #     os.environ['CONSUMERGROUP']
    #     )

    # consumer = KafkaConsumer(
    #     'numtest',
    #     bootstrap_servers=['kafka:9092'],
    #     auto_offset_reset='earliest',
    #     enable_auto_commit=True,
    #     group_id='my-group',
    #     value_deserializer=lambda x: json.loads(x.decode('utf-8')))
    # consumer.subscribe('numtest')

    # for message in consumer:
    #     message = message.value
    #     print('{} added to topic'.format(message))

    #     while not self.stop_event.is_set():
    #         for message in consumer:
    #             print(message)
    #             if self.stop_event.is_set():
    #                 break

    #     consumer.close()

    # consumer.subscribe(['timeseries_2'])

    # print("waiting for messages")
    # counter = 0
    # for message in consumer:
    #     counter = counter + 1
    #     print("%s:%d:%d: key=%s value=%s" % (message.topic, message.partition, message.offset, message.key, message.value))
    #     print(counter)    

    # consumer.close()



    # To consume latest messages and auto-commit offsets
    consumer = KafkaConsumer('timeseries_2',
                            group_id='hi',
                            client_id='kuh',
                            bootstrap_servers=['kafka:9092'],
                            fetch_max_wait_ms=100,
                            request_timeout_ms=500,
                            metadata_max_age_ms = 100, 
                            )

    for message in consumer:
        # message value and key are raw bytes -- decode if necessary!
        # e.g., for unicode: `message.value.decode('utf-8')`
        print(message.value.decode('utf-8'))
        # consumer.close()                                    
    

    # '''
    # client = KafkaClient(hosts=os.environ['KAFKAPORT'])
    # topic = client.topics[os.environ['TOPICNAME']]
    # # zk = KazooClient()

    # balanced_consumer = topic.get_balanced_consumer(
    #     consumer_group="testing1",
    #     # managed = True,
    #     auto_commit_enable=True,
    #     auto_offset_reset='earliest',
    #     # zookeeper_connect='zookeeper:2181',
    #     zookeeper_hosts = 'zookeeper:2181',
    #     consumer_timeout_ms=1000,
    #     )    
    # for message in balanced_consumer:
    #     if message is not None:
    #         # print(message)
    #         print(message.offset, message.value)
    #     # if message.value == b'test message 1':
    #     #     break
    # '''

    return

if __name__ == "__main__":
    #Call main function
    main()
