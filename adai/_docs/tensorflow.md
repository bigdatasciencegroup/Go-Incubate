---
title: "TensorFlow"
permalink: /docs/tensorflow/
---

Every machine learning project must produce a saved model as its final deliverable.

## SavedModel
+ Contains complete architecture and weights. 
+ Independent of the source code that created the model.
+ Suitable for deployment via TensorFlow Serving, TensorFlow Lite, in many programming languages (e.g., C++, Java, Go)
+ A saved model produces a folder consisting of the following files:
    ```text
    tfserving                                           # TensorFlow Serving 
    ├── cnn                                             # model name 
    |    ├── 1                                          # model version
    |    |   ├── assets/
    |    |   ├── variables/
    |    |   |   ├── variables.data-*****-of-*****
    |    |   |   └── variables.index
    |    |   └── saved_model.pb
    |    └── 2                                          # model version
    |        ├── assets/
    |        ├── variables/
    |        |   ├── variables.data-*****-of-*****
    |        |   └── variables.index
    |        └── saved_model.pb 
    └── mtcnn                                           # model name
        └── 1                                           # model version
            ├── assets/
            ├── variables/
            |   ├── variables.data-*****-of-*****
            |   └── variables.index
            └── saved_model.pb
    ```
    + The .pb is the MetaGraphDef which holds the graph structure. MetaGraph is a dataflow graph, plus its associated variables, assets, and signatures. A MetaGraphDef is the protocol buffer representation of a MetaGraph.
    + The variables folder holds your learned weights. 
    + The assets folder allow you to add external files that may be needed.
+ Refer to the [official tutorial](https://www.tensorflow.org/tfx/tutorials/serving/rest_simple#save_your_model) for saving and loading a TensorFlow model. 
+ A simplified example of writing TensorFlow machine leanring model and saving it into SavedModel in Python is given below.
<br/><img src="/assets/images/wiki/wiki_tensorflow_01.jpg" width="100%" height="200%"/>
+ An overview of the TensorFlow SavedModel building and deploying process is illustrated here.
<br/><img src="/assets/images/wiki/wiki_tensorflow_02.jpg" width="100%" height="200%"/>