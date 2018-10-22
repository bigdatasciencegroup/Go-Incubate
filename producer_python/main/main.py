import os
from kafkapc_python import Producer
import json

def main():

    producer = Producer(
        os.environ['KAFKAPORT']
        )

    for e in range(10):
        data = {"number": e}    
        producer.send(os.environ['TOPICNAME'], value=data)
        print("Data sent ---> ",data)

    # block until all async messages are sent
    producer.flush()

    return

if __name__ == "__main__":
    #Call main function
    main()
