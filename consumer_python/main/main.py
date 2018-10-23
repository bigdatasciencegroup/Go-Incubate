import os, sys
# from kafkapc_python import Consumer
from kafka import KafkaConsumer
import json

def main():

    consumer = KafkaConsumer(
        os.environ['TOPICNAME'],
        bootstrap_servers=os.environ['KAFKAPORT'],
        auto_offset_reset='earliest',
        enable_auto_commit=True,
        group_id=os.environ['CONSUMERGROUP'],
        value_deserializer=lambda x: json.loads(x.decode('utf-8'))
        )

    # consumer = Consumer(
    #     os.environ['TOPICNAME'],
    #     os.environ['KAFKAPORT'],
    #     os.environ['CONSUMERGROUP']
    #     )

    print("waiting for messages")
    counter = 0
    for message in consumer:
        counter = counter + 1
        print("%s:%d:%d: key=%s value=%s" % (
            message.topic,
            message.partition,
            message.offset,
            message.key,
            message.value))
        print(counter)    
        # if message.value == "Hihi 999":
        #     break

    print("ENDNEDNEND")
    print(counter)
    return

if __name__ == "__main__":
    #Call main function
    main()
