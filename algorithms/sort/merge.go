package sort

import "math"

func Merge(a []int) {
	m(a, 0, len(a))
}

func m(a []int, q, r int) {
	if q >= r-1 {
		return
	}
	mid := (q + r) / 2
	m(a, q, mid)
	m(a, mid, r)
	merge(a, q, mid, r)
}

func merge(a []int, q, mid, r int) {
	a1 := make([]int, mid-q+1)
	for i, j := q, 0; i < mid; i, j = i+1, j+1 {
		a1[j] = a[i]
	}
	a1[mid-q] = math.MaxInt64

	a2 := make([]int, r-mid+1)
	for i, j := mid, 0; i < r; i, j = i+1, j+1 {
		a2[j] = a[i]
	}
	a2[r-mid] = math.MaxInt64

	var a1c, a2c int
	for i := q; i < r; i++ {
		if a1[a1c] < a2[a2c] {
			a[i] = a1[a1c]
			a1c++
		} else {
			a[i] = a2[a2c]
			a2c++
		}
	}
}
