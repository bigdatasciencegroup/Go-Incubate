import os, sys
from kafkapc_python import Consumer

def main():

    consumer = Consumer(
        os.environ['TOPICNAME'],
        os.environ['KAFKAPORT'],
        )

    print("waiting for messages")
    print("waiting for messages")
    print("waiting for messages")
    print("waiting for messages")
    for message in consumer:
        print(message.value)
        print("%s:%d:%d: key=%s value=%s" % (
            message.topic,
            message.partition,
            message.offset,
            message.key,
            message.value))

    print("ENDNEDNEND")
    return

if __name__ == "__main__":
    #Call main function
    main()
