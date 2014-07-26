
all: hello_server hello_client

bin:
	[ ! -d bin ] && mkdir bin

hello_server: bin
	clang -Wall hello_server.c -o bin/hello_server -L/opt/local/lib -I/opt/local/include -lzmq

hello_client: bin
	clang -Wall hello_client.c -o bin/hello_client -L/opt/local/lib -I/opt/local/include -lzmq

clean:
	rm -r bin
