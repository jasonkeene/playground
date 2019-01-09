package sort

func Selection(a []int) {
	for i := 0; i < len(a)-1; i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		if i != min {
			a[i], a[min] = a[min], a[i]
		}
	}
}
