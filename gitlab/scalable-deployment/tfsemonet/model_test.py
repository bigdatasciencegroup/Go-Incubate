import unittest 
import tensorflow as tf
import os

savedModelPath = os.path.join(os.path.dirname(__file__), 'cnn/1/')

def savedModel(path):
    # Recreate the exact same model
    model = tf.keras.experimental.load_from_saved_model(path)
    # model.summary()
    inputLayer = model.input
    outputLayer = model.output
    return inputLayer, outputLayer

class SimpleTest(unittest.TestCase): 
    def testModelInputOutput(self):		 
        path = savedModelPath
        inputLayer, outputLayer = savedModel(path)
        self.assertEqual(inputLayer.shape.as_list(), [None,48,48,1])
        self.assertEqual(inputLayer.dtype, tf.float32)
        self.assertEqual(outputLayer.shape.as_list(), [None,7])
        self.assertEqual(outputLayer.dtype, tf.float32) 

if __name__ == '__main__': 
    unittest.main()
