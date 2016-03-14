package main

import "fmt"

func main() {
	values := []int{1, 2, 3, 4, 5}
	for i := 0; i < 15; i++ {
		s := make([]int, len(values))
		copy(s, values)
		rotate(s, i)
		fmt.Printf("rotated %d times: %v\n", i, s)
	}
}

func rotate(s []int, c int) {
	c = c % len(s)
	if c == 0 {
		return
	}
	var ac []int
	for i, v := range s {
		ac = append(ac, v)
		if i < c {
			// look behind
			s[i] = s[len(s)-c+i]
		} else {
			// copy from ac
			s[i] = ac[0]
			ac = ac[1:]
		}
	}
}
