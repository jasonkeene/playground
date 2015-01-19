package main

import (
    "fmt"
    "log"
    "net/http"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s\n", s)
}

type Struct struct {
    Greeting string
    Punct string
    Who string
}

func (s Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s%s%s\n", s.Greeting, s.Punct, s.Who)
}

func main() {
    http.Handle("/string", String("I'm a frayed knot."))
    http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
    log.Fatal(http.ListenAndServe("localhost:54321", nil))
}
