package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
    n := len(b)
    tmp := make([]byte, n)
    for {
        // read out of source
        count, err := r.r.Read(tmp)

        // error checking on read
        if err == io.EOF {
            return 0, io.EOF
        } else if err != nil {
            return 0, err
        }

        // map characters
        rot13(tmp)

        // write out to dest
        copy(b, tmp)
        return count, nil
    }
}

// rot13 a given slice of bytes
func rot13(b []byte) {
    for i := range b {
        if b[i] >= 65 && b[i] <= 90 {
            // byte is A-Z
            b[i] = (b[i] + 13 - 65) % 26 + 65
        } else if b[i] >= 97 && b[i] <= 122 {
            // byte is a-z
            b[i] = (b[i] + 13 - 97) % 26 + 97
        }
    }
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
    print("\n")
}
