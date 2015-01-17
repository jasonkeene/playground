package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    // I decided to use Sprintf here vs Sprint, however to answer the question
    // posed in the exercise, the reason Sprint would recurse forever is
    // because it probably finds the Error method on e as the default format
    // method and calls it, then it calls Sprint again, and so on..
    return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
    if x < 0 {
        return 0.0, ErrNegativeSqrt(x)
    }
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
    return z, nil
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
