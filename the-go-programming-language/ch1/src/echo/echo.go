package echo

import (
	"fmt"
	"strings"
)

func Echo1(fragments []string) {
	var s, sep string
	for _, f := range fragments {
		s += sep + f
		sep = " "
	}
	fmt.Println(s)
}

func Echo2(fragments []string) {
	fmt.Println(strings.Join(fragments, " "))
}
