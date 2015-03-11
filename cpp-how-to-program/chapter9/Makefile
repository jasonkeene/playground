all: bin/fig09_03 bin/copy

clean:
	[ -d bin ] && rm -r bin || exit 0

bin:
	mkdir bin

bin/fig09_03: bin fig09_03.cpp bin/Time.o
	g++ -std=c++11 fig09_03.cpp bin/Time.o -o bin/fig09_03

bin/Time.o: bin Time.cpp
	g++ -std=c++11 -c Time.cpp -o bin/Time.o

bin/copy: bin copy.cpp
	g++ -std=c++11 copy.cpp -o bin/copy
