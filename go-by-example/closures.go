package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		defer func() {
			i += 1
		}()
		return i
	}
}

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
