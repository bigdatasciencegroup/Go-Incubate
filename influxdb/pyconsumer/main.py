from dotenv import load_dotenv
import os, sys, json
import logging
import cv2
from influxdb import InfluxDBClient
import message
import dataprocessing.alg as alg
import time
import base64
import numpy as np
import json
import datetime

def decode(val):
    return json.loads(val.decode('utf-8'))

def main():

    # Prepare openCV window
    print("OpenCV:",cv2.__version__)
    windowName = os.getenv('WINDOWNAME')
    cv2.namedWindow(windowName)
    cv2.resizeWindow(windowName, 240, 160)

    #Instantiate a signal processing model
    model = alg.Model()

    client = InfluxDBClient(host=os.getenv('INFLUXDBHOSTNAME'),
                            port=8086,
                            username='admin',
                            password='admin',
                            database='mydb'
                            )
    # Setup query
    query = 'SELECT last("frame") FROM "VideoFrames"'

    try:
        while True:
            #Consume message
            result = client.query(query,database='mydb')
            msgStr = list(result.get_points(measurement='VideoFrames'))[0]['last']
            # nparray = np.array([int(msg[ii:ii+2],16) for ii in range(0,len(msg),2)])

            #Print message info        
            print('Received : {}'.format(datetime.datetime.now()))

            #Message handler
            msg = message.decode(msgStr)
            img = message.handle(msg)

            #Show image
            cv2.imshow(windowName, img)
            cv2.waitKey(1)

            #Process the message
            model.run(img)
            # time.sleep(10)

    except KeyboardInterrupt:
        sys.stderr.write('Operation aborted by user\n')

    finally:
        #Close down consumer
        client.close()

    return

if __name__ == "__main__":
    # Load environment variable
    load_dotenv(override=True)
    # Call main function
    main()



