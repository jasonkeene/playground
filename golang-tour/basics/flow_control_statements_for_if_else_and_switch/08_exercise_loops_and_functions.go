package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) (float64, int) {
    y, z := x, x
    i := 0
    for {
        i++
        z = z - (z * z - x) / (2 * z);
        if math.Abs(y - z) < 1e-15 {
            break
        }
        y = z
    }
    return z, i
}

func main() {
    fmt_str := "%d: %g (%d) - %g = %g\n"
    for i := 1; i < 100; i++ {
        value, iterations := Sqrt(float64(i))
        actual := math.Sqrt(float64(i))
        diff := value - actual
        fmt.Printf(fmt_str, i, value, iterations, actual, diff)
    }
}
