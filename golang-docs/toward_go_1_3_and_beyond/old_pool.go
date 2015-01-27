package main

import "fmt"

type Thing struct{}

var pool = make(chan *Thing, 10)

var count = 0

func GetThing() *Thing {
	select {
	case t := <-pool:
		return t
	default:
	}
	fmt.Println("creating new thing!")
	count++
	return new(Thing)
}

func PutThing(t *Thing) {
	pool <- t
}

func main() {
	t1, t2, t3 := GetThing(), GetThing(), GetThing()
	PutThing(t1)
	PutThing(t2)
	t4, t5, t6 := GetThing(), GetThing(), GetThing()
	PutThing(t3)
	PutThing(t4)
	PutThing(t5)
	PutThing(t6)
	fmt.Printf("%d Things created!\n", count)
}
