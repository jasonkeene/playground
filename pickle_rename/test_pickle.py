#!/usr/bin/env python

import pickle
import sys

from test_lib import Foo


DBNAME = 'pickle_data'
COMMANDS = {}


def command(func):
    COMMANDS[func.__name__] = func
    return func


@command
def write(*name):
    """Write foo object from db file."""
    foo = Foo(' '.join(name))
    with open(DBNAME, 'w') as f:
        f.write(pickle.dumps(foo))


@command
def read():
    """Read foo object from db file."""
    with open(DBNAME) as f:
        foo = pickle.loads(f.read())
        print foo


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print "must choose from commands: " + ', '.join(COMMANDS.keys())
    else:
        COMMANDS[sys.argv[1]](*sys.argv[2:])
