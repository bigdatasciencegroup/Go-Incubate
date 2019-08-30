---
title: "How objects are called in Keras"
tags: 
  - Keras 
categories:
  - tensorflow   
last_modified_at: 2018-03-07
---

This post elucidates how layers (for instance ```class MyLayer(Layer)```) are called in Keras when we perform operations such as ``x_output = MyLayer( .... , ...)(x_input)``. In this example, the ``MyLayer`` class instance is called directly with input ``x_input``. This action calls the ``def __call__()`` function, found in the parent of ``MyLayer`` which is ``class Layer()``. The ``def __call__()`` function in the parent class, in-turn calls the ``def call()`` in ``MyLayer`` class which defines the logic of ``MyLayer`` layer.

I have given a simplified example below describing this operation flow.

```python
class Layer():
    def __init__(self,x):
        print('Inside Layer -> __init__ :',x)

    def __call__(self, y):
        print('Inside Layer -> __call__ :',y)
        self.call(y)

    def call(self,z):
        print('Inside Layer -> call :',z)

class MyLayer(Layer):
    def __init__(self,k):
        print('Inside MyLayer -> __init__ :',k)
        super(MyLayer, self).__init__(k)

    def call(self, m):
        print('Inside MyLayer -> call :',m)    

if __name__ == '__main__':
    obj = MyLayer(3)
    obj(4)
```

Upon running the above code, the output is as follows:

```python
Inside MyLayer -> __init__ : 3
Inside Layer -> __init__ : 3
Inside Layer -> __call__ : 4
Inside MyLayer -> call : 4
``` 

<!--
This post elucidates how layers (for instance [```class LSTM(RNN)```](https://github.com/keras-team/keras/blob/1c9a49781da2101507db23e2014e4e5d16bd2e52/keras/layers/recurrent.py#L1996)) are called in Keras when we perform operations such as ``x_output = LSTM( .... , ...)(x_input)``. In this example, the ``LSTM`` class instance is called directly with input ``x_input``. This action calls the ``def __call__()`` function, which is found in the parent class of ``LSTM`` which is the [``class RNN ()``](https://github.com/keras-team/keras/blob/1c9a49781da2101507db23e2014e4e5d16bd2e52/keras/layers/recurrent.py#L212). The ``def __call__()`` function in the parent class, in turn calls the [``def call()``](https://github.com/keras-team/keras/blob/1c9a49781da2101507db23e2014e4e5d16bd2e52/keras/layers/recurrent.py#L2145) in ``LSTM`` which defines the logic of the LSTM layer.

I have given a simplified example below describing this operation flow. Here, ``class MyLayer()`` corresponds to a layer/object-instance in Keras and ``class Layer()`` corresponds to the parent class. 
-->
