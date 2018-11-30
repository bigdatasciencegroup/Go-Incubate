from dotenv import load_dotenv
import os, sys, json
import cv2
from confluent_kafka import Consumer, KafkaError
import message
import dataprocessing.alg as alg

def main():
    # Create consumer
    consumer = Consumer({
        'bootstrap.servers': os.getenv('KAFKAPORT'), #Address of Kafka
        'group.id': os.getenv('CONSUMERGROUP'),
        'client.id': os.getenv('CONSUMERID'),
        'enable.auto.commit': True,
        'auto.offset.reset': 'latest',  #Options: 'earliest','latest'
        'enable.partition.eof':False, #Partition EOF informs that consumer has reached end of partition. To ignore/avoid EOF message, set `enable.partition.eof` = False.
        'api.version.request': True,
    })
    #Subscribe to the topic
    consumer.subscribe([os.getenv('TOPICNAME')])

    # Prepare openCV window
    print("OpenCV:",cv2.__version__)
    windowName = "RTSPvideo1"
    cv2.namedWindow(windowName)
    cv2.resizeWindow(windowName, 240, 160)

    #Instantiate a signal processing model
    model = alg.Model()

    try:
        while True:
            #Consume message
            msg = consumer.poll(timeout=0.1)

            #Check whether message is None or if it contains error
            if msg is None:
                continue
            if msg.error():
                if msg.error().code() == KafkaError._PARTITION_EOF:
                    #Partition EOF indicates that consumer has reached end of partition. 
                    #To ignore/avoid EOF message, set "`enable.partition.eof` = False" in consumer config.
                    continue
                else:
                    print(msg.error())
                    break

            #Print message info        
            print('Received:{}, {}, {}, {}'.format(msg.timestamp(), msg.topic(), msg.partition(), msg.offset()))

            #Message handler
            val = message.decode(msg.value())
            img = message.handler(val)

            #Show image
            cv2.imshow(windowName, img)
            cv2.waitKey(1)

            #Process the message
            model.run(img)

    except KeyboardInterrupt:
        sys.stderr.write('Operation aborted by user\n')

    finally:
        #Close down consumer to commit final offsets
        consumer.close()
        sys.stderr.write("Consumer closed: {}\n".format(os.getenv('CONSUMERID')))

    return

if __name__ == "__main__":
    # Load environment variable
    load_dotenv(override=True)
    # Call main function
    main()
