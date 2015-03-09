all: bin/Time.o

clean:
	[ -d bin ] && rm -r bin || exit 0

bin:
	mkdir bin

bin/Time.o: bin Time.cpp
	g++ -std=c++11 -c Time.cpp -o bin/Time.o
