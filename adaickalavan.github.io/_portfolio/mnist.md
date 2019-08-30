---
title: "MNIST"
excerpt: "Neural network and convolutional neural networks for MNIST"
header:
  teaser: /assets/images/MNIST_1.jpg
feature_row:
  - image_path: /assets/images/MNIST_1.jpg
  - image_path: /assets/images/MNIST_2.jpg
---

{% include feature_row %}

{% include toc %}

## Summary
This page contains different Python implementations for recognising handwritten digits from the famous MNIST database.

## Repository
The [repository](https://github.com/Adaickalavan/MNIST) consists of the following: 
* Multilayer Perceptron using NumPy - Python codes
* Multilayer Perceptron using Keras and Theano - Python codes
* Convolutional Neural Network using Theano - Python codes
* Multiclass linear regression using TensorFlow - Python codes

## Info

### MNIST MLP Numpy
Implemented a 2-layer feedforward neural network (30 hidden nodes with sigmoid activation, 10 output nodes with multiclass sigmoid activation, cross entropy cost function) in Python using NumPy for handwritten digit recognition from MNIST database. Test set accuracy is >95%. 

### MNIST MLP Keras
Implemented a 3-layer feedforward neural network (50 nodes in each hidden layer with tanh activation, 10
output nodes with softmax activation, cross entropy cost function) in Python using Theano & Keras for
handwritten digit recognition from MNIST database. Test set accuracy is >94%. 

### MNIST CNN Theano
Implemented a convolutional neural network (2 convolution + max pooling layers consisting of 20 and 50
filters with 2-by-2 subsampling, fully-connected hidden layer using 500 nodes with tanh activation, 10 output nodes with softmax activation, and negative log-likelihood cost function) in Python using Theano for handwritten digit recognition from MNIST database. Test set accuracy is >98%. 

### MNIST Multiclass Linear Regression TensorFlow
Implemented a single hidden layer feedforward neural network (784x10 weight matrix, output node with softmax, cross entropy cost function, and backpropagation with stochatic gradient descent) in Python using TensorFlow for handwritten digit recognition from MNIST database. Test set accuracy is >91%. 
