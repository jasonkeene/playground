package main

import (
	"sync/atomic"
	"time"

	"golang.org/x/text/message"
)

var (
	count uint64
	p     = message.NewPrinter(message.MatchLanguage("en"))
)

func f() {
	atomic.AddUint64(&count, 1)
}

func reportCounters() {
	for {
		c := atomic.SwapUint64(&count, 0)
		p.Printf("%15d ops/s\n", c)
		time.Sleep(time.Second)
	}
}

func main() {
	go reportCounters()
	for {
		f()
	}
}
