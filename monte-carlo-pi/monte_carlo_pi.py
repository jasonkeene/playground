#!/usr/bin/env python
"""Naive Monte Carlo approximation of pi."""

from __future__ import division

import math
from random import random
from itertools import repeat


ITERATIONS = 10 ** 7


def in_circle(x, y):
    """Determine if a given point is inside the unit circle."""
    x, y = abs(x), abs(y)
    return y <= math.sqrt(1.0 - x ** 2.0)


def main(iterations=ITERATIONS):
    """Approximate pi and print the result."""
    count = 0
    for x in repeat(None, iterations):
        if in_circle(random(), random()):
            count += 1
    print 'iterations:', iterations
    print '    approx:', count / iterations * 4
    print '   math.pi:', math.pi


if __name__ == '__main__':
    main()
