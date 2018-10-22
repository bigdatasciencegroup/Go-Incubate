import os, sys
from kafkapc_python import Consumer

def main():

    consumer = Consumer(
        os.environ['TOPICNAME'],
        os.environ['KAFKAPORT'],
        os.environ['CONSUMERGROUP']
        )

    try:
        print("waiting for messages")
        for message in consumer:
            print("waiting for messages")
            # print("-------",message)
            val = message.value
            print(val)
    except KeyboardInterrupt:
        sys.exit()


    return

if __name__ == "__main__":
    #Call main function
    main()
