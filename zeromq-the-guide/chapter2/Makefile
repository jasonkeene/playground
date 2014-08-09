
all: ventilator worker sink rr
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

rr: bin
	clang -Wall -Wno-unused-function rrclient.c -o bin/rrclient -L/opt/local/lib -I/opt/local/include -lzmq
	clang -Wall -Wno-unused-function rrbroker.c -o bin/rrbroker -L/opt/local/lib -I/opt/local/include -lzmq
	clang -Wall -Wno-unused-function rrworker.c -o bin/rrworker -L/opt/local/lib -I/opt/local/include -lzmq

clean:
	rm -r bin lib
