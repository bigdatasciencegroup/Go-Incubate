from dotenv import load_dotenv
import kafkapc_python as pc
import os
import cv2
import json
import base64
import numpy as np
import message

def main():
    # Create consumer 
    consumer = pc.Consumer(
        os.getenv('TOPICNAME'),
        os.getenv('KAFKAPORT'),
        os.getenv('CONSUMERGROUP')
        )

    # Prepare openCV window
    print(cv2.__version__)
    cv2.namedWindow("RTSPvideo")
    cv2.resizeWindow("RTSPvideo", 240, 160)

    # Start consuming video
    for m in consumer:
        val = m.value
        timestamp = m.timestamp
        print(type(timestamp))

        #Message handler
        img = message.handler(val)
        
        # img = cv2.imdecode(npArray, cv2.IMREAD_COLOR)  # cv2.IMREAD_COLOR in OpenCV 3.1
        # img = cv2.imread(npArray, 1)

        print("---->>>>",timestamp)         
        cv2.imshow('RTSPvideo', img)
        cv2.waitKey(1)
        
    consumer.close()                                    
    
    return

if __name__ == "__main__":
    # Load environment variable
    load_dotenv(override=True)
    # Call main function
    main()
