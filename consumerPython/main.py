from dotenv import load_dotenv
import kafkapc_python as pc
import os
import cv2
import json
import numpy as np

# class Consumer(multiprocessing.Process):
#     def __init__(self):
#         multiprocessing.Process.__init__(self)
#         self.stop_event = multiprocessing.Event()
        
#     def stop(self):
#         self.stop_event.set()
        
#     def run(self):
#         consumer = KafkaConsumer(bootstrap_servers='localhost:9092',
#                                  auto_offset_reset='earliest',
#                                  consumer_timeout_ms=1000)
#         consumer.subscribe(['my-topic'])

#         while not self.stop_event.is_set():
#             for message in consumer:
#                 print(message)
#                 if self.stop_event.is_set():
#                     break

#         consumer.close()

def main():
    consumer = pc.Consumer(
        os.getenv('TOPICNAME'),
        os.getenv('KAFKAPORT'),
        os.getenv('CONSUMERGROUP')
        )

    # print(cv2.__version__)
    # cv2.namedWindow("RTSP video")

    for message in consumer:
        val = message.value
    #     print('{} added to topic'.format(val))
        pix = val['pix']
        print(pix)
        print(len(pix))
        # pixArray = np.array(pix)

        # print("Pix Array =", pixArray)
        
        # cv2.imshow('frame', val['pix'])
        # cv2.waitKey(1)
        
        # consumer.close()                                    

    return

if __name__ == "__main__":
    # Load environment variable
    load_dotenv(override=True)
    # Call main function
    main()
