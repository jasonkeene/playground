package main

import "fmt"

func main() {
    var a []int
    printSlice("a", a)

    // append stuff to nil slice
    a = append(a, 0)
    printSlice("a", a)

    // the slice grows!
    a = append(a, 1)
    printSlice("a", a)

    // append is variadic
    a = append(a, 2, 3, 4)
    printSlice("a", a)
}

func printSlice(s string, x []int) {
    fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
