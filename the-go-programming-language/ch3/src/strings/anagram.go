package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Anagram("cba", "abc"))
	fmt.Println(Anagram("cba", "cba"))

	fmt.Println(Anagram("cba", "cbaa"))
	fmt.Println(Anagram("cba", "caa"))
	fmt.Println(Anagram("cba", "ccc"))
}

func Anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for _, r := range s1 {
		index := strings.IndexRune(s2, r)
		if index < 0 {
			return false
		}
		s2 = s2[:index] + s2[index+1:]
	}
	return true
}
