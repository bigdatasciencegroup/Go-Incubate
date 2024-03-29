import base64
import numpy as np
import json

def decode(val):
    return json.loads(val)

def handle(val):
    # Conversion: base-64 string --> array of bytes --> array of integers
    base64string = val['pix'] #pix is base-64 encoded string
    byteArray = base64.b64decode(base64string) #byteArray is an array of bytes
    npArray = np.frombuffer(byteArray, np.uint8) #npArray is an array of integers

    # Reshape array into an RGB image matrix of shape (channels, rows, cols)
    rows = val['rows']
    cols = val['cols']
    imgR = npArray[0::4].reshape((rows, cols))
    imgG = npArray[1::4].reshape((rows, cols))
    imgB = npArray[2::4].reshape((rows, cols))
    img = np.stack((imgR, imgG, imgB))
    img = np.moveaxis(img, 0, -1)

    # img = cv2.imdecode(npArray, cv2.IMREAD_COLOR)  # cv2.IMREAD_COLOR in OpenCV 3.1
    # img = cv2.imread(npArray, 1)

    return img


    