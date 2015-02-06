package main

import (
	"fmt"
	"sync"
)

type Thing struct{}

var pool = sync.Pool{
	New: func() interface{} {
		fmt.Println("creating new thing!")
		count++
		return new(Thing)
	},
}

var count = 0

func main() {
	t1, t2, t3 := pool.Get().(*Thing), pool.Get().(*Thing), pool.Get().(*Thing)
	pool.Put(t1)
	pool.Put(t2)
	t4, t5, t6 := pool.Get().(*Thing), pool.Get().(*Thing), pool.Get().(*Thing)
	pool.Put(t3)
	pool.Put(t4)
	pool.Put(t5)
	pool.Put(t6)
	fmt.Printf("%d Things created!\n", count)
}
