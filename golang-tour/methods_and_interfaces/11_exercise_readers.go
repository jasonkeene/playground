package main

import "fmt"

type MyReader struct {}

func (r *MyReader) Read(b []byte) (int, error) {
    i := 0
    for i = range b {
        b[i] = 'A'
    }
    return i + 1, nil
}

func main() {
    r := new(MyReader)
    b := make([]byte, 8)
    for {
        n, err := r.Read(b)
        if err != nil {
            fmt.Println("An error occured: %s", err)
        }
        for i := 0; i < n; i++ {
            fmt.Printf("%c", b[i])
        }
        print("\n")
    }
}
