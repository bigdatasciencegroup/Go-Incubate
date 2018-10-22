import os
from kafkapc_python import Producer

def on_send_success(record_metadata):
    print("on success ---------------------------")
    print(record_metadata.topic)
    print(record_metadata.partition)
    print(record_metadata.offset)
    print("---------------------------")

def main():

    producer = Producer(
        os.environ['KAFKAPORT']
        )

    for e in range(10):
        data = {'number': e}
        producer.send(os.environ['TOPICNAME'], data).add_callback(on_send_success)
        # future = producer.send(os.environ['TOPICNAME'], b'tester string')
        print("manual data print ---> ",data)

    # block until all async messages are sent
    producer.flush()

    return

if __name__ == "__main__":
    #Call main function
    main()
