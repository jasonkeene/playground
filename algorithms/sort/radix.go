package sort

import (
	"log"
)

func Radix(a []string, size int) {
	for i := size - 1; i >= 0; i-- {
		counting(a, i)
	}
}

func counting(a []string, p int) {
	// equal
	bucket := make([]int, 36)
	for _, v := range a {
		bucket[toIndex(v[p])]++
	}

	// less
	tmp := bucket[0]
	bucket[0] = 0
	for i := 1; i < len(bucket); i++ {
		tmp, bucket[i] = bucket[i], tmp+bucket[i-1]
	}

	// sort
	sorted := make([]string, len(a))
	for _, v := range a {
		i := bucket[toIndex(v[p])]
		sorted[i] = v
		bucket[toIndex(v[p])]++
	}

	// replace
	for i, v := range sorted {
		a[i] = v
	}
}

func toIndex(r byte) int {
	// digit
	if r >= 48 && r < 58 {
		return int(r) - 48
	}
	// capital letter
	if r >= 65 && r < 91 {
		return int(r) - 55
	}
	log.Panicf("bad byte: %d", r)
	return -1
}
