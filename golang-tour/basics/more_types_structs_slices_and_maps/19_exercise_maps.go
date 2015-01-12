package main

import "strings"

import "code.google.com/p/go-tour/wc"

func WordCount(s string) map[string]int {
    counter := make(map[string]int)
    for _, word := range strings.Fields(s) {
        counter[word]++
    }
    return counter
}

func main() {
    wc.Test(WordCount)
}
