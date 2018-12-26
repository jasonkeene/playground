#!/usr/bin/env python

from __future__ import print_function
import frida
import sys
import time

session = frida.attach("tracee")
script = session.create_script("""
var count = 0;

setInterval(function () {
    send(count);
    count = 0;
}, 1000);

Interceptor.attach(ptr("%s"), {
    onEnter: function(args) {
        count++;
    }
});
""" % int(sys.argv[1], 16))

def on_message(message, data):
    print(message)

script.on('message', on_message)

script.load()

try:
    time.sleep(99999999)
except KeyboardInterrupt:
    pass
