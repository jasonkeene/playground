package main

import (
    "fmt"
)

type Fetcher interface {
    // Fetch returns the body of URL and
    // a slice of URLs found on that page.
    Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, signal chan int,
           visited *[]string) {
    // keep track of visited urls
    for _, s := range *visited {
        if url == s {
            fmt.Println("Already visited url ", url)
            signal <- -1
            return
        }
    }
    *visited = append(*visited, url)

    // check if max depth exceeded
    if depth <= 0 {
        signal <- -1
        return
    }

    // do the fetch
    body, urls, err := fetcher.Fetch(url)
    if err != nil {
        fmt.Println(err)
        signal <- -1
        return
    }
    fmt.Printf("found: %s %q\n", url, body)

    // crawl children
    for _, u := range urls {
        signal <- 1
        go Crawl(u, depth-1, fetcher, signal, visited)
    }

    signal <- -1
    return
}

func main() {
    signal := make(chan int)
    visited := make([]string, 0)
    go Crawl("http://golang.org/", 4, fetcher, signal, &visited)
    for n := 1; n > 0; {
        select {
        case v := <-signal:
            n += v
        }
    }
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
    body string
    urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
    if res, ok := f[url]; ok {
        return res.body, res.urls, nil
    }
    return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
    "http://golang.org/": &fakeResult{
        "The Go Programming Language",
        []string{
            "http://golang.org/pkg/",
            "http://golang.org/cmd/",
        },
    },
    "http://golang.org/pkg/": &fakeResult{
        "Packages",
        []string{
            "http://golang.org/",
            "http://golang.org/cmd/",
            "http://golang.org/pkg/fmt/",
            "http://golang.org/pkg/os/",
        },
    },
    "http://golang.org/pkg/fmt/": &fakeResult{
        "Package fmt",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
    "http://golang.org/pkg/os/": &fakeResult{
        "Package os",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
}
