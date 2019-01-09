package sort

func Counting(max int) func([]int) {
	return func(a []int) {
		new := rearrange(a, keysLess(keysEqual(a, max), max), max)
		for i, v := range new {
			a[i] = v
		}
	}
}

func rearrange(a, less []int, max int) []int {
	b := make([]int, len(a))
	next := make([]int, max)

	for i := 0; i < max; i++ {
		next[i] = less[i]
	}
	for i := 0; i < len(a); i++ {
		key := a[i]
		b[next[key]] = key
		next[key]++
	}

	return b
}

func keysLess(equal []int, max int) []int {
	less := make([]int, len(equal))
	less[0] = 0
	for i := 1; i < len(equal); i++ {
		less[i] = equal[i-1] + less[i-1]
	}
	return less
}

func keysEqual(a []int, max int) []int {
	equal := make([]int, max)
	for i := 0; i < len(a); i++ {
		equal[a[i]]++
	}
	return equal
}

func ReallySimpleSort(a []int) {
	var count int
	for i := 0; i < len(a); i++ {
		if a[i] == 0 {
			count++
		}
	}
	for i := 0; i < len(a); i++ {
		if count > 0 {
			a[i] = 0
			count--
		} else {
			a[i] = 1
		}
	}
}
