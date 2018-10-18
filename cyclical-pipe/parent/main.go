package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	cmd := exec.Command(os.Getenv("CHILD"))
	childIn, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer childIn.Close()
	childOut, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer childOut.Close()

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	// prime the pump
	fmt.Fprintln(childIn, "12")

	in := bufio.NewReader(childOut)

	for {
		text, err := in.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		n, err := strconv.Atoi(strings.TrimRight(text, "\n"))
		if err != nil {
			log.Fatal(err)
		}
		new := n + 2
		log.Println("sending new value:", new)
		fmt.Fprintf(childIn, "%d\n", new)
		time.Sleep(time.Second)
	}
}
