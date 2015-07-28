package main

import (
	"flag"
	"fmt"
	"strconv"
)

var n = flag.Int("n", 1, "count of how many primes to calculate")

func Primes() chan int {
	out := make(chan int)
	in := NewSieve(out)
	go func() {
		for i := 2; ; i++ {
			in <- i
		}
	}()
	return out
}

func NewSieve(out chan int) chan int {
	in := make(chan int)
	go sieve(in, out)
	return in
}

func sieve(in, out chan int) {
	prime := <-in
	out <- prime

	child_in := NewSieve(out)
	for {
		possible_prime := <-in
		if possible_prime%prime != 0 {
			child_in <- possible_prime
		}
	}
}

func main() {
	flag.Parse()
	prime_chan := Primes()
	for i := 0; i < *n; i++ {
		fmt.Println(strconv.Itoa(<-prime_chan))
	}
}
