package sort

// Θ(n*log(n))
func QuickSort(data []int) []int {
	length := len(data)
	if length < 2 {
		return data
	}
	pivot := data[length-1]

	lower := make([]int, 0, length)
	middle := make([]int, 0, length)
	higher := make([]int, 0, length)

	for _, val := range data {
		switch {
		case val < pivot:
			lower = append(lower, val)
		case val == pivot:
			middle = append(middle, val)
		case val > pivot:
			higher = append(higher, val)
		}
	}
	lower, higher = QuickSort(lower), QuickSort(higher)
	lower = append(lower, middle...)
	lower = append(lower, higher...)
	return lower
}
