package main

import "fmt"

func world() string {
    fmt.Println("Generated world")
    return "world"
}

func main() {
    defer fmt.Println(world())
    fmt.Println("hello")
}
