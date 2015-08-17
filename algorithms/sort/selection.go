package sort

// Θ(n^2)
func SelectionSort(data []int) {
	for i := 0; i < len(data); i++ {
		// find the smallest value
		min_index := minIndex(data, i)

		// swap it out
		Swap(data, i, min_index)
	}
}

// Θ(n)
func minIndex(data []int, startIndex int) int {
	min_index := startIndex
	for i := startIndex + 1; i < len(data); i++ {
		if data[i] < data[min_index] {
			min_index = i
		}
	}
	return min_index
}

// Θ(1)
func Swap(data []int, i, j int) {
	data[i], data[j] = data[j], data[i]
}
