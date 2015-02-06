package main

import "fmt"

type Thingy struct{}

func main() {
	ch := make(chan interface{})
	go func() {
		ch <- 1
		ch <- 1.1
		ch <- "test"
		ch <- true
		ch <- new(Thingy)
		ch <- Thingy{}
		close(ch)
	}()
	for x := range ch {
		switch x.(type) {
		case bool:
			fmt.Println("bool")
		case string:
			fmt.Println("string")
		case int:
			fmt.Println("int")
		case float64:
			fmt.Println("float64")
		default:
			fmt.Printf("unknown type %T\n", x)
		}
	}
}
