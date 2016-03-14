package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var algo string

func main() {
	flag.StringVar(&algo, "algo", "sha256", "specifies hashing algorithm to use")
	flag.Parse()

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "unable to read from stdin")
		os.Exit(1)
	}

	switch algo {
	case "sha1":
		hash := sha1.Sum(data)
		fmt.Printf("%#x\n", hash)
	case "sha256":
		hash := sha256.Sum256(data)
		fmt.Printf("%#x\n", hash)
	case "sha384":
		hash := sha512.Sum384(data)
		fmt.Printf("%#x\n", hash)
	case "sha512":
		hash := sha512.Sum512(data)
		fmt.Printf("%#x\n", hash)
	default:
		fmt.Fprintln(os.Stderr, "unsupported hashing algorithm specified")
		os.Exit(2)
	}
}
