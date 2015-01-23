package main

import (
    "fmt"
    "sort"
    "code.google.com/p/go-tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    ch <- t.Value
    if t.Left != nil {
        Walk(t.Left, ch)
    }
    if t.Right != nil {
        Walk(t.Right, ch)
    }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    ch1, ch2, done := make(chan int), make(chan int), make(chan int)
    r1, r2 := make([]int, 0), make([]int, 0)

    go func () {
        Walk(t1, ch1)
        done <- 1
    }()
    go func () {
        Walk(t2, ch2)
        done <- 1
    }()
    for n := 2; n > 0; {
        select {
        // these two cases are almost identical, I wish I knew how to
        // write a case that reads when both channels are ready
        case v1 := <-ch1:
            // this read could block indefinitely if t1 and t2
            // have different sizes
            v2 := <-ch2
            if v1 != v2 {
                // push these values into a slice for comparison later
                r1 = append(r1, v1)
                r2 = append(r2, v2)
            }
        case v2 := <-ch2:
            // this read could block indefinitely if t1 and t2
            // have different sizes
            v1 := <-ch1
            if v1 != v2 {
                // push these values into a slice for comparison later
                r1 = append(r1, v1)
                r2 = append(r2, v2)
            }
        case <-done:
            n--
        }
    }

    // if there are any remaining values, make sure they are all the same
    if len(r1) != 0 || len(r2) != 0 {
        if len(r1) != len(r2) {
            // this will never get hit as it would block indefinitely in the
            // select above
            return false
        }
        // sort slices to make them easy to compare
        is1 := sort.IntSlice(r1)
        is2 := sort.IntSlice(r2)
        sort.Sort(is1)
        sort.Sort(is2)
        for i, x := range is1 {
            if x != is2[i] {
                return false
            }
        }
    }
    return true
}

func main() {
    fmt.Println(Same(tree.New(4), tree.New(4)))
    fmt.Println(Same(tree.New(3), tree.New(4)))
}
