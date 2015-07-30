package search

// Θ(n)
func LinearSearch(data []int, number int) int {
	for i, num := range data {
		if num == number {
			return i
		}
	}
	return -1
}
