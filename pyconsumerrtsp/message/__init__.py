import base64
import numpy as np
         
def handler(val):
    # Conversion: base-64 string --> array of bytes --> array of integers
    base64string = val['pix'] #pix is base-64 encoded string
    byteArray = base64.b64decode(base64string)
    npArray = np.frombuffer(byteArray, np.uint8)

    # Reshape array into image matrix 
    rows = val['rows']
    cols = val['cols']
    imgR = npArray[0::4].reshape((rows, cols))
    imgG = npArray[1::4].reshape((rows, cols))
    imgB = npArray[2::4].reshape((rows, cols))
    img = np.stack((imgR, imgG, imgB))
    img = np.moveaxis(img, 0, -1)

    return img


    