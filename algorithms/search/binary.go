package search

// Θ(log n)
func BinarySearch(data []int, number int) int {
	min_index := 0
	max_index := len(data) - 1

	for {
		test_index := (min_index + max_index) / 2

		// sanity check {min,max}_index
		if min_index > max_index {
			break
		}

		// out of bounds check for test_index
		if test_index > len(data)-1 || test_index < 0 {
			break
		}

		// test if index found
		if data[test_index] == number {
			return test_index
		}

		// since index not found, check if there is anything left
		if min_index == max_index {
			break
		}

		// set {min,max}_index to appropriate values
		if data[test_index] > number {
			max_index = test_index - 1
		} else {
			min_index = test_index + 1
		}
	}
	return -1
}
