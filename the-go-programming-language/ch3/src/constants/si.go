package main

import "fmt"

const (
	KB = 1000
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB
)

func main() {
	fmt.Println("KB =", KB)
	fmt.Println("MB =", MB)
	fmt.Println("GB =", GB)
	fmt.Println("TB =", TB)
	fmt.Println("PB =", PB)
	fmt.Println("EB =", EB)
}
