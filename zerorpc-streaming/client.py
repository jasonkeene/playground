#!/usr/bin/env python

import zerorpc
import time


def main():
    client = zerorpc.Client()
    client.connect("tcp://127.0.0.1:12345")
    for i, item in enumerate(client.stream_randomness()):
        print "got", item
        if i == 3:
            print "simulating network outage"
            time.sleep(5)
        elif i == 9:
            print "done, got", i + 1, "messages"
            client.close()
            break


if __name__ == "__main__":
    main()
