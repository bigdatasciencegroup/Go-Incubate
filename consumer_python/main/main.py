import os, sys
from kafkapc_python import Consumer

def main():

    consumer = Consumer(
        os.environ['TOPICNAME'],
        os.environ['KAFKAPORT'],
        os.environ['CONSUMERGROUP']
        )

    try:
        for message in consumer:
            # print("-------",message)
            val = message.value
            print(val)
    except KeyboardInterrupt:
        sys.exit()


    return

if __name__ == "__main__":
    #Call main function
    main()
