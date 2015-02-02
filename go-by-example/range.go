package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index(3):", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	unistr := "也习乡乢乣乤乥书乧乨乩乪"
	for i, r := range unistr {
		fmt.Printf("unistr[%d]: %#U (%#v)\n", i, r, r)
	}
}
