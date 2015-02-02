package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["testers1"] = 1
	m["testers2"] = 2

	fmt.Println("map:", m)

	v1 := m["testers1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "testers2")
	fmt.Println("map:", m)

	v2, present := m["testers2"]
	fmt.Println("present:", present)
	fmt.Println("v2:", v2)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)
	fmt.Printf("%#v\n", n)
}
