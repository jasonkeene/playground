all: $(shell ls *.cpp | sed -E 's/(.*).cpp/bin\/\1/')

clean:
	[ -d bin ] && rm -r bin || exit 0

bin:
	mkdir bin

bin/fig08_04: bin fig08_04.cpp
	g++ -std=c++11 fig08_04.cpp -o bin/fig08_04

bin/fig08_06: bin fig08_06.cpp
	g++ -std=c++11 fig08_06.cpp -o bin/fig08_06

bin/fig08_07: bin fig08_07.cpp
	g++ -std=c++11 fig08_07.cpp -o bin/fig08_07

bin/nullptr: bin nullptr.cpp
	g++ -std=c++11 nullptr.cpp -o bin/nullptr

bin/arrays: bin arrays.cpp
	g++ -std=c++11 arrays.cpp -o bin/arrays

bin/sizeof: bin sizeof.cpp
	g++ -std=c++11 sizeof.cpp -o bin/sizeof

bin/cstrings: bin cstrings.cpp
	g++ -std=c++11 cstrings.cpp -o bin/cstrings
