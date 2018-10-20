import os
from kafkapc_python import Producer

def main():

    producer = Producer(
        os.environ['KAFKAPORT']
        )
    for e in range(1000):
        data = {'number': e}
        producer.send(os.environ['TOPICNAME'], value=data)
        print(data)

    return

if __name__ == "__main__":
    #Call main function
    main()
