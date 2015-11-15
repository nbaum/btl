CC=clang++ -std=c++14 -Wall
CFLAGS=-Iinclude

all: golem

golem: build/golem.o
	$(CC) $(LDFLAGS) -o $@ $+

build/%.o: src/%.cpp
	mkdir -p build
	$(CC) $(CFLAGS) -c -o $@ $+
