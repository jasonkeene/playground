
all: client server
.PHONY: all

bin:
	[ ! -d bin ] && mkdir bin

lib:
	[ ! -d lib ] && mkdir lib

client: bin
	clang -Wall -Wno-unused-function client.c -o bin/client -L/opt/local/lib -I/opt/local/include -lzmq

server: bin
	clang -Wall -Wno-unused-function server.c -o bin/server -L/opt/local/lib -I/opt/local/include -lzmq

clean:
	rm -r bin lib
