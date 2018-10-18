package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	for {
		text, err := in.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		n, err := strconv.Atoi(strings.TrimRight(text, "\n"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d\n", n+2)
	}
}
