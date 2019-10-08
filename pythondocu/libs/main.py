"""
main.py
====================================
The core module of my example project
"""


def fetch_data_rows(ml_layer, keys, other_variable=None):
    """Fetches rows from a data.

    Retrieves rows pertaining to the given keys from the Table instance
    represented by ml_layer. 

    Args:
        ml_layer: An open data Table instance.
        keys: A sequence of strings representing the key of each table row
            to fetch.
        other_variable: Another optional variable, that has a much
            longer name than the other args.

    Returns:
        dict mapping keys to the corresponding table row data
        fetched. Each row is represented as a tuple of strings. 

        {'Layer1': ('CNN', 'LSTM'),
        'Layer2': ('Ridge', 'Regression')}

    Raises:
        IOError: An error occurred accessing the data.Table object.
    """
    return "The wise {} loves Python.".format(ml_layer)

class SampleClass(object):
    """Summary of class here.

    Longer class information....
    Longer class information....

    Attributes:
        likes_spam: A boolean indicating if we like spam or not.
        eggs: An integer count of the eggs.
    """

    def __init__(self, likes_spam=False):
        """Inits SampleClass with bla."""
        self.likes_spam = likes_spam
        self.eggs = 0

    def public_method(self):
        """Performs operation multiplication."""    
        return "I am a {} object.".format(self.eggs)