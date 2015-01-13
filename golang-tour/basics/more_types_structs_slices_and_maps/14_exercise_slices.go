package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
    parent := make([][]uint8, dy)
    for i := range parent {
        row := make([]uint8, dx)
        for j := range row {
            row[j] = uint8(i) * uint8(j)
        }
        parent[i] = row
    }
    return parent
}

func main() {
    pic.Show(Pic)
}
