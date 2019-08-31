---
title: "Python"
permalink: /docs/python/
---

## Code style
+ Follow the [PEP 8 Python style guide](https://www.python.org/dev/peps/pep-0008/), except TensorFlow uses 2 spaces instead of 4. 

+ Please conform to the [Google Python Style Guide](https://github.com/google/styleguide/blob/gh-pages/pyguide.md), and use [pylint](https://www.pylint.org/) to check your Python changes.

## Code linting
+ `black` is a very unforgiving code formatter. It doesn’t have any configuration options, and it has a very specific style. Code formatters will change the code automatically to meet a collection of style and layout practices.

    Note: `black` requires Python 3.6+.

+ Install black via pip:
    ```bash
    $ pip install black
    ```
    Then to run black at the command line, provide the file or directory you want to format:
    ```bash
    $ black test.py     
    ```

## Paths for importing modules and for opening files
+ Strictly do not perform relative import for any files or modules in Python. 
+ Always perform absolute import. 
    1. Assume the following project structure.
        ```text
        emotion-recognition                         # Main project directory 
        ├── model                                   # Machine learning model folder 
        |   ├── facial_expression_model_weights.h5  # Machine learning saved weights
        |   └── facial-expression-recognition.py    # Python code 
        ├── tfgraph                                 # TensorFlow graph folder 
        |   └── cnn.py                              # Graph architecture
        └── .env                                    # Environment variables 
        ```
    1. First, install the `dotenv` library.
        ```python
        pip install -U python-dotenv
        ```
    1. Specify the `ROOT` environment variable, which refers to the directory path containing the project, in the `.env` file as follows.
        ```
        # File: emotion-recognition/.env
        ROOT = /home/admin/src/github.com/emotion-recognition
        ```
    1. Place the following piece of code at the top of the Python code file. It will add the project directory given by `ROOT` to `sys.path` and make it searchable by Python. 
        ```python
        # File: emotion-recognition/model/facial-expression-recognition.py
        # Setup
        import os
        import sys
        from dotenv import load_dotenv, find_dotenv
        load_dotenv(find_dotenv())
        ROOT = os.getenv("ROOT")
        sys.path.append(ROOT)
        ```
    1. Import local packages as follows.
        ```python
        # File: emotion-recognition/model/facial-expression-recognition.py
        # Import local packages
        from tfgraph import cnn
        ```
    1. To open files in Python, create absolute paths by adding the file path within the project directory and the `ROOT`. An example is as follows. 
        ```python
        # File: emotion-recognition/model/facial-expression-recognition.py
        model.load_weights(ROOT+"/model/facial_expression_model_weights.h5")  # Load weights for TensorFlow model
        ```

## Dependency management   
+ To generate `requirements.txt` file for Python dependencies.
    ```python
    pipreqs [options] <path to main project folder>

    [options]
        --force : to overwrite existing file
    ```
    Example command assuming the project is located at `home/admin/src/github.com/scalable-deployment/tfsemonet` . 
    ```
    pipreqs --force /home/admin/src/github.com/scalable-deployment/tfsemonet
    ```

+ To install dependencies, issue the following commands.
    ```python
    pip install -r /requirements.txt
    ```