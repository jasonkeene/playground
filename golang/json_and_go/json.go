package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Message struct {
	Name string `json:"name"`
	Body string `json:"body"`
	Time int64  `json:"time"`
}

type FamilyMember struct {
	Name    string
	Age     int
	Parents []string
}

func json_test(v interface{}) {
	fmt.Printf("original type: %T\n", v)
	fmt.Printf("original value: %#v\n", v)

	bytes, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		fmt.Printf("An error occured in marshalling! %v\n\n", err)
		return
	}
	print("json bytes: ")
	os.Stdout.Write(bytes)
	print("\n")

	var result interface{}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		fmt.Printf("An error occured in unmarshalling! %v\n\n", err)
		return
	}
	fmt.Printf("result type: %T\n", result)
	fmt.Printf("result value: %#v\n\n", result)
}

func main() {
	json_test(nil)
	json_test(true)
	json_test(false)
	json_test(42)
	json_test(42.42)
	json_test("raw string")
	json_test("Hello, 世界")
	json_test([]interface{}{
		"raw string",
		42,
		map[string]string{
			"key": "value",
		},
	})

	m := Message{"Jason", "Hello friend!", 1294706395881547000}
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	var fm FamilyMember
	err = json.Unmarshal(b, &fm)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", fm)
}
