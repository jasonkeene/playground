
all: hello version weather
.PHONY: all

bin:
	[ ! -d bin ] && mkdir bin

lib:
	[ ! -d lib ] && mkdir lib

hello: hello_server hello_client
.PHONY: hello

hello_server: bin
	clang -Wall -Wno-unused-function hello_server.c -o bin/hello_server -L/opt/local/lib -I/opt/local/include -lzmq

hello_client: bin
	clang -Wall -Wno-unused-function hello_client.c -o bin/hello_client -L/opt/local/lib -I/opt/local/include -lzmq

version: bin
	clang -Wall -Wno-unused-function version.c -o bin/version -L/opt/local/lib -I/opt/local/include -lzmq

weather: weather_server weather_client
.PHONY: weather

weather_lib: lib
	clang -Wall -c weather.c -o lib/weather.o

weather_server: bin weather_lib
	clang -Wall -Wno-unused-function lib/weather.o weather_server.c -o bin/weather_server -L/opt/local/lib -I/opt/local/include -lzmq

weather_client: bin weather_lib
	clang -Wall -Wno-unused-function lib/weather.o weather_client.c -o bin/weather_client -L/opt/local/lib -I/opt/local/include -lzmq

clean:
	rm -r bin lib
