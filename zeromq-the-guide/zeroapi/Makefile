
all: client server
.PHONY: all

bin:
	[ ! -d bin ] && mkdir bin

lib:
	[ ! -d lib ] && mkdir lib

utils: lib
	clang -Wall -c utils.c -o lib/utils.o

client: bin utils
	clang -Wall -Wno-unused-function lib/utils.o client.c -o bin/client -L/opt/local/lib -I/opt/local/include -lzmq

server: bin utils
	clang -Wall -Wno-unused-function lib/utils.o server.c -o bin/server -L/opt/local/lib -I/opt/local/include -lzmq

clean:
	rm -r bin lib
