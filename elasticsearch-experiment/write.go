package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"
)

var (
	url     = flag.String("url", "", "elastic search api url")
	type_   = flag.String("type", "", "the type of document to insert")
	index   = flag.String("index", "", "the index to insert the document into")
	message = flag.String("message", "", "message to insert into elastic search")
)

const urlFmt = "%s/%s/%s"

func fatal(err error) {
	if err != nil {
		panic(err)
	}
}

func Insert(document interface{}) {
	documentJSON, err := json.Marshal(document)
	fatal(err)
	postURL := fmt.Sprintf(urlFmt, *url, *index, *type_)
	request, err := http.NewRequest("POST", postURL, bytes.NewReader(documentJSON))
	fatal(err)
	_, err = http.DefaultClient.Do(request)
	fatal(err)
}

func main() {
	flag.Parse()

	Insert(map[string]interface{}{
		"time":    time.Now().UnixNano(),
		"message": message,
	})
}
