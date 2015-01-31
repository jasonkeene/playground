package main

import "fmt"

func main() {
	// basic
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// classic
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// while
	for {
		fmt.Println("loop")
		break
	}
}
