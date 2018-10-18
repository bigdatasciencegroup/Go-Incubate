import os
from kafkapc import Consumer

def main():

    consumer = Consumer(
        os.environ['TOPICNAME'],
        os.environ['KAFKAPORT'],
        os.environ['CONSUMERGROUP']
        )

    for message in consumer:
        val = message.value
        print(val)

    return

if __name__ == "__main__":
    #Call main function
    main()
