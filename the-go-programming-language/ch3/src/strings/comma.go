package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println(Comma("1234.5"))
	fmt.Println(Comma(".5"))
	fmt.Println(Comma("1234"))
	fmt.Println(Comma("123"))
	fmt.Println(Comma("123456"))
	fmt.Println(Comma("-123456"))
	fmt.Println(Comma("+123456"))
}

func Comma(s string) string {
	var b bytes.Buffer

	sign := 0
	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		sign = 1
	}
	b.WriteString(s[:sign])

	period := strings.LastIndex(s, ".")
	if period < 0 {
		period = len(s)
	}

	lhs := s[sign:period]
	rc := utf8.RuneCountInString(lhs)
	for i, r := range lhs {
		if i != 0 && (i-rc)%3 == 0 {
			b.WriteByte(',')
		}
		b.WriteRune(r)
	}

	b.WriteString(s[period:])
	return b.String()
}
