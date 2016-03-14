package main

import (
	"fmt"
	"unicode/utf8"
)

const size = 5

func main() {
	a := [size]int{1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a)

	b := []byte("Hello, 世界")
	reverseUTF8(b)
	fmt.Println(string(b))
}

func reverse(a *[size]int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func reverseUTF8(d []byte) {
	tmp := make([]byte, len(d))
	read := len(d)
	write := 0

	for read > 0 {
		r, s := utf8.DecodeLastRune(d[:read])
		read -= s
		write += utf8.EncodeRune(tmp[write:], r)
	}

	copy(d, tmp)
}
