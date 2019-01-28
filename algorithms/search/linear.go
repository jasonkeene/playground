package search

// Θ(n)
func Linear(data []int, number int) int {
	for i, num := range data {
		if num == number {
			return i
		}
	}
	return -1
}
