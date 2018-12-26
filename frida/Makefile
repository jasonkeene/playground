
all: frida

bin:
	mkdir -p bin

frida: bin
	cd tracee && go build -o ../bin/tracee -gcflags '-N -l'

measure:
	./measure.py `nm bin/tracee | grep -E '[^\.]main.f' | awk '{print $1}'`

clean:
	rm -r bin
