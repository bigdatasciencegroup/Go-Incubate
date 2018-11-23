from dotenv import load_dotenv
import kafkapc_python as pc
import os
import cv2
import json
import base64
import numpy as np

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
    cv2.resizeWindow("RTSPvideo", 160*10, 240*10)

    # Start consuming video
    for message in consumer:
        val = message.value
        rows = val['rows']
        cols = val['cols']
        timestamp = message.timestamp

        base64string = val['pix'] #pix is base-64 encoded string
        byteArray = base64.b64decode(base64string)
        npArray = np.frombuffer(byteArray, np.uint8)

        imgR = npArray[0::4].reshape((rows, cols))
        imgG = npArray[1::4].reshape((rows, cols))
        imgB = npArray[2::4].reshape((rows, cols))
        img = np.stack((imgR, imgG, imgB))
        img = np.moveaxis(img, 0, -1)
        
        # img = cv2.imdecode(npArray, cv2.IMREAD_COLOR)  # cv2.IMREAD_COLOR in OpenCV 3.1
        # img = cv2.imread(npArray, 1)
        # print(npArray)

        print(timestamp)         
        cv2.imshow('RTSPvideo', img)
        cv2.waitKey(1)
        
        # consumer.close()                                    
    return

if __name__ == "__main__":
    # Load environment variable
    load_dotenv(override=True)
    # Call main function
    main()
