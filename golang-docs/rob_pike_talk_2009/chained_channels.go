package main

import (
    "flag"
    "fmt"
)

var ngoroutine = flag.Int("n", 100000, "how many")

func add_chans(left, right chan int) {
    left <- 1 + <-right
}

func main() {
    flag.Parse()
    leftmost := make(chan int)
    var left, right chan int = nil, leftmost
    fmt.Println("creating channels")
    for i := 0; i < *ngoroutine; i++ {
        left, right = right, make(chan int)
        go add_chans(left, right)
    }
    fmt.Println("start computation")
    right <- 0
    fmt.Println(<-leftmost)
}
