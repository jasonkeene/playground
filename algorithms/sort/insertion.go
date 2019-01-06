package sort

func Insertion(a []int) {
	for i := 1; i < len(a); i++ {
		key := a[i]
		var j int
		for j = i - 1; j >= 0; j-- {
			if a[j] <= key {
				break
			}
			a[j+1] = a[j]
		}
		a[j+1] = key
	}
}
