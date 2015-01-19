package main

import (
    "fmt"
    "log"
    "net/http"
)

type HelloServer struct {}

func (h *HelloServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("%s %s\n", r.Method, r.URL)
    fmt.Fprintf(w, "Requested %s\n", r.URL)
    for i := 0; i < 2000000; i++ {
        fmt.Fprintf(w, "%d\n", i)
    }
}

func main() {
    listen := "localhost:8123"
    server := new(HelloServer)
    fmt.Printf("Starting server listening on http://%s/\n", listen)
    err := http.ListenAndServe(listen, server)
    if err != nil {
        log.Fatal(err)
    }
}
