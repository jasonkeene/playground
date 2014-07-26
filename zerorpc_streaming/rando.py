import hashlib
import itertools
from random import SystemRandom
import time

import gevent
import zerorpc


def random_hash():
    return hashlib.sha1(str(SystemRandom().random())).hexdigest()


@zerorpc.stream
def stream_randomness():
    for i in itertools.count():
        rando = random_hash()
        print "yielding", rando
        yield rando
        if i == 7:
            print "simulating network outage"
            time.sleep(5)
        gevent.sleep(0.5)
