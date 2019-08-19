import unittest 
import tensorflow as tf
import os
import json
import pprint

savedModelPath = os.path.join(os.path.dirname(__file__), 'cnn/1/')
print(savedModelPath)


def savedmodel(path):
    # Recreate the exact same model
    model = tf.keras.experimental.load_from_saved_model(path)
    model.summary()
    # json_string = model.to_json()
    # pprint.pprint(json.loads(json_string))
    # json_string
    print("---------->>>>>>>",model.input)
    print("---------->>>>>>>",model.output.shape)
    k = model.output.shape
    print("----++++++___+++___",type(k))
    # print(p)
    # with tf.Session(graph=tf.Graph()) as sess:
    #     tf.saved_model.loader.load(sess, ["serve"], export_path)
    # graph = tf.get_default_graph()
    # print(graph.get_operations())
    # sess.run('myOutput:0', feed_dict={'myInput:0': })

class SimpleTest(unittest.TestCase): 

	# Returns True or False. 
	def test(self):		 
        path = savedModelPath
		self.assertEqual(savedmodel(path),float) 

    

if __name__ == '__main__': 
    print("Testing")
    unittest.main()