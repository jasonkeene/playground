
all: ventilator worker sink
.PHONY: all

bin:
	[ ! -d bin ] && mkdir bin

lib:
	[ ! -d lib ] && mkdir lib

ventilator: bin
	clang -Wall -Wno-unused-function ventilator.c -o bin/ventilator -L/opt/local/lib -I/opt/local/include -lzmq

worker: bin
	clang -Wall -Wno-unused-function worker.c -o bin/worker -L/opt/local/lib -I/opt/local/include -lzmq

sink: bin
	clang -Wall -Wno-unused-function sink.c -o bin/sink -L/opt/local/lib -I/opt/local/include -lzmq

clean:
	rm -r bin lib
