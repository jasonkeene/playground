#!/usr/bin/env python

import logging

import zerorpc

import rando


def main():
    logging.basicConfig()
    server = zerorpc.Server(rando)
    server.bind("tcp://0.0.0.0:12345")
    server.run()


if __name__ == "__main__":
    main()
