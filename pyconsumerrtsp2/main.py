from dotenv import load_dotenv
import os, sys, json
import logging
import cv2
from confluent_kafka import Consumer, KafkaError, OFFSET_END
import message
import dataprocessing.alg as alg

#Callback to provide handling of customized offsets on completion of a successful partition re-assignment 
def assignStrategy(consumer, partitions):
    for p in partitions:
        # OFFSET_BEGINNING = Set offset to start of partition 
        # OFFSET_END = Set offset to end of partition 
        # OFFSET_STORED = Set offset to last commit of partition. If no last commit, defaults to auto.offset.reset . 
        p.offset = OFFSET_END   
    consumer.assign(partitions)
    # print("Offset assignment:", partitions) 

def main():
    # Create consumer
    consumer = Consumer({
        'bootstrap.servers': os.getenv('KAFKAPORT'), #Address of Kafka
        'group.id': os.getenv('CONSUMERGROUP'),
        'client.id': os.getenv('CONSUMERID'),
        'enable.auto.commit': True,
        'auto.offset.reset': 'earliest', #Consume messages from earliest offset committed
        'enable.partition.eof':False, #Partition EOF informs that consumer has reached end of partition. To ignore/avoid EOF message, set `enable.partition.eof` = False.
        'api.version.request': True,
    })
    #Subscribe to the topic
    consumer.subscribe([os.getenv('TOPICNAME')],on_assign=assignStrategy)

    # Prepare openCV window
    print("OpenCV:",cv2.__version__)
    windowName = os.getenv('WINDOWNAME')
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
            img = message.handle(val)

            #Show image
            cv2.imshow(windowName, img)
            cv2.waitKey(1)

            #Process the message
            model.run(img)

            #Move the commit to the last message in the topic queue
            assignStrategy(consumer,consumer.assignment())

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
