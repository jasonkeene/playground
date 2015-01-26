package main

import "fmt"

// it occured to me at 36:59 that it would be possible to have untyped channels
func main() {
    ch := make(chan interface {})
    go func () {
        ch <- 1234
        ch <- "asdf"
        ch <- 1.3
        ch <- struct {x, y int}{22, 33}
        close(ch)
    }()
    for val := range ch {
        fmt.Println(val)
    }
}
