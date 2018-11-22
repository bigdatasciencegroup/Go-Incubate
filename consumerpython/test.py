import base64
import io
import cv2 as cv21
from imageio import imread
import matplotlib.pyplot as plt
import numpy as np

base64_string = "CQrIQA=="
imgdata = base64.b64decode(base64_string)
print(type(imgdata))
nparr = np.fromstring(imgdata, np.uint8)
img_np = cv21.imdecode(nparr, cv21.IMREAD_COLOR) # cv2.IMREAD_COLOR in OpenCV 3.1

print(type(img_np))
print(img_np)

