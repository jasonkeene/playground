package main

import "testing"

func TestKnownPrimes(t *testing.T) {
	prime_chan := Primes()
	known_primes := []int{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53,
		59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113,
	}
	for _, prime := range known_primes {
		got := <-prime_chan
		if prime != got {
			t.Errorf("bad prime, expected: %d, got: %d", prime, got)
		}
	}
}
