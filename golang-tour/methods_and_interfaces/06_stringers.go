package main

import "fmt"

type Person struct {
    Name string
    Age int
}

func (p Person) String() string {
    return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
    a := Person{"Nicole Keene", 28}
    z := Person{"Jason Keene", 30}
    fmt.Println(a)
    fmt.Println(z)
}
