---
title: "Deep Learning for Time Series and NLP"
excerpt: "Project"
header:
  teaser: assets/images/time_series_nlp_01.jpg
---

{% include feature_row %}

<!-- {% include toc %} -->

## Project

### Problem
In this project we develop a deep learning algorithm, based on a combination of convolutional neural network(CNN) and long-short term memory(LSTM) network, which is applicable for time series data classification and for natural language processing (NLP) applications. The algorithm development will be demonstrated using a 12-way text classification problem. A limited sample of training sentences, corresponding labels, and test sentences are given in the repository.

```
sentence: letwmktwviuhskenamuhtwlpuhraletwlrvienskuhqvmvamuhxepmuhlekrpmamuhtwamuluhezpmlexeuhtwulultwvienqjuhtwmvypkrxzuhsktwmkpmiwuhskenamuhtwlpuhenuhxepmuhtwmkpmiwuhtwamuluhxepmuhsaendfuhtwamulbhbhsaendfuhqvijsaenvileenmcuhqvtwiwleenameebhbhsaendfuhtwvipmuhtwvipmlruhsaiwtvenmvleenmkvimvuhqvenamuhvienezuhenuhxepmuhskiwlepmdfuhtwamuluhqgqvtwskkrulmvuleniwuhvitwiwiwenxeuhvimvuhletwulvimvdfuhsaiwulqvpmezuhqvmvuhulmvuhvitwamdf
label: 10
```
Sentences in the file have been encoded in a deterministic one-to-one mapping, to prove that the algorithm developed is data agnostic such that it can be applied to any other NLP or time series data classification problem.

