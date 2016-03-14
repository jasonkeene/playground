package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	var d []byte

	d = []byte("a    bb")
	d = squashSpace(d)
	fmt.Println(string(d))

	d = []byte("  \t \n a  \n\n  b")
	d = squashSpace(d)
	fmt.Println(string(d))
}

func squashSpace(data []byte) []byte {
	write := 0
	read := 0
	space := false

	for read < len(data) {
		r, s := utf8.DecodeRune(data[read:])
		read += s

		if unicode.IsSpace(r) {
			space = true
		} else {
			if space {
				space = false
				data[write] = byte(' ')
				write++
			}
			write += utf8.EncodeRune(data[write:], r)
		}
	}
	return data[:write]
}
