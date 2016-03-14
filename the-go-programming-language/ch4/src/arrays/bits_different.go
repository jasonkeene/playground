package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func BitsDifferent(h1, h2 [sha256.Size]byte) int {
	diff := 0
	for i := 0; i < sha256.Size; i++ {
		diff += PopCount(uint64(h1[i] ^ h2[i]))
	}
	return diff
}

func main() {
	h1, h2 := sha256.Sum256([]byte("a")), sha256.Sum256([]byte("A"))
	fmt.Printf("%d bits different\n", BitsDifferent(h1, h2))
}
