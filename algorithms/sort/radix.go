package sort

func Radix(a []string, size int) {
	for i := size - 1; i >= 0; i-- {
		counting(a, i)
	}
}

func counting(a []string, p int) {
	// equal
	var bucket [36]int
	for _, v := range a {
		bucket[bv[v[p]]]++
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
		i := bv[v[p]]
		sorted[bucket[i]] = v
		bucket[i]++
	}

	// replace
	for i, v := range sorted {
		a[i] = v
	}
}

var bv = func() [256]int {
	var bv [256]int

	for r := 0; r < 256; r++ {
		// digit
		if r >= 48 && r < 58 {
			bv[r] = r - 48
			continue
		}
		// capital letter
		if r >= 65 && r < 91 {
			bv[r] = r - 55
			continue
		}
	}
	return bv
}()
