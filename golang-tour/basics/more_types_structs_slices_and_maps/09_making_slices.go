package main

import "fmt"

func main() {
    a := make([]int, 5)
    a = []int{0, 1, 2, 3, 4}
    printSlice("a", a)
    b := make([]int, 0, 6)
    b = []int{5, 6, 7, 8, 9}
    printSlice("b", b)
    c := b[:2]
    printSlice("c", c)
    d := c[2:5]
    printSlice("d", d)
}

func printSlice(s string, x []int) {
    fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
