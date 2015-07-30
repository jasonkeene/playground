package sort

// Θ(n^2)
func InsertionSort(data []int) {
	for i := 1; i < len(data); i++ {
		insert(data, i-1, data[i])
	}
}

// Θ(n)
func insert(data []int, sorted, key int) {
	var i int
	for i = sorted; i >= 0 && data[i] > key; i-- {
		// move over
		data[i+1] = data[i]
	}
	data[i+1] = key
}
