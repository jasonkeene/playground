package sort

func Insertion(a []int) {
	for i := 1; i < len(a); i++ {
		key := a[i]
		for j := i - 1; j >= 0; j-- {
			if a[j] <= key {
				a[j+1] = key
				break
			}
			a[j+1] = a[j]
		}
	}
}
