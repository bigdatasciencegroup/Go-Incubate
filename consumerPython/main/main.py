import json

from kafka import KafkaConsumer
from kafka.structs import OffsetAndMetadata, TopicPartition

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

    consumer = KafkaConsumer(
        'numtest',
        bootstrap_servers=['kafka:9092'],
        auto_offset_reset='earliest',
        enable_auto_commit=True,
        group_id='my-group',
        value_deserializer=lambda x: json.loads(x.decode('utf-8')))

    for message in consumer:
        message = message.value
        print('{} added to topic'.format(message))

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

    consumer.close()

    return

if __name__ == "__main__":
    #Call main function
    main()