### Repository
The [repository](https://github.com/Adaickalavan/Deep-Learning-for-Time-Series-and-NLP) contains the following project files:
```
Deep-Learning-for-Time-Series-and-NLP   # main folder
├── challenge.py                        # code in Python script 
├── challenge.ipynb                     # code in iPython notebook
├── xtrain.txt                          # limited sample training/validation set
├── ytrain.txt                          # limited sample labels for training/validation set
├── xtest.txt                           # limited sample test set
└── cnn_lstm-180-0.87.hdf5              # sample saved tensorflow model 
```

### Solution 
A complete step-by-step approach to solving the problem is described next. The following are included:
* Reading in of files
* Data wrangling
* Partitioning train/validate/test sets
* Building tensorflow machine learning model
* Checking validation set accuracy
* Saving best tensorflow model
* Writing test set predictions to file

We will be using the following tools:
* Python
* Keras
* Jupyter Notebook

Import the required libraries
```python
import pdb
import numpy as np
import csv
import tensorflow as tf
from sklearn.model_selection import train_test_split
from sklearn.utils.class_weight import compute_class_weight

from keras.models import load_model, Model
from keras.layers import Dense, Activation, Dropout, Input, LSTM, Lambda, Bidirectional, Conv1D, MaxPooling1D
import keras.callbacks
```
Read in the files
```python
def readIn(filepath):
    '''
    Arguments:
    filepath -- path of file to be read

    Returns:    
    x -- list of strings read from the input file    
    '''
    # Read in the files
    with open(filepath, 'r') as x:
        x = x.read()
        x = x.splitlines()  # Split files into sentences

    return x
```
```python
x_traindev = readIn('./xtrain.txt')
y_traindev = readIn('./ytrain.txt')
```
Throughout this python example, we shall use this notation:
* m is number of examples
* n_v is vocabulary size (i.e., vocab_size)
* n_c is number of classes (i.e., number of novels to predict)
* n_a is dimensionality of the hidden state in LSTM cell
* Tx is length of each string. All strings will be appended or truncated to the same length  

Set the desired string length and number of classes
```python
# Compute the maximum string length
Txmax = len(max(x_traindev, key=len))
print(Txmax)

# We set desired string length to 448 which is 5 char shorter than Txmax. 
# This simplifies the dimensions during convolution later, without significant performance loss
Tx = 448
# Set the number of classes
n_c = 12
```
Each sentence in the data set is a sequence of characters. Hence, we will build a character level sentence classification model.

First, build a character dictionary. Then, convert each training sample (i.e., sentence) from character sequence to integer sequence using the dictionary.
```python
def char2ix(x):
    '''
    Arguments:
    x -- list of strings
    
    Returns:
    char_to_ix -- dictionary mapping character to integer
    n_v -- vocabulary size
    '''
    # Create a python dictionary to map each character to an index 0-26.
    string = ''.join(x)
    chars = list(string)  # Create a list of unique characters (i.e., a to z and \n newline)
    chars_set = set(chars)
    n_v = len(chars_set)  # Vocabulary size
    char_to_ix = {ch: ii for ii, ch in enumerate(sorted(chars_set))}
    print(char_to_ix)

    return char_to_ix, n_v

def listOfStr_to_array(x, Tx, func):
    '''
    Arguments:
    x -- list of strings.  
    Tx -- maximum length of string
    func -- function to map characters to integers
    
    Returns:
    x_out -- numpy array. Shape = (m, Tx).
    '''

    x_out = np.ones((len(x), Tx), dtype=np.int64)*-1  # Create empty array to store one-hot encoding result
    for ii, string in enumerate(x):  # Iterate over each string in the file
        x_array = np.array(list(map(func, string)))  # Map characters to integers
        if len(x_array) <= Tx:
            x_out[ii, : len(x_array)] = x_array
        else:
            x_out[ii, :] = x_array[: Tx]

    return x_out
```
```python
char_to_ix, n_v = char2ix(x_traindev)
x_traindev = listOfStr_to_array(x_traindev, Tx, lambda x: char_to_ix[x])

# Check type and shape x_traindev
print('x_traindev is of type {} and of shape {}'.format(type(x_traindev), x_traindev.shape))
```
Cast the y labels into an integer sequence
```python
y_traindev = np.array(list(map(int, y_traindev)))
```
Shuffle stratify-split the data into train and dev sets
```python
x_train, x_dev, y_train, y_dev = train_test_split(x_traindev, 
                                                  y_traindev,
                                                  test_size=0.20, 
                                                  random_state=1,
                                                  shuffle=True, 
                                                  stratify=y_traindev)
# Check type and shape x_train, x_dev
print('x_train is of type {} and of shape {}'.format(type(x_train), x_train.shape))
print('x_dev is of type {} and of shape {}'.format(type(x_dev), x_dev.shape))
```
Since the number of classes are imbalanced, we compute and apply class weights to the cost function
```python
#Compute class weights using y_train     
unique, counts = np.unique(y_train, return_counts=True)
print('Class distribution:')
print(dict(zip(unique, counts)))
class_weight_vec = compute_class_weight('balanced',np.array(range(12)),y_train)
class_weight_dict = {ii: w for ii, w in enumerate(class_weight_vec)}
```
Convert training and dev labels to one hot matrices
```python
# Convert a list of characters into one-hot encoding numpy array
def convert_to_one_hot(x, n):
    '''
    Arguments:
    x -- array of integers. 
    n -- number of unique values. scalar value.
   
    Returns:
    y -- one hot encoded matrix of integers. shape = (len(x), n).
    '''

    y = np.eye(n)[x.reshape(-1)]
    return y
```
```python
y_train_onehot = convert_to_one_hot(y_train, n_c)
y_dev_onehot = convert_to_one_hot(y_dev, n_c)
print('y_train_onehot is of type {} and of shape {}'.format(type(y_train_onehot), y_train_onehot.shape))
print('y_dev_onehot is of type {} and of shape {}'.format(type(y_dev_onehot), y_dev_onehot.shape))
```
Build the machine learning model as follows: 

CNN (4 layer) --> LSTM (2 layer) --> Dense (1 layer) --> Softmax --> Cross entropy loss

1. CNN is suitable for character level sequence (or time series) classification. It helps to extract relevant patterns from the sequences along the feature and time dimensions.
2. LSTM helps to recognise sequential information
3. Fully connected Dense layer with Softmax reduces the output to 12 desired class probability distribution
4. Cross entropy loss is used as the cost for this multiclass classification problem.

```python
def one_hot(x, numOfVocab):
    '''
    Arguments:
    x -- input array of shape (m,Tx) where m = len(x)
    numOfVocab -- vocabulary size
    
    Returns:
    return -- output array of shape (m, Tx, numOfVocab)
    '''
    import tensorflow as tf
    return tf.to_float(tf.one_hot(x, numOfVocab, on_value=1, off_value=0, axis=-1))

def one_hot_outshape(x):
    '''
    Arguments:
    x -- input shape
    
    Returns:
    return -- one hot encoded shape of x, i.e., (m, Tx, numOfVocab)     
    '''
    return x[0], x[1], 26
```
```python
print('Building the model')
kernel_length = [ 5,   3,   3,   3]
filter_num    = [64,  64, 128, 128]
pool_length   = [ 2,   2,   2,   2]
x_input = Input(shape=(Tx,), dtype='int64')
# The one-hot encoded vectors for each input string is built on the fly using the Lambda layer
# Desired shape of input data X = (m, Tx, n_v)
x_one_hot = Lambda(one_hot, output_shape=one_hot_outshape, arguments={'numOfVocab': 26})(x_input)
x_conv = x_one_hot
for ii in range(len(filter_num)):
    x_conv = Conv1D(filters=filter_num[ii],
                    kernel_size=kernel_length[ii],
                    padding='same')(x_conv)
    x_conv = Activation('relu')(x_conv)
    x_conv = MaxPooling1D(pool_size=pool_length[ii])(x_conv)

x_lstm1 = Bidirectional(LSTM(128, return_sequences=True))(x_conv)
x_lstm1_dropout = Dropout(0.3)(x_lstm1)
x_lstm2 = LSTM(128, return_sequences=False)(x_lstm1_dropout)
x_lstm2_dropout = Dropout(0.3)(x_lstm2)
x_out = Dense(n_c, activation='softmax')(x_lstm2_dropout)
model = Model(inputs=x_input, outputs=x_out, name='textClassifier')
# Check model summary
model.summary()
```
```
Building the model
_________________________________________________________________
Layer (type)                 Output Shape              Param #   
=================================================================
input_1 (InputLayer)         (None, 448)               0         
_________________________________________________________________
lambda_1 (Lambda)            (None, 448, 26)           0         
_________________________________________________________________
conv1d_1 (Conv1D)            (None, 448, 64)           8384      
_________________________________________________________________
activation_1 (Activation)    (None, 448, 64)           0         
_________________________________________________________________
max_pooling1d_1 (MaxPooling1 (None, 224, 64)           0         
_________________________________________________________________
conv1d_2 (Conv1D)            (None, 224, 64)           12352     
_________________________________________________________________
activation_2 (Activation)    (None, 224, 64)           0         
_________________________________________________________________
max_pooling1d_2 (MaxPooling1 (None, 112, 64)           0         
_________________________________________________________________
conv1d_3 (Conv1D)            (None, 112, 128)          24704     
_________________________________________________________________
activation_3 (Activation)    (None, 112, 128)          0         
_________________________________________________________________
max_pooling1d_3 (MaxPooling1 (None, 56, 128)           0         
_________________________________________________________________
conv1d_4 (Conv1D)            (None, 56, 128)           49280     
_________________________________________________________________
activation_4 (Activation)    (None, 56, 128)           0         
_________________________________________________________________
max_pooling1d_4 (MaxPooling1 (None, 28, 128)           0         
_________________________________________________________________
bidirectional_1 (Bidirection (None, 28, 256)           263168    
_________________________________________________________________
dropout_1 (Dropout)          (None, 28, 256)           0         
_________________________________________________________________
lstm_2 (LSTM)                (None, 128)               197120    
_________________________________________________________________
dropout_2 (Dropout)          (None, 128)               0         
_________________________________________________________________
dense_1 (Dense)              (None, 12)                1548      
=================================================================
Total params: 556,556
Trainable params: 556,556
Non-trainable params: 0
_________________________________________________________________
```
Define an optimizer and compile the model
```python
model.compile(loss='categorical_crossentropy', optimizer='adam', metrics=['accuracy'])
```
Save the model at regular epoch intervals
```python
checkpoint = keras.callbacks.ModelCheckpoint('./cnn_lstm'+'-{epoch:02d}-{val_acc:.2f}.hdf5',
                                             monitor='val_loss', 
                                             verbose=1, 
                                             save_best_only=False, 
                                             save_weights_only=False, 
                                             mode='auto', 
                                             period=10)
```
Fit the model
```python
print('Fitting the model now')
model.fit(x_train, y_train_onehot,
          validation_data=(x_dev, y_dev_onehot),
          batch_size=256, 
          epochs=130,
          shuffle=True,
          callbacks=[checkpoint],
          class_weight=class_weight_dict)

# Save model to continue later
model.save('./cnn_lstm.hdf5')
```
We have already trained the above model over 180 epochs using our full training set, which yielded a training set accuracy of ~99% and validation set accuracy of ~87%.

Hence, we now load the trained model containing the trained parameters.
```python
#Load model to continue testing
#Our saved file is named "cnn_lstm-180-0.87.hdf5"
model = load_model('./cnn_lstm-180-0.87.hdf5')
```
First, check accuracy on training and validation set.
```python
_, acc_train = model.evaluate(x_train, y_train_onehot)
_, acc_dev = model.evaluate(x_dev, y_dev_onehot)
print('Train set accuracy: {}'.format(acc_train))
print('Validation set accuracy: {}'.format(acc_dev))
```
Read in the test set
```python
#Read in the test set
x_test = readIn('./xtest.txt')
x_test = listOfStr_to_array(x_test, Tx, lambda x: char_to_ix[x])

# Check type and shape x_test
print('x_test is of type {} and of shape {}'.format(type(x_test), x_test.shape))
```
Make prediction on test set and write the predcited class labels of test sample into file
```python
def write2file(filename,resultList):
    with open(filename,'w',newline='') as outputFile:
        writer = csv.writer(outputFile)
        for result in resultList:
            writer.writerow([result]) 
```
```python
#Make predictions
pred = model.predict(x_test)

#Write predictions to file
pred_labels = np.argmax(pred,axis=1)
write2file('ytest.txt',pred_labels)
```

The end
