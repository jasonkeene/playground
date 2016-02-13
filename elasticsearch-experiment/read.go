package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	url   = flag.String("url", "", "elastic search api url")
	type_ = flag.String("type", "", "the type of document to insert")
	index = flag.String("index", "", "the index to insert the document into")
)

const urlFmt = "%s/%s/%s/_search?pretty"

func fatal(err error) {
	if err != nil {
		panic(err)
	}
}

func Search() {
	getURL := fmt.Sprintf(urlFmt, *url, *index, *type_)
	println(getURL)
	request, err := http.NewRequest("GET", getURL, nil)
	fatal(err)
	resp, err := http.DefaultClient.Do(request)
	fatal(err)
	data, err := ioutil.ReadAll(resp.Body)
	fatal(err)
	fmt.Println(string(data))
}

func main() {
	flag.Parse()

	Search()
}
