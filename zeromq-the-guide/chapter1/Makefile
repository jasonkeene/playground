
all: hello_server hello_client version

bin:
	[ ! -d bin ] && mkdir bin

hello_server: bin
	clang -Wall -Wno-unused-function hello_server.c -o bin/hello_server -L/opt/local/lib -I/opt/local/include -lzmq

hello_client: bin
	clang -Wall -Wno-unused-function hello_client.c -o bin/hello_client -L/opt/local/lib -I/opt/local/include -lzmq

version: bin
	clang -Wall -Wno-unused-function version.c -o bin/version -L/opt/local/lib -I/opt/local/include -lzmq

clean:
	rm -r bin
