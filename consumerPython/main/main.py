# import json
# from kafka import KafkaConsumer
# from kafka.structs import OffsetAndMetadata, TopicPartition
from pykafka import KafkaClient

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

    client = KafkaClient(hosts="kafka:9092")
    topic = client.topics['my_test']
    balanced_consumer = topic.get_balanced_consumer(
        consumer_group=('testgroup').encode(),
        # managed = True,
        auto_commit_enable=True,
        auto_offset_reset='earliest',
        zookeeper_connect='zookeeper:2181')
        # consumer_timeout_ms=1000)
    # balanced_consumer = balanced_consumer(
    #     auto_commit_enable=True, 
    #     auto_commit_interval_ms=6, 
    #     queued_max_messages=1, 
    #     fetch_min_bytes=1, 
    #     fetch_wait_max_ms=100, 
    #     zookeeper_hosts='zookeeper:2181', 
    #     auto_start=True, 
    #     reset_offset_on_start=False,
    #     reset_offset_on_fetch=True)    
    for message in balanced_consumer:
        if message is not None:
            print(message.offset, message.value)
        if message.value == b'test message 1':
            break


    return

if __name__ == "__main__":
    #Call main function
    main()
