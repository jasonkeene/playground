package main

import (
    "fmt"
    "time"
    "sync"
)

func say(s string, wg *sync.WaitGroup) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
    wg.Done()
}

func main() {
    wg := new(sync.WaitGroup)
    wg.Add(2)
    go say("world", wg)
    say("hello", wg)
    wg.Wait()
}
