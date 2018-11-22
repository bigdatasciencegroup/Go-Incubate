from dotenv import load_dotenv
import kafkapc_python as pc
import os
import cv2
import json
import base64
import numpy as np



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

        base64string = val['pix'] #pix is base-64 encoded string
        byteArray = base64.b64decode(base64string)
        npArray = np.frombuffer(byteArray, np.uint8)
        img = cv2.imdecode(npArray, cv2.IMREAD_COLOR) # cv2.IMREAD_COLOR in OpenCV 3.1
        print(npArray)


        
        # cv2.imshow('frame', val['pix'])
        # cv2.waitKey(1)
        
        # consumer.close()                                    

# Pix:      []uint8{9, 10, 200, 64},
# "pix":"CQrIQA=="
# "pix":CQrIQA=="
# []byte encodes as a base64-encoded string

    return

if __name__ == "__main__":
    # Load environment variable
    load_dotenv(override=True)
    # Call main function
    main()
