package sort

func Quick(a []int) {
	quick(a, 0, len(a))
}

func quick(a []int, q, r int) {
	if len(a[q:r]) <= 1 {
		return
	}
	pivot := partition(a, q, r)
	quick(a, q, pivot)
	quick(a, pivot+1, r)
}

func partition(a []int, q, r int) int {
	rhs := q
	for i := q; i < r-1; i++ {
		if a[i] >= a[r-1] {
			continue
		}
		a[rhs], a[i] = a[i], a[rhs]
		rhs++
	}

	a[rhs], a[r-1] = a[r-1], a[rhs]
	return rhs
}
