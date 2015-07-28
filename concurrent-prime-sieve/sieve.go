package main

import (
	"flag"
	"fmt"
	"strconv"
)

var n = flag.Int("n", 1, "count of how many primes to calculate")

func Primes() chan int {
	prime_chan := make(chan int)
	go func() {
		for {
			prime_chan <- 2
		}
	}()
	return prime_chan
}

func main() {
	flag.Parse()
	prime_chan := Primes()
	for i := 0; i < *n; i++ {
		fmt.Println(strconv.Itoa(<-prime_chan))
	}
}
