package main

import "fmt"

func main() {
	s := make([]string, 5)
	fmt.Println("emp:", s)

	var t []int
	t = append(t, 1)
	t = append(t, 2)
	t = append(t, 3)
	fmt.Println("t:", t)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s[3] = "d"
	s[4] = "e"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:4]
	fmt.Println("sl1:", l)
	l = s[:5]
	fmt.Println("sl2:", l)
	l = s[2:]
	fmt.Println("sl3:", l)
	l = s[:]
	fmt.Println("sl4:", l)

	u := []string{"g", "h", "i"}
	fmt.Println("dcl:", u)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerlen := i + 1
		twoD[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
