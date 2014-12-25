package main

import "fmt"

const (
    Big   = 1 << 63
    Small = Big >> 99
)

func needInt(x uint64) uint64 {
    return x
}

func needFloat(x float64) float64 {
    return x
}

func main() {
    fmt.Println(needInt(Small))
    fmt.Println(needInt(Big))
    fmt.Println(needFloat(Small))
    fmt.Println(needFloat(Big))
}
