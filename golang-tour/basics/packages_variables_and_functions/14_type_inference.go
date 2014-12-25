package main

import "fmt"

func main() {
    a := 42
    b := 42.42
    c := 42i
    fmt.Printf("a is of type %T\n", a)
    fmt.Printf("b is of type %T\n", b)
    fmt.Printf("c is of type %T\n", c)
}
