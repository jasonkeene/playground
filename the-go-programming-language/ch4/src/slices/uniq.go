package main

import "fmt"

func main() {
	s1 := []string{"a", "a", "b", "b"}
	s1 = uniq(s1)
	fmt.Println(s1)

	s2 := []string{"a", "a", "b", "b", "c", "b"}
	s2 = uniq(s2)
	fmt.Println(s2)

	s3 := []string{}
	s3 = uniq(s3)
	fmt.Println(s3)

	s4 := []string{"a"}
	s4 = uniq(s4)
	fmt.Println(s4)

	s5 := []string{"a", "b"}
	s5 = uniq(s5)
	fmt.Println(s5)

	s6 := []string{"a", "a"}
	s6 = uniq(s6)
	fmt.Println(s6)
}

func uniq(s []string) []string {
	if len(s) == 0 {
		return s
	}
	last := s[0]
	dropped := 0
	for i, v := range s[1:] {
		if v == last {
			dropped++
		}
		s[i+1-dropped], last = v, v
	}
	return s[:len(s)-dropped]
}
