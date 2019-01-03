package search

// Θ(log n)
func BinarySearch(data []int, number int) int {
	min_index := 0
	max_index := len(data) - 1

	for min_index <= max_index {
		test_index := (min_index + max_index) / 2

		// test if index found
		if data[test_index] == number {
			return test_index
		}

		// set indicies to appropriate values
		if data[test_index] > number {
			max_index = test_index - 1
		} else {
			min_index = test_index + 1
		}
	}

	return -1
}
