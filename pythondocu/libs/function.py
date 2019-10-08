"""
function.py
====================================
The core module of my example project
"""

# TODO(kl@email.com): Use a "*" here for string repetition.
# TODO(Zeke) Change this to use relations.

def fetch_ml_model(big_value, keys, other_variable=None):
    """Fetches rows from a ML model.

    Retrieves rows pertaining to the given keys from the Table instance
    represented by ML model.

    Args:
        big_table: An open ML model Table instance.
        keys: A sequence of strings representing the key of each table row
            to fetch.

    Returns:
        A dict mapping keys to the corresponding table row data
        fetched. Each row is represented as a tuple of strings.

    Raises:
        IOError: An error occurred accessing the ML model.Table object.
    """
    return "The wise {} loves Python.".format(big_value)
